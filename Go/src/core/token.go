package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Token - struct for response type for token api
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Expiration   int    `json:"expiration"`
	Scope        string `json:"scope"`
}

// RetrieveToken - func to retrieve token
func RetrieveToken(apikey string) {
	// Create payload
	payloadSlice := []string{"grant_type=urn%3Aibm%3Aparams%3Aoauth%3Agrant-type%3Aapikey&response_type=cloud_iam&apikey=", apikey}
	payload := strings.NewReader(strings.Join(payloadSlice, ""))

	// Create a new request given a method, URL, and optional body.
	req, err := http.NewRequest("POST", IAM_endpoint, payload)
	if err != nil {
		log.Fatal(err)
	}

	// Adding headers to the request
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	// Requesting server
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Reading response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// body is []byte format
	// parse the JSON-encoded body and stores the result in the struct object for the res
	var token Token
	json.Unmarshal([]byte(body), &token)
	// saving the token
	Iam_token = token.TokenType + " " + token.AccessToken
}
