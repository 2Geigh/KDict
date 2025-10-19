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

	// Create HTTP client request
	req, err := http.NewRequest("GET", urlWithQuery, nil)
	if err != nil {
		return dictSearch{}, fmt.Errorf("Failed to create request: %v", err)
	}

	// Mozilla/5.0 is set as the header to mimic web browsers,
	// as the Korean government blocks generic headers to block scrapers
	req.Header.Set("User-Agent", "Mozilla/5.0")

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return dictSearch{}, fmt.Errorf("Failed to execute request: %v", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return dictSearch{}, fmt.Errorf("HTTP status code: %v", resp.StatusCode)
	}

	// Read response body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dictSearch{}, fmt.Errorf("Could not read response body: %v", err)
	}

	// Parse the response body's XML
	var xml_data dictSearch
	err = xml.Unmarshal(body, &xml_data)
	if err != nil {
		return dictSearch{}, fmt.Errorf("Could not parse XML: %v", err)
	}

	// log.Printf("Successfully fetched dictionary data for '%s'", word)
	return xml_data, nil
}

func resultsHandler(w http.ResponseWriter, req *http.Request, apiUrl string) {

	var (
		data = templateData{}
	)

	// Get ƒorm data
	data.SearchQuery = req.FormValue("search_query")

	// Validate form data
	if data.SearchQuery == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// Parse sentence
	words, err := parseSentence(data.SearchQuery)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		return
	}
	fmt.Println(words)

	// Send each parsed word to API
	for _, v := range words {
		wordSearchData, err := fetchDictionaryData(v, apiUrl)
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		// Append found data to the SearchResults field of the data variable
		data.SearchResults = append(data.SearchResults, wordSearchData)
	}
	fmt.Println(data)

	// data.SearchResults, err = fetchDictionaryData(data.SearchQuery, apiUrl)
	// if err != nil {
	// 	http.Error(w, fmt.Sprint(err), 500)
	// 	return
	// }

	// Parse HTML template
	tmpl := template.Must(template.ParseFiles("./templates/results.html"))

	// Insert data into parsed HTML template
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		return
	}

}

func parseSentence(query string) ([]string, error) {

	// Get current directory
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
	var out bytes.Buffer
	cmd.Stdout = &out

	// Execute command and wait for it to complete
	log.Printf("Calling %s...", filename)
	start := time.Now()
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to call %s: %v\n", filename, err)
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

// For item in dictSearch.Results
// {
// 	한국어 기초사전 개발 지원(Open API) - 사전 검색
// 	5
// 	[
// 		{
// 			72461
// 			한자
// 			0
// 			漢字
// 			한ː짜
// 			중급
// 			명사
// 			https://krdict.korean.go.kr/kor/dicSearch/SearchView?ParaWordNo=72461
// 			{
// 				1
// 				중국에서 만들어 오늘날에도 쓰고 있는 중국 고유의 문자.
// 			}
// 		}

// 		{
// 			85621
// 			한자리
// 			0
// 			한자리
// 			고급
// 			명사
// 			https://krdict.korean.go.kr/kor/dicSearch/SearchView?ParaWordNo=85621
// 			{
// 				2
// 				중요하거나 높은 직위. 또는 어느 한 직위.
// 			}
// 		}

// 		{93420 한자어 0 漢字語 한ː짜어 고급 명사 https://krdict.korean.go.kr/kor/dicSearch/SearchView?ParaWordNo=93420 {1 한자에 기초하여 만들어진 말.}} {85622 한자리하다 0  한자리하다  동사 https://krdict.korean.go.kr/kor/dicSearch/SearchView?ParaWordNo=85622 {1 중요하거나 높은 직위에 오르다.}} {88977 한자음 0 漢字音 한ː짜음  명사 https://krdict.korean.go.kr/kor/dicSearch/SearchView?ParaWordNo=88977 {1 한자의 발음이나 소리.}}]}
