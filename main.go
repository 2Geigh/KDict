package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	// "encoding/xml"

	"github.com/joho/godotenv"
)

var (
	API_BASE_URL string
)

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

func main() {

	godotenv.Load()
	API_KEY := os.Getenv("API_KEY")
	API_BASE_URL = "https://krdict.korean.go.kr/api/search?key=" + API_KEY

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

	// fmt.Println(search)
	out, err := search.Output()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(out))

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write(out)
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(":3000", nil))

}
