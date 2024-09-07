import requests


data = {"name": "nico", "age": 30}
r = requests.post("https://www.httpbin.org/post", data=data)
print(r.text)
