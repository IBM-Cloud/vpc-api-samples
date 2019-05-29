import http.client
import gzip
import json
from apikey import APIKEY

def get_token():
    # URL for token
    conn = http.client.HTTPSConnection("iam.cloud.ibm.com")

    # Payload for retrieving token. Note: An API key will need to be generated and replaced here
    payload = f'grant_type=urn%3Aibm%3Aparams%3Aoauth%3Agrant-type%3Aapikey&apikey={APIKEY}&response_type=cloud_iam'

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
