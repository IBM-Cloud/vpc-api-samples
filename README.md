# Getting Started

This repo is a variety of code snippets to be used by developers to interact with resources that belong to VPC on Classic IaaS Offering. This repository is intended to be used for documentation purposes only and not to be included as a dependency.

These examples are provided in the following languages.

1. [Go](#go)
2. [Python](#python)

These examples will walk you through the following steps.

1. Retrieve API key for your account.
2. Get an IAM access token using your api key.
3. Get a list of the resources.
4. Post a resource.

## Go

1. Retrieve API key for your account.

2. Get an IAM access token using your api key.

3. Get a list of the resources.

4. Post a resource.

## Python

The following steps gives an example on how to retrieve a token, list all VPCs, and create a VPC.

1. Retrieve API key for your account.

2. Get an IAM access token using your API key.

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

3. Get a list of all VPCs.

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

version = "2019-01-01"

payload = ""

try:
    # Connect to rias endpoint for vpcs
    conn.request("GET", "/v1/vpcs?version=" + version, payload, headers)

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

4. Create a VPC.

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

version = "2019-01-01"

# Required payload for creating a VPC
payload = f'{{"name": "NAME_OF_VPC"}}'

try:
    # Connect to rias endpoint for vpcs
    conn.request("POST", "/v1/vpcs?version=" + version, payload, headers)

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