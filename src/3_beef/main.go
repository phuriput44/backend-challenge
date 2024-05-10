package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type BeefSummary map[string]int

func isBeef(word string) bool {
	beefNames := map[string]bool{
		"bacon":    true,
		"filet":    true,
		"mignon":   true,
		"ribeye":   true,
		"flank":    true,
		"t-bone":   true,
		"pastrami": true,
		"pork":     true,
		"meatloaf": true,
		"jowl":     true,
		"bresaola": true,
	}

	_, found := beefNames[word]
	return found
}

func CountBeef(text string) BeefSummary {
	counts := make(BeefSummary)

	words := strings.Fields(text)

	for _, word := range words {
		if isBeef(word) {
			counts[word]++
		}
	}

	return counts
}

func BeefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	text := string(body)

	summary := CountBeef(text)

	jsonData, err := json.Marshal(summary)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/beef/summary", BeefSummaryHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
