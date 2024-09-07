import urllib.parse
import urllib.request
import urllib.error
import socket


data = bytes(urllib.parse.urlencode({"name": "nico"}), encoding="utf-8")
response = urllib.request.urlopen("https://www.httpbin.org/post", data=data)
print(response.read().decode("utf-8"))


try:
    response = urllib.request.urlopen("https://www.httpbin.org/get", timeout=0.1)
except urllib.error.URLError as e:
    if isinstance(e.reason, socket.timeout):
        print("TIMEOUT!")
