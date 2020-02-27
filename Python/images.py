import http.client
import json
from config import conn, headers, version, print_json


# Fetch all Images
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_scoped/#/Images/list_images
def fetch_images():

    payload = ""

    try:
        # Connect to api endpoint for images
        conn.request("GET", "/v1/images?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    except Exception as error:
        print(f"Error fetching Images. {error}")
        raise


fetch_images()
