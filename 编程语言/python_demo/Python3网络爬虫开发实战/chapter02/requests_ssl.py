import requests
import logging
from requests.packages import urllib3

r = requests.get("https://ssr2.scrape.center/", verify=False)
print(r.status_code)

logging.captureWarnings(True)
r = requests.get("https://ssr2.scrape.center/", verify=False)
print(r.status_code)


urllib3.disable_warnings()
r = requests.get("https://ssr2.scrape.center/", verify=False)
print(r.status_code)


# r = requests.get("https://ssr2.scrape.center/", cert=("server.crt", "server.key"))
# print(r.status_code)
