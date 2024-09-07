import httpx

with httpx.Client(http2=True) as client:
    response = client.get("https://www.httpbin.org/get")
    print(response)
    print(response.http_version)


url = "http://www.httpbin.org/headers"
headers = {"User-Agent": "my-app/0.0.1"}
with httpx.Client(headers=headers) as client:
    r = client.get(url)
    print(r.json()["headers"]["User-Agent"])
