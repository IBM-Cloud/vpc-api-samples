import http.client
import json
from config import conn, headers, version, print_json


# Fetch all Security Groups
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/Security%20Groups/list_security_groups
def fetch_security_groups():

    payload = ""

    try:
        # Connect to api endpoint for security groups
        conn.request("GET", "/v1/security_groups?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    # If an error happens while fetching security groups
    except Exception as error:
        print(f"Error fetching security groups. {error}")
        raise


# Create Security Group
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/Security%20Groups/create_security_group
# Params:
#   name: str
#   rules: [{
#       protocol: str (required)
#       direction: str (required)
#   }]
#   vpc: str (required)
def create_security_group(name, protocol, direction, vpc_id):

    # Required payload for creating a security group
    payload = f'''
        {{
            "name": "{name}",
            "rules": [{{
                "protocol": "{protocol}",
                "direction": "{direction}"
            }}],
            "vpc": {{
                "id": "{vpc_id}"
            }}
        }}
    '''

    try:
        # Connect to api endpoint for creating a security group
        conn.request("POST", "/v1/security_groups?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    # If an error happens while creating a security group
    except Exception as error:
        print(f"Error creating a security group. {error}")
        raise


fetch_security_groups()
create_security_group("SECURITY_GROUP_NAME", "PROTOCOL", "DIRECTION", "VPC_ID")
