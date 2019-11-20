package serp

/*
 * This package enables to interact with SerpApi server
 */

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Hold SerpApi user key
var apiKey string

// Query hold query parameter
type Query struct {
	parameter map[string]string
}

// Response hold response
type Response map[string]interface{}

// ResponseArray hold response array
type ResponseArray []interface{}

// NewGoogleSearch initialize the query
func NewGoogleSearch(parameter map[string]string) Query {
	if len(apiKey) > 0 {
		parameter["api_key"] = apiKey
	}
	return Query{parameter: parameter}
}

// Set your API KEY
func setAPIKey(key string) {
	apiKey = key
}

// GetJSON returns SerpResponse containing
func (sq *Query) GetJSON() (Results, error) {
	rsp, err := sq.execute("/search", "json")
	if err != nil {
		return Results{}, err
	}
	return sq.decodeJSON(rsp.Body)
}

// GetHTML returns html as a string
func (sq *Query) GetHTML() (*string, error) {
	rsp, err := sq.execute("/search", "html")
	if err != nil {
		return nil, err
	}
	return sq.decodeHTML(rsp.Body)
}

// GetLocation returns closest location
func GetLocation(q string, limit int) (ResponseArray, error) {
	client := NewGoogleSearch(map[string]string{
		"q":     q,
		"limit": string(limit),
	})
	rsp, err := client.execute("/locations.json", "json")
	if err != nil {
		return ResponseArray{}, err
	}
	return client.decodeJSONArray(rsp.Body)
}

// GetAccount return account information
// func GetAccount() (Results, error) {
// 	client := NewGoogleSearch(map[string]string{})
// 	rsp := client.execute("/account", "json")
// 	return client.decodeJSON(rsp.Body)
// }

// GetSearchArchive retrieve search from the archive using the Search Archive API
func (sq *Query) GetSearchArchive(searchID string) (Results, error) {
	rsp, err := sq.execute("/searches/"+searchID+".json", "json")
	if err != nil {
		return Results{}, err
	}
	return sq.decodeJSON(rsp.Body)
}

// decodeJson response
func (sq *Query) decodeJSON(body io.ReadCloser) (Results, error) {
	// Decode JSON from response body
	decoder := json.NewDecoder(body)
	// var serpResponse SerpResponse
	var resp Results
	err := decoder.Decode(&resp)
	if err != nil {
		return resp, fmt.Errorf("fail to decode: %w", err)
	}

	// check error message
	if resp.ErrorMessage != "" {
		return resp, errors.New(resp.ErrorMessage)
	}

	return resp, nil
}

// decodeJSONArray primitive function
func (sq *Query) decodeJSONArray(body io.ReadCloser) (ResponseArray, error) {
	decoder := json.NewDecoder(body)
	var rsp ResponseArray
	err := decoder.Decode(&rsp)
	if err != nil {
		return nil, errors.New("fail to decode array")
	}
	return rsp, nil
}

// decodeHTML primitive function
func (sq *Query) decodeHTML(body io.ReadCloser) (*string, error) {
	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	text := string(buffer)
	return &text, nil
}

// Execute the HTTP get
func (sq *Query) execute(path string, output string) (*http.Response, error) {
	query := url.Values{}
	for k, v := range sq.parameter {
		query.Add(k, v)
	}
	query.Add("source", "go")
	query.Add("output", output)
	endpoint := "https://serpapi.com" + path + "?" + query.Encode()
	var client = &http.Client{
		Timeout: time.Second * 60,
	}
	rsp, err := client.Get(endpoint)
	return rsp, err
}
