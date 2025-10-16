package main

import (
	"encoding/xml"
	"fmt"
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

func searchWord(word string, urlWithApiKey string) (*exec.Cmd, error) {
	curl, err := exec.LookPath("curl")
	if err != nil {
		log.Fatal(err)
	}

	url_query := urlWithApiKey + "&q=" + url.QueryEscape(word)

	// Mozilla/5.0 is set as the header to mimic web browsers,
	// as the Korean government blocks generic headers to block scrapers
	cmd := exec.Command(curl, "-s", "-A", "Mozilla/5.0", url_query)
	return cmd, err
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

	search, err := searchWord("한자", apiUrlWithKey)
	if err != nil {
		log.Fatal(err)
	}

	out, err := search.Output()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(out))

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var data dictSearch

		err := xml.Unmarshal(out, &data)
		if err != nil {
			log.Fatal(err)
		}

		// Set proper content type and status
		w.Header().Set("Content-Type", "text/xml")
		w.WriteHeader(http.StatusOK)

		_, err = w.Write([]byte(fmt.Sprint(xmlData)))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(fmt.Sprint(data))
	})
	log.Fatal(http.ListenAndServe(":3000", nil))

}
