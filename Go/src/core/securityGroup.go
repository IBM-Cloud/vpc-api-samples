package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// GetSecurityGroups - request to get the list of security groups
func GetSecurityGroups() {
	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/security_groups" + QueryParams
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

// Remote - to create json with with id
type Remote struct {
	CidrBlock string `json:"cidr_block"`
}

// Rule - to create a request body
type Rule struct {
	Direction string  `json:"direction"`
	IPVersion string  `json:"ip_version,omitempty"`
	Remote    *Remote `json:"remote"`
	Protocol  string  `json:"protocol,omitempty"`
}

// CreateSecurityGroupInput - to create a request body
type CreateSecurityGroupInput struct {
	Name  string        `json:"name"`
	Rules []*Rule       `json:"rules"`
	Vpc   *ResourceByID `json:"vpc"`
}

// PostSecurityGroup - request to create a security group
func PostSecurityGroup(sg *CreateSecurityGroupInput) {
	// Create payload
	payload, err := json.Marshal(sg)
	if err != nil {
		log.Fatal(err)
	}
	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/security_groups" + QueryParams
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Printing response
	fmt.Println("Response Status -", res.StatusCode)
	fmt.Println("Response Body -", string(body))
}
