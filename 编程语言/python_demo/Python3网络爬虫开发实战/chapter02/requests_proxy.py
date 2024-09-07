import requests

proxies = {
    "http": "http://127.0.0.1:8080",
    "https": "https://user:password@127.0.0.1:8080",
    "https": "sock5://user:password@host:port",
}
requests.get("https://www.httpbin.org/get", proxies=proxies)
