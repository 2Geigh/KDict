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
	API_URL = "https://krdict.korean.go.kr/api/search?key="
)

var (
	API_BASE_URL string
)

type dictSearch struct {
	Title   string     `xml:"title"`
	Total   int        `xml:"total"`
	Results []dictItem `xml:"item"`
}

type dictItem struct {
	Target_code      int              `xml:"target_code"`
	Word             string           `xml:"word"`
	Sup_no           int              `xml:"sup_no"`
	Etymology        string           `xml:"origin"`
	Pronunciation    string           `xml:"pronunciation"`
	Word_grade_level string           `xml:"word_grade"`
	Word_type        string           `xml:"pos"`
	Entry_link       string           `xml:"link"`
	Sense            dict_entry_sense `xml:"sense"`
}

type dict_entry_sense struct {
	Order      int    `xml:"sense_order"`
	Definition string `xml:"definition"`
}

func search_word(word string) (*exec.Cmd, error) {
	curl, err := exec.LookPath("curl")
	if err != nil {
		log.Fatal(err)
	}

	url_query := API_BASE_URL + "&q=" + url.QueryEscape(word)

	// Mozilla/5.0 is set as the header to mimic web browsers,
	// as the Korean government blocks generic headers to block scrapers
	cmd := exec.Command(curl, "-s", "-A", "Mozilla/5.0", url_query)
	return cmd, err
}

func parseXML(data any) {

}

func main() {

	godotenv.Load()
	API_KEY := os.Getenv("API_KEY")
	API_BASE_URL = API_URL + API_KEY

	curl, err := exec.LookPath("curl")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(curl, API_BASE_URL)

	fmt.Println(API_BASE_URL)
	fmt.Println()
	fmt.Println(cmd)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	search, err := search_word("한자")
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

		_, err = w.Write([]byte(fmt.Sprint(data)))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(fmt.Sprint(data))
	})
	log.Fatal(http.ListenAndServe(":3000", nil))

}
