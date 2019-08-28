package main

import (
	"fmt"
	"os"

	g "github.com/ericgreene/google-search-results-golang"
)

/***
 * demo how to create a client for SerpApi
 *
 * go get -u github.com/serpapi/google_search_results_golang
 */
func main() {
	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Austin,Texas",
		"api_key":  os.Getenv("API_KEY"), // your api key
	}

	client := g.NewGoogleSearch(parameter)
	serpResponse, err := client.GetJSON()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(serpResponse.OrganicResults[0].Title)
}
