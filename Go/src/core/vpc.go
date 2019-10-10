package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// GetVPCs - request to get the list of vpcs
func GetVPCs() {
	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/vpcs" + QueryParams

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

// CreateVPCInput - to create a request body
type CreateVPCInput struct {
	Name          string        `json:"name"`
	DefaultACL    *ResourceByID `json:"default_network_acl"`
	ClassicPeered string        `json:"classic_peered,omitempty"`
	ResourceGroup *ResourceByID `json:"resource_group,omitempty"`
}

// PostVPC - request to create a vpc
func PostVPC(vpcInput *CreateVPCInput) {
	// Create payload
	payload, err := json.Marshal(vpcInput)
	if err != nil {
		log.Fatal(err)
	}
	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/vpcs" + QueryParams
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
