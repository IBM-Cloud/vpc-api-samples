import http.client
import json
from token import get_token

region = "us-south"

conn = http.client.HTTPSConnection(f"{region}.iaas.cloud.ibm.com")

headers = {
    'Content-Type': 'application/json',
    'Cache-Control': 'no-cache',
    'Accept': 'application/json',
    'Authorization': get_token(),
    'cache-control': 'no-cache'
}

version = "2019-06-01"

def print_json(data):
    print(json.dumps(json.loads(data), indent=2, sort_keys=True))
