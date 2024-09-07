import urllib.parse
import urllib.request


request = urllib.request.Request("https://python.org")
response = urllib.request.urlopen(request)
print(response.read().decode("utf-8"))


url = "https://www.httpbin.org/post"
headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
    "Host": "www.httpbin.org",
}
dict = {"name": "nico"}
data = bytes(urllib.parse.urlencode(dict), encoding="utf-8")
request = urllib.request.Request(url=url, data=data, headers=headers, method="POST")
response = urllib.request.urlopen(request)
print(response.read().decode("utf-8"))
