import http.client
import json
from config import conn, headers, version, print_json

# Fetch all Floating Ips
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/Floating%20IPs/list_floating_ips
def fetch_floating_ips():

    payload = ""

    try:
        # Connect to rias endpoint for floating ips
        conn.request("GET", "/v1/floating_ips?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    # If an error happens while fetching floating ips
    except Exception as error:
        print(f"Error fetching floating ips. {error}")
        raise


# Reserve Floating IP
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/Floating%20IPs/reserve_floating_ip
# Params:
#   target: {
#       id: str (required)
#   }
def reserve_floating_ip(network_interface_id):

    # Required payload for reserving a floating ip
    payload = f'''
        {{
            "target": {{
                "id": "{network_interface_id}"
            }}
        }}
    '''

    try:
        # Connect to rias endpoint for reserving a floating ip
        conn.request("POST", "/v1/floating_ips?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()
        
        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    # If an error happens while reserving a floating ip
    except Exception as error:
        print(f"Error reserving floating ip. {error}")
        raise


fetch_floating_ips()
reserve_floating_ip("NETWORK_INTERFACE_ID")
