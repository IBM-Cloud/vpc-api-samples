package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetProfiles - request to get the list of instance profiles
func GetProfiles() {

	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/instance/profiles" + QueryParams

	// Create a new request given a method, URL, and optional body.
	req, err := http.NewRequest("GET", url, nil)

	// Handle error occured while creating a new request
	if err != nil {
		log.Fatal(err)
	}

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
