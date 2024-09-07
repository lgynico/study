import httpx
import requests


response = httpx.get("https://www.httpbin.org/get")
print(response.status_code)
print(response.headers)
print(response.text)


url = "https://spa16.scrape.center/"
try:
    response = requests.get(url)
    print(response.text)
except requests.exceptions.ConnectionError as e:
    print("requests.get:", e)


try:
    response = httpx.get(url)
    print(response.text)
except httpx.RemoteProtocolError as e:
    print("httpx.get:", e)


client = httpx.Client(http2=True)
response = client.get(url)
print(response.text)
