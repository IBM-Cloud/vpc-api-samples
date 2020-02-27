import http.client
import json
from config import conn, headers, version, print_json


# Fetch all SSH Keys
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/SSH%20Keys/list_keys
def fetch_ssk_keys():

    payload = ""

    try:
        # Connect to api endpoint for ssh keys
        conn.request("GET", "/v1/keys?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    # If an error happens while fetching ssh keys
    except Exception as error:
        print(f"Error fetching SSH keys. {error}")
        raise


# Create SSH Key
# Spec: https://pages.github.ibm.com/riaas/api-spec/spec_2019-05-07/#/SSH%20Keys/create_key
# Params:
#   name: str
#   public_key: str (required)
#   rsa: str
def createSSHKey(name, public_key, key_type):

    # Required payload for creating an ssh key
    payload = f'''
        {{
            "name": "{name}",
            "public_key": "{public_key}",
            "type": "{key_type}"
        }}
    '''

    try:
        # Connect to api endpoint for ssh keys
        conn.request("POST", "/v1/keys?version=" + version, payload, headers)

        # Get and read response data
        res = conn.getresponse()
        data = res.read()

        # Print and return response data
        print_json(data.decode("utf-8"))
        return data.decode("utf-8")

    # If an error happens while creating an ssh key
    except Exception as error:
        print(f"Error creating SSH key. {error}")
        raise


fetch_ssk_keys()
createSSHKey("SSH_KEY_NAME", "PUBLIC_KEY", "KEY_TYPE")
