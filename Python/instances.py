import http.client
import json
from config import conn, headers, version, print_json


# Fetch all Instances
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/Instances/list_instances
def fetch_instances():

    payload = ""

    try:
        # Connect to api endpoint for instances
        conn.request("GET", "/v1/instances?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    # If an error happens while fetching instances
    except Exception as error:
        print(f"Error fetching instances. {error}")
        raise



# Create Instance
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/Instances/create_instance
# Params:
#   image: {
#       id: str (required)
#   }
#   name: str
#   profile: {
#       name: str (required)
#   }
#   keys: [{
#       id: str (required)
#   }]
#   primary_network_interface {
#       port_speed: int (required),
#       subnet: {
#           id: str (required)
#       }
#   }
#   vpc: {
#       id: str (required)
#   }
#   zone: {
#       name: str (required)
#   }
def create_instance(image, name, profile, key_id, port_speed, subnet_id, vpc_id, zone):

    # Required payload for creating an instance
    payload = f'''
        {{
            "image": {{
                "id": "{image}"
            }},
            "name": "{name}",
            "profile": {{
                "name": "{profile}"
            }},
            "keys": [{{
                "id": "{key_id}"
            }}],
            "primary_network_interface": {{
                "port_speed": {port_speed},
                "subnet": {{
                    "id": "{subnet_id}"
                }}
            }},
            "vpc": {{
                "id": "{vpc_id}"
            }},
            "zone": {{
                "name": "{zone}"
            }}
        }}
    '''

    try:
        # Connect to api endpoint for instances
        conn.request("POST", "/v1/instances?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    # If an error happens while creating an instance
    except Exception as error:
        print(f"Error creating instance. {error}")
        raise


fetch_instances()
create_instance("IMAGE_ID", "INSTANCE_NAME", "PROFILE", "KEY_ID", "PORT_SPEED(int)", "SUBNET_ID", "VPC_ID", "ZONE")
