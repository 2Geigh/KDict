package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	apiUrlWithoutKey = "https://krdict.korean.go.kr/api/search?key="
)

type templateData struct {
	SearchQuery   string
	SearchResults []dictSearch
}

type dictSearch struct {
	Title   string     `xml:"title"`
	Total   int        `xml:"total"`
	Results []dictItem `xml:"item"`
}

type dictItem struct {
	Target_code      int            `xml:"target_code"`
	Word             string         `xml:"word"`
	Sup_no           int            `xml:"sup_no"`
	Etymology        string         `xml:"origin"`
	Pronunciation    string         `xml:"pronunciation"`
	Word_grade_level string         `xml:"word_grade"`
	Word_type        string         `xml:"pos"`
	Entry_link       string         `xml:"link"`
	Sense            dictEntrySense `xml:"sense"`
}

type dictEntrySense struct {
	Order      int    `xml:"sense_order"`
	Definition string `xml:"definition"`
}

func fetchDictionaryData(word string, urlWithApiKey string) (dictSearch, error) {

	urlWithQuery := urlWithApiKey + "&q=" + url.QueryEscape(word)

	req, err := http.NewRequest("GET", urlWithQuery, nil)
	if err != nil {
		return dictSearch{}, fmt.Errorf("Failed to create request: %v", err)
	}

	// Mozilla/5.0 is set as the header to mimic web browsers
	// because the Korean government blocks generic headers to block scrapers
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return dictSearch{}, fmt.Errorf("Failed to execute request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return dictSearch{}, fmt.Errorf("HTTP status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dictSearch{}, fmt.Errorf("Could not read response body: %v", err)
	}

	var xml_data dictSearch
	err = xml.Unmarshal(body, &xml_data)
	if err != nil {
		return dictSearch{}, fmt.Errorf("Could not parse XML: %v", err)
	}

	return xml_data, nil
}

func removeEmojis(input string) string {
	// This regex pattern matches emojis by their Unicode ranges using the correct escape sequence.
	emojiRegex := regexp.MustCompile(
		`[\x{1F600}-\x{1F64F}` + // Emoticons
			`|\x{1F300}-\x{1F5FF}` + // Miscellaneous Symbols and Pictographs
			`|\x{1F680}-\x{1F6FF}` + // Transport and Map Symbols
			`|\x{1F700}-\x{1F8FF}` + // Alchemical Symbols
			`|\x{1F900}-\x{1F9FF}` + // Supplemental Symbols and Pictographs
			`|\x{2700}-\x{27BF}` + // Dingbats
			`|\x{1F1E6}-\x{1F1FF}]`) // Regional Indicator Symbols

	return emojiRegex.ReplaceAllString(input, "")
}

func resultsHandler(w http.ResponseWriter, req *http.Request, apiUrl string) {

	fmt.Println()

	var (
		data = templateData{}
	)

	data.SearchQuery = req.FormValue("search_query")

	if data.SearchQuery == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	data.SearchQuery = removeEmojis(data.SearchQuery)

	if strings.Contains(data.SearchQuery, " ") {
		words, err := parseSentence(data.SearchQuery)
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		// Send each parsed word to API
		for _, v := range words {
			wordSearchData, err := fetchDictionaryData(v, apiUrl)
			if err != nil {
				http.Error(w, fmt.Sprint(err), 500)
				return
			}

			data.SearchResults = append(data.SearchResults, wordSearchData)
		}
	} else {
		wordSearchData, err := fetchDictionaryData(data.SearchQuery, apiUrl)
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		data.SearchResults = append(data.SearchResults, wordSearchData)
	}

	tmpl := template.Must(template.ParseFiles("./templates/results.html"))

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		return
	}

}

func parseSentence(query string) ([]string, error) {

	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Could not get working directory: %v\n", err)
	}
	fmt.Println("Current working directory:", dir)

	filename := "parseSentence.py"
	pythonProgramPath := fmt.Sprintf("%s/src/sentenceParsing/%s", dir, filename)

	// Create parseSentence.py terminal call
	cmd := exec.Command("python", pythonProgramPath, query)

	// Capture the standard output
	var (
		out    bytes.Buffer
		stderr bytes.Buffer // To capture standard error
	)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Execute command and wait for it to complete
	log.Printf("Calling %s...", filename)
	start := time.Now()
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to call %s: %v, stderr: %s", filename, err, stderr.String())
	}
	elapsed := time.Since(start)
	log.Printf("%s call completed in %v\n", filename, elapsed)

	// Get output as string
	output := out.String()

	// Split the output into words by line
	words := strings.Split((strings.TrimSpace(output)), "\n")

	return words, err

}

func main() {

	// Create API query
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	apiUrlWithKey := apiUrlWithoutKey + apiKey

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "index.html")
	})

	// route "/results"
	http.HandleFunc("/results", func(w http.ResponseWriter, req *http.Request) {
		resultsHandler(w, req, apiUrlWithKey)
	})

	// Start server
	log.Fatal(http.ListenAndServe(":3000", nil))

}
