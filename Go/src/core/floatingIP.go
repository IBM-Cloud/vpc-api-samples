package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// GetFloatingIPs - request to get the list of floating ips
func GetFloatingIPs() {

	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/floating_ips" + RiasVersion

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

// CreateFloatingIPInput - to create a request body
type CreateFloatingIPInput struct {
	Name   string          `json:"name"`
	Zone   *ResourceByName `json:"zone,omitempty"`
	Target *ResourceByID   `json:"target"`
}

// PostReserveFloatingIP - request to create a security group
func PostReserveFloatingIP(fip *CreateFloatingIPInput) {
	// Create payload
	payload, _ := json.Marshal(fip)

	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/floating_ips" + QueryParams

	// Create a new request given a method, URL, and optional body.
	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
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
	body, _ := ioutil.ReadAll(res.Body)

	// Printing response
	fmt.Println("Response Status -", res.StatusCode)
	fmt.Println("Response Body -", string(body))
}
