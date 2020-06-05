# Getting Started with VPC on Classic

This repository consists of a variety of code snippets intended to aid developer interaction with resources belonging to VPC IaaS offering. This repository is intended to be used for documentation purposes only and not to be included as a dependency.

Examples are available in the following languages:
1. [Go](#go)
2. [Python](#python)

Examples break down into following steps:
1. Retrieve API key for your account.
2. Get an IAM access token using your api key.
3. Get a list of the resources.
4. Post a resource.

## Retrieve API key for your account.

  Instructions to retrieve your API key [here](https://cloud.ibm.com/docs/iam?topic=iam-userapikey#create_user_key)

## List of available resources

| Resource          | Respective endpoint   |
| ----------------- |:---------------------:|
| VPC               | /vpcs                 |
| Subnets           | /subnets              |
| Profiles          | /profiles             |
| Security Groups   | /security_groups      |
| Images            | /images               |
| Floating IPs      | /floating_ips         |
| Instances         | /instances            |
| Regions           | /regions              |
| Volumes           | /volumes              |
| Network ACL       | /network_acls         |
| Regions           | /regions              |
| Volumes           | /volumes              |
| Load Balancers    | /load_balancers       |
| SSH keys          | /keys                 |


*VPC API Version*

Refer to the [VPC on Classic Release Notes](https://cloud.ibm.com/docs/vpc-on-classic?topic=vpc-on-classic-release-notes) document to find out latest version release.


## Go

This section displays the process to call VPC APIs in a Go workspace.


1. Set up global/account variables.
  Set up following global variables in your workspace

    ```go
    var Iam_token string
    const API_version = "2020-06-02"
    const Generation = "1"
    const VPC_api_endpoint = "https://us-south.iaas.cloud.ibm.com/v1"
    const IAM_endpoint = "https://iam.cloud.ibm.com/identity/token"
    const API_key = "Your API key here"
    const QueryParams = `?version=` + API_version + `&generation=` + Generation
    ```


2. Get an IAM access token using your API key.

      Define a struct to hold token returned by API.

      ```go
      type Token struct {
        AccessToken  string `json:"access_token"`
        RefreshToken string `json:"refresh_token"`
        TokenType    string `json:"token_type"`
        ExpiresIn    int    `json:"expires_in"`
        Expiration   int    `json:"expiration"`
        Scope        string `json:"scope"`
      }

      ```


      Make the payload.

      ```go
      payloadSlice := []string{"grant_type=urn%3Aibm%3Aparams%3Aoauth%3Agrant-type%3Aapikey&response_type=cloud_iam&apikey=", apikey}
      payload := strings.NewReader(strings.Join(payloadSlice, ""))
      ```

      Make the request by passing the required endpoint and payload.

      ```go
      req, err := http.NewRequest("POST", IAMEndpoint, payload)
      ```

      Set the headers.

      ```go
      req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
      req.Header.Add("Accept", "application/json")
      ```

      Request server and unmarshall the body in the `Token` type struct. Set the global variable.

      ```go
      res, err := http.DefaultClient.Do(req)
      body, err := ioutil.ReadAll(res.Body)
      var token Token
      json.Unmarshal([]byte(body), &token)
      IamToken = token.TokenType + " " + token.AccessToken
      ```


      Now, you should have the token stored in global variable `IamToken`.
      Refer to the code [here.](https://github.com/IBM-Cloud/vpc-api-samples/blob/master/Go/src/core/token.go)


3. Get a list of the resources.

    This section shows how to perform a GET API call on VPC APIs.

    Create the URL to be used to make GET rest API call.

    ```go
    url := API_version + "/subnets" + QueryParams
    ```
    This URL will get the list of subnets. Get other resources by using the appropriate endpoint.


    Create a new request given a method, above URL, and optional body.

    ```go
    req, err := http.NewRequest("GET", url, nil)
    ```


    Set the headers and IAM token.

    ```go
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", IamToken)
    ```


    Request server and read the response.

    ```go
    res, err := http.DefaultClient.Do(req)
    body, err := ioutil.ReadAll(res.Body)
    // Printing response
    fmt.Println("Response Status -", res.StatusCode)
    fmt.Println("Response Body -", string(body))
    ```


4. Post a resource.

    4-1. Define a struct for the input. Subnet POST API accepts two kinds of request body defined by the following structs.

    ```go
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
    ```

      4-2. Define a struct for the response. The structure for response body is defined in API [spec](https://cloud.ibm.com/apidocs/vpc-on-classic).

    ```go
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
    ```

    The POST call is shown below.

    ```go
    // PostSubnet - request to create a subnet
    func PostSubnet(subnetInput interface{}) {
      // Create payload
      payload, err := json.Marshal(subnetInput)
      if err != nil {
        log.Fatal(err)
      }
      // Create URL adding endpoint, path to the resource and query parameters
      url := API_version + "/subnets" + QueryParams

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
    }
    ```

    Calling `PostSubnet` function.

    ```go
    vpcID := &core.ResourceByID{ID: "VPC_ID"}
    zone := &core.ResourceByName{Name: "ZONE_ID"}
    subnetCountOnly := &core.CreateSubnetCountOnlyTemplateInput{
      Name: "SUBNET_NAME", Vpc: vpcID, Zone: zone,
      TotalIpv4AddressCount: 8, //number of addresses
    }
    core.PostSubnet(subnetCountOnly)
    ```


## Python

The example below shows how to retrieve a token, list all VPCs, and create a VPC.

1. Get an IAM access token using your API key.

```python
import http.client
import json

# URL for token
conn = http.client.HTTPSConnection("iam.cloud.ibm.com")

# Payload for retrieving token. Note: An API key will need to be generated and replaced here
payload = 'grant_type=urn%3Aibm%3Aparams%3Aoauth%3Agrant-type%3Aapikey&apikey=YOUR_API_KEY&response_type=cloud_iam'

# Required headers
headers = {
    'Content-Type': 'application/x-www-form-urlencoded',
    'Accept': 'application/json',
    'Cache-Control': 'no-cache'
}

try:
    # Connect to endpoint for retrieving a token
    conn.request("POST", "/identity/token", payload, headers)

    # Get and read response data
    res = conn.getresponse().read()
    data = res.decode("utf-8")

    # Format response in JSON
    json_res = json.loads(data)

    # Concatenate token type and token value
    return json_res['token_type'] + ' ' + json_res['access_token']

# If an error happens while retrieving token
except Exception as error:
    print(f"Error getting token. {error}")
    raise
```



2. Get a list of all VPCs.

```python
import http.client
import json

region = "us-south"

conn = http.client.HTTPSConnection(f"{region}.iaas.cloud.ibm.com")

headers = {
    'Content-Type': 'application/json',
    'Cache-Control': 'no-cache',
    'Accept': 'application/json',
    'Authorization': YOUR_TOKEN,
    'cache-control': 'no-cache'
}

version = "2020-06-02"

payload = ""

try:
    # Connect to api endpoint for vpcs
    conn.request("GET", "/v1/vpcs?generation=1&version=" + version, payload, headers)

    # Get and read response data
    res = conn.getresponse()
    data = res.read()

    # Print and return response data
    print(json.dumps(json.loads(data.decode("utf-8")), indent=2, sort_keys=True))
    return data.decode("utf-8")

except Exception as error:
    print(f"Error fetching VPCs. {error}")
    raise
```

3. Create a VPC.

```python
import http.client
import json

region = "us-south"

conn = http.client.HTTPSConnection(f"{region}.iaas.cloud.ibm.com")

headers = {
    'Content-Type': 'application/json',
    'Cache-Control': 'no-cache',
    'Accept': 'application/json',
    'Authorization': YOUR_TOKEN,
    'cache-control': 'no-cache'
}

version = "2020-06-02"

# Required payload for creating a VPC
payload = f'{{"name": "NAME_OF_VPC"}}'

try:
    # Connect to api endpoint for vpcs
    conn.request("POST", "/v1/vpcs?generation=1&version=" + version, payload, headers)

    # Get and read response data
    res = conn.getresponse()
    data = res.read()

    # Print and return response data
    print_json(data.decode("utf-8"))
    return data.decode("utf-8")

# If an error happens while creating a VPC
except Exception as error:
    print(f"Error creating VPC. {error}")
    raise
```

## API Spec
[IBM Cloud VPC on Classic API](https://cloud.ibm.com/apidocs/vpc-on-classic)

## Documentation
[IBM Cloud Virtual Private Cloud](https://cloud.ibm.com/docs/vpc-on-classic?topic=vpc-on-classic-getting-started)
