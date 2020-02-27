package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// GetVSIs - request to get the list of security groups
func GetVSIs() {
	// Create URL adding endpoint, path to the resource and query parameters
	url := VPC_api_endpoint + "/instances" + QueryParams

	// Create a new request given a method, URL, and optional body.
	req, _ := http.NewRequest("GET", url, nil)

	// Adding headers to the request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", Iam_token)

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

// NetworkInterface - to create a request body
type NetworkInterface struct {
	Name           string          `json:"name"`
	PortSpeed      int             `json:"port_speed"`
	SecurityGroups []*ResourceByID `json:"security_groups,omitempty"`
	Subnet         *ResourceByID   `json:"subnet"`
	FloatingIPs    []*ResourceByID `json:"floating_ips,omitempty"`
}

// CreateVSIInput - to create a request body
type CreateVSIInput struct {
	Name                    string              `json:"name"`
	Keys                    []*ResourceByID     `json:"keys"`
	PrimaryNetworkInterface *NetworkInterface   `json:"primary_network_interface"`
	NetworkInterface        []*NetworkInterface `json:"network_interfaces,omitempty"`
	Profile                 *ResourceByName     `json:"profile"`
	Vpc                     *ResourceByID       `json:"vpc"`
	Zone                    *ResourceByName     `json:"zone"`
	Image                   *ResourceByID       `json:"image"`
}

// PostVSI - request to create a security group
func PostVSI(vsi *CreateVSIInput) {
	// Create payload
	payload, err := json.Marshal(vsi)
	if err != nil {
		log.Fatal(err)
	}

	// Create URL adding endpoint, path to the resource and query parameters
	url := VPC_api_endpoint + "/instances" + QueryParams

	// Create a new request given a method, URL, and optional body.
	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		log.Fatal(err)
	}

	// Adding headers to the request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", Iam_token)

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
