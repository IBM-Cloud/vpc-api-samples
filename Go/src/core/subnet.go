package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// GetSubnets - request to get the list of subnets
func GetSubnets() {

	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/subnets" + QueryParams

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

// CreateSubnetTemplateInput - to create a request body
type CreateSubnetTemplateInput struct {
	Name          string          `json:"name"`
	NetworkACL    *ResourceByID   `json:"network_acl,omitempty"`
	PublicGateway *ResourceByID   `json:"public_gateway,omitempty"`
	Vpc           *ResourceByID   `json:"vpc"`
	Zone          *ResourceByName `json:"zone"`
	Ipv4CidrBlock string          `json:"ipv4_cidr_block"`
}

// CreateSubnetCountOnlyTemplateInput - to create a request body
type CreateSubnetCountOnlyTemplateInput struct {
	Name                  string          `json:"name"`
	NetworkACL            *ResourceByID   `json:"network_acl,omitempty"`
	PublicGateway         *ResourceByID   `json:"public_gateway,omitempty"`
	Vpc                   *ResourceByID   `json:"vpc"`
	Zone                  *ResourceByName `json:"zone"`
	TotalIpv4AddressCount int64           `json:"total_ipv4_address_count"`
}

// Subnet - Create a struct to mimic your json response structure
type Subnet struct {
	ID                        string     `json:"id"`
	Name                      string     `json:"name"`
	Href                      string     `json:"href"`
	AvailableIpv4AddressCount int        `json:"available_ipv4_address_count"`
	CreatedAt                 string     `json:"created_at"`
	Ipv4CidrBlock             string     `json:"ipv4_cidr_block"`
	NetworkACL                *Reference `json:"network_acl"`
	PublicGateway             *Reference `json:"public_gateway,omitempty"`
	Status                    string     `json:"status"`
	TotalIpv4AddressCount     int        `json:"total_ipv4_address_count"`
	Vpc                       *Reference `json:"vpc"`
	Zone                      *Reference `json:"zone"`
}

// PostSubnet - request to create a subnet
func PostSubnet(subnetInput interface{}) {
	// Create payload
	payload, err := json.Marshal(subnetInput)
	if err != nil {
		log.Fatal(err)
	}
	// Create URL adding endpoint, path to the resource and query parameters
	url := RiasEndpoint + "/subnets" + QueryParams

	// Create a new request given a method, URL, and optional body.
	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		log.Fatal(err)
	}

	// Adding headers to request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", IamToken)

	// Requesting server
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// Reading response and converting it to a JSON format
	decoder := json.NewDecoder(res.Body)
	var subnet Subnet
	err = decoder.Decode(&subnet)
	if err != nil {
		log.Fatal(err)
	}

	// Printing response
	fmt.Println("Response Status -", res.StatusCode)
	fmt.Println("Subnet created successfully!!")
	fmt.Println("Subnet ID-", subnet.ID)
	fmt.Println("Subnet Name-", subnet.Name)
	fmt.Println("Subnet VPC-", subnet.Vpc)
}
