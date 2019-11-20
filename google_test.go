package serp

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func setup() {
	v := os.Getenv("API_KEY")
	if len(v) == 0 {
		setAPIKey("demo")
	} else {
		setAPIKey(v)
	}
}

func shoulSkip() bool {
	return len(os.Getenv("API_KEY")) == 0
}

func TestQuickStart(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	setup()

	parameter := map[string]string{
		"q":             "Coffee",
		"location":      "Portland, Oregon, United States",
		"hl":            "en",
		"gl":            "us",
		"google_domain": "google.com",
		"api_key":       os.Getenv("API_KEY"),
		"safe":          "active",
		"start":         "10",
		"num":           "10",
		"device":        "desktop",
	}

	client := NewGoogleSearch(parameter)
	rsp, err := client.GetJSON()

	if err != nil {
		t.Error(err)
		return
	}

	if len(rsp.OrganicResults[0].Title) == 0 {
		t.Error("empty title in local results")
		return
	}
}

// basic use case
func TestJSON(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	setup()
	parameter := map[string]string{
		"api_key":  os.Getenv("API_KEY"),
		"q":        "Coffee",
		"location": "Portland"}

	client := NewGoogleSearch(parameter)
	rsp, err := client.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}
	result := rsp.OrganicResults[0]
	if len(result.Title) == 0 {
		t.Error("empty title in local results")
		return
	}
}

func TestJSONwithGlobalKey(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	setup()
	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Portland"}

	client := NewGoogleSearch(parameter)
	rsp, err := client.GetJSON()
	if err != nil {
		t.Error("unexpected error", err)
		return
	}
	result := rsp.OrganicResults[0]
	if len(result.Title) == 0 {
		t.Error("empty title in local results")
		return
	}
}

func TestGetHTML(t *testing.T) {
	if shoulSkip() {
		t.Skip("API_KEY required")
		return
	}

	parameter := map[string]string{
		"q":        "Coffee",
		"location": "Portland"}

	setup()

	client := NewGoogleSearch(parameter)
	data, err := client.GetHTML()
	if err != nil {
		t.Error("err must be nil")
		return
	}
	if !strings.Contains(*data, "</html>") {
		t.Error("data does not contains <html> tag")
	}
}

func TestDecodeJson(t *testing.T) {
	reader, err := os.Open("./data/search_coffee_sample.json")
	if err != nil {
		panic(err)
	}
	var sq Query
	rsp, err := sq.decodeJSON(reader)
	if err != nil {
		t.Error("error should be nil", err)
		return
	}

	if rsp.OrganicResults[0].Title != "Portland Roasting Coffee" {
		t.Error("empty title in local results")
		return
	}
}

func TestDecodeJsonPage20(t *testing.T) {
	t.Log("run test")
	reader, err := os.Open("./data/search_coffee_sample_page20.json")
	if err != nil {
		panic(err)
	}
	var sq Query
	rsp, err := sq.decodeJSON(reader)
	if err != nil {
		t.Error("error should be nil")
		t.Error(err)
	}
	t.Log(reflect.ValueOf(rsp).MapKeys())

	if rsp.OrganicResults[0].Title != "Coffee | HuffPost" {
		t.Error("fail decoding the title ")
	}
}

func TestDecodeJsonError(t *testing.T) {
	reader, err := os.Open("./data/error_sample.json")
	if err != nil {
		panic(err)
	}
	var sq Query
	_, err = sq.decodeJSON(reader)
	if err != nil {
		t.Error(err)
		return
	}

	if err == nil {
		t.Error("unexcepted err is nil")
	} else if strings.Compare(err.Error(), "Your account credit is too low, plesae add more credits.") == 0 {
		t.Error("empty title in local results")
		return
	}
}

func TestGetLocation(t *testing.T) {
	setup()

	var rsp ResponseArray
	var err error
	rsp, err = GetLocation("Austin", 3)

	if err != nil {
		t.Error(err)
	}

	//log.Println(rsp[0])
	first := rsp[0].(map[string]interface{})
	googleID := first["google_id"].(float64)
	if googleID != float64(200635) {
		t.Error(googleID)
		return
	}
}

// func TestGetAccount(t *testing.T) {
// 	// Skip this test
// 	if len(os.Getenv("API_KEY")) == 0 {
// 		t.Skip("API_KEY required")
// 		return
// 	}

// 	setup()

// 	var rsp SerpResponse
// 	var err error
// 	rsp, err = GetAccount()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if rsp["account_id"] == nil {
// 		t.Error("no account_id found")
// 		return
// 	}
// }

// Search archive API
func TestSearchArchive(t *testing.T) {
	if len(os.Getenv("API_KEY")) == 0 {
		t.Skip("API_KEY required")
		return
	}

	setup()
	parameter := map[string]string{
		"api_key":  os.Getenv("API_KEY"),
		"q":        "Coffee",
		"location": "Portland"}

	client := NewGoogleSearch(parameter)
	rsp, err := client.GetJSON()

	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	searchID := rsp.SearchMetadata.ID

	if len(searchID) == 0 {
		t.Error("search_metadata.id must be defined")
	}

	searchArchive, err := client.GetSearchArchive(searchID)
	if err != nil {
		t.Error(err)
		return
	}

	if searchArchive.SearchMetadata.ID != searchID {
		t.Error("search_metadata.id do not match", searchArchive.SearchMetadata.ID, searchID)
	}
}
