import http.client
import json
from config import conn, headers, version, print_json


# Fetch all Profiles
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_scoped/#/Instances/list_instance_profiles
def fetch_profiles():

    payload = ""

    try:
        # Connect to rias endpoint for profiles
        conn.request("GET", "/v1/instance/profiles?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    except Exception as error:
        print(f"Error fetching Profiles. {error}")
        raise


fetch_profiles()
