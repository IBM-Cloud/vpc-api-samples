package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// GetSSHKeys - request to get the list of ssh keys
func GetSSHKeys() {
	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/keys" + QueryParams

	// Create a new request given a method, URL, and optional body.
	req, _ := http.NewRequest("GET", url, nil)

	// Adding headers to the request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", IamToken)

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

	// Printing response
	fmt.Println("Response Status -", res.StatusCode)
	fmt.Println("Response Body -", string(body))
}

// CreateSSHKeyInput - to create a request body
type CreateSSHKeyInput struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
	Type      string `json:"type"`
}

// PostSSHKey - request to create ssh key
func PostSSHKey(sskKeyInput *CreateSSHKeyInput) {
	// Create payload
	payload, err := json.Marshal(sskKeyInput)
	if err != nil {
		log.Fatal(err)
	}

	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/keys" + QueryParams

	// Create a new request given a method, URL, and optional body.
	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		log.Fatal(err)
	}

	// Adding headers to the request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", IamToken)

	// Reading response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Requesting server
	body, _ := ioutil.ReadAll(res.Body)

	// Printing response
	fmt.Println("Response Status -", res.StatusCode)
	fmt.Println("Response Body -", string(body))
}
