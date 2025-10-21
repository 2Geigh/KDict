package controllers

import (
	"KDict/src/config"
	"KDict/src/models"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func RootHandler(w http.ResponseWriter, req *http.Request) {

	http.ServeFile(w, req, "../index.html")

}

func ResultsHandler(w http.ResponseWriter, req *http.Request) {

	// For ease of readability in the terminal
	fmt.Println()

	data := models.TemplateData{}

	// Get ƒorm data
	data.SearchQuery = req.FormValue("search_query")

	// Validate form data
	if data.SearchQuery == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// Query before emoji removal
	fmt.Printf("Pre-de-emoji: %s\n", data.SearchQuery)

	// Remove emojis from form data
	data.SearchQuery = removeEmojis(data.SearchQuery)

	// Query after emoji removal
	fmt.Printf("Post-de-emoji: %s\n", data.SearchQuery)

	// Skip parsing step if the query only has one word
	if strings.Contains(data.SearchQuery, " ") {
		// Parse multi-word query
		words, err := parseSentence(data.SearchQuery)
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}
		fmt.Println(words)

		// Send each parsed word to API
		for _, v := range words {
			wordSearchData, err := fetchDictionaryData(v, config.ApiUrlWithKey)
			if err != nil {
				http.Error(w, fmt.Sprint(err), 500)
				return
			}

			// Append found data to the SearchResults field of the data variable
			data.SearchResults = append(data.SearchResults, wordSearchData)
		}
	} else {
		wordSearchData, err := fetchDictionaryData(data.SearchQuery, config.ApiUrlWithKey)
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		// Append found data to the SearchResults field of the data variable
		data.SearchResults = append(data.SearchResults, wordSearchData)
	}

	// Parse HTML template
	tmpl := template.Must(template.ParseFiles("../templates/results.html"))

	// Insert data into parsed HTML template
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		return
	}

}
