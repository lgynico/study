import requests
from requests.auth import HTTPBasicAuth
from requests_oauthlib import OAuth1


r = requests.get("https://ssr3.scrape.center/", auth=HTTPBasicAuth("admin", "admin"))
# r = requests.get("https://ssr3.scrape.center/", auth=("admin", "admin"))
print(r.status_code)


url = "https://api.twitter.com/1.1/account/verify_credentials.json"
auth = OAuth1(
    "YOUR_APP_KEY", "YOUR_APP_SECRET", "USER_OAUTH_TOKEN", "USER_OAUTH_TOKEN_SECRET"
)
requests.get(url, auth=auth, timeout=(5, 30))
