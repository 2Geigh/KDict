package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	// "html/template"

	"github.com/joho/godotenv"
)

const (
	apiUrlWithoutKey = "https://krdict.korean.go.kr/api/search?key="
)

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

func fetchDictionaryData(word string, urlWithApiKey string) ([]byte, error) {

	urlWithQuery := urlWithApiKey + "&q=" + url.QueryEscape(word)

	// Create HTTP client request
	req, err := http.NewRequest("GET", urlWithQuery, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}

	// Mozilla/5.0 is set as the header to mimic web browsers,
	// as the Korean government blocks generic headers to block scrapers
	req.Header.Set("User-Agent", "Mozilla/5.0")

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute request: %v", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code: %v", resp.StatusCode)
	}

	// Read response body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read response body: %v", err)
	}

	return body, nil

}

func parseXML(data any) {

}

func main() {

	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	apiUrlWithKey := apiUrlWithoutKey + apiKey

	curl, err := exec.LookPath("curl")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(curl, apiUrlWithKey)

	fmt.Println(apiUrlWithKey)
	fmt.Println()
	fmt.Println(cmd)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	search, err := fetchDictionaryData("한자", apiUrlWithKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(search)

	// out, err := search.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(out))

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		xmlData := search

		// Unmarshal XML data
		// err := xml.Unmarshal(out, &xmlData)
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("Failed to parse XML: %v", err), http.StatusBadRequest)
		// }

		// Set proper content type and status
		w.Header().Set("Content-Type", "text/xml")
		w.WriteHeader(http.StatusOK)

		// _, err = w.Write([]byte(fmt.Sprint(xmlData)))
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("Failed to write response: %v", err), http.StatusInternalServerError)
		// }

		fmt.Println(fmt.Sprint(xmlData))
	})
	log.Fatal(http.ListenAndServe(":3000", nil))

}
