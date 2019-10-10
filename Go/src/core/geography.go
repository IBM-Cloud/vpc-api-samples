package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetRegions - request to get the list of regions
func GetRegions() {
	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/regions" + RiasVersion

	// Create a new request given a method, URL, and optional body.
	req, err := http.NewRequest("GET", url, nil)
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

// GetZones - request to get the list of regions
func GetZones(regionName string) {
	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + `/regions/` + regionName + `/zones` + RiasVersion

	// Create a new request given a method, URL, and optional body.
	req, err := http.NewRequest("GET", url, nil)
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
