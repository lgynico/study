import requests
import re


r = requests.get("https://www.httpbin.org/get")
print(r.text)

r = requests.get("https://www.httpbin.org/get?name=nico&age=30")
print(r.text)


data = {"name": "nico", "age": 30}
r = requests.get("https://www.httpbin.org/get", params=data)
print(r.text)
print(type(r.text))
print(r.json())
print(type(r.json()))


headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"
}
r = requests.get("https://ssr1.scrape.center/", headers=headers)
pattern = re.compile("<h2.*?>(.*?)</h2>", re.S)
titles = re.findall(pattern, r.text)
print(titles)


r = requests.get("https://scrape.center/favicon.ico")
print(r.text)
print(r.content)
with open("favicon.ico", "wb") as f:
    f.write(r.content)
