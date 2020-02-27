import http.client
import json
from config import conn, headers, version, print_json

# Fetch all Subnets
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/Subnets/list_subnets
def fetch_subnets():

    payload = ""

    try:
        # Connect to api endpoint for subnets
        conn.request("GET", "/v1/subnets?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    # If an error happens while fetching subnets
    except Exception as error:
        print(f"Error fetching subnets. {error}")
        raise



# Create Subnet
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/Subnets/create_subnet
# Params:
#   name: str
#   vpc: str (required)
#   total_ipv4_address_count: str (required)
#   zone: str (required)
def create_subnet(name, vpc_id, total_ipv4_address_count, zone):

    # Required payload for creating a subnet
    payload = f'''
        {{
            "name": "{name}",
            "vpc": {{
                "id": "{vpc_id}"
            }},
            "total_ipv4_address_count": {total_ipv4_address_count},
            "zone": {{
                "name": "{zone}"
            }}
        }}
    '''

    try:
        # Connect to api endpoint for subnets
        conn.request("POST", "/v1/subnets?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    # If an error happens while creating a subnet
    except Exception as error:
        print(f"Error creating subnet. {error}")
        raise


fetch_subnets()
create_subnet("SUBNET_NAME", "VPC_ID", "IPV4_ADDRESS_COUNT", "ZONE")
