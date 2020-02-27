import http.client
import json
from config import conn, headers, version, print_json


# Fetch all VPCs
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/VPCs/list_vpcs
def fetch_vpcs():

    payload = ""

    try:
        # Connect to api endpoint for vpcs
        conn.request("GET", "/v1/vpcs?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    except Exception as error:
        print(f"Error fetching VPCs. {error}")
        raise



# Create VPC
# Spec:https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/VPCs/create_vpc
# Params:
#   name: string (required)
def create_vpc(name):

    # Required payload for creating a VPC
    payload = f'{{"name": "{name}"}}'

    try:
        # Connect to api endpoint for vpcs
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


fetch_vpcs()
create_vpc("VPC_NAME")
