import json
import requests


data_file = 'sample_json.txt'
with open(data_file) as json_file:
    data = json.load(json_file)
json_str = json.dumps(data)
print('json data:', json.dumps(json_str))

response = requests.post("http://localhost:60730/newtask", data=json_str)
print('response:', response)
