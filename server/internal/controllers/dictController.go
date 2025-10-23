package controllers

import (
	"KDict/src/models"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func fetchDictionaryData(word string, urlWithApiKey string) (models.DictSearch, error) {

	urlWithQuery := urlWithApiKey + "&q=" + url.QueryEscape(word)

	// Create HTTP client request
	req, err := http.NewRequest("GET", urlWithQuery, nil)
	if err != nil {
		return models.DictSearch{}, fmt.Errorf("Failed to create request: %v", err)
	}

	// Mozilla/5.0 is set as the header to mimic web browsers,
	// as the Korean government blocks generic headers to block scrapers
	req.Header.Set("User-Agent", "Mozilla/5.0")

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.DictSearch{}, fmt.Errorf("Failed to execute request: %v", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return models.DictSearch{}, fmt.Errorf("HTTP status code: %v", resp.StatusCode)
	}

	// Read response body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.DictSearch{}, fmt.Errorf("Could not read response body: %v", err)
	}

	// Parse the response body's XML
	var xml_data models.DictSearch
	err = xml.Unmarshal(body, &xml_data)
	if err != nil {
		return models.DictSearch{}, fmt.Errorf("Could not parse XML: %v", err)
	}

	// log.Printf("Successfully fetched dictionary data for '%s'", word)
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

func parseSentence(query string) ([]string, error) {

	// Get current directory
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Could not get working directory: %v\n", err)
	}
	fmt.Println("Current working directory:", dir)

	filename := "parseSentence.py"
	pythonProgramPath := fmt.Sprintf("%s/src/utils/sentenceParsing/%s", dir, filename)

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
