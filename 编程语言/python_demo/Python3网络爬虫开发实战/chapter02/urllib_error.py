from urllib import request, error
import socket


try:
    response = request.urlopen("https://cuiqingcai.com/404")
except error.HTTPError as e:
    print(e.reason, e.code, e.headers, sep="\n")
except error.URLError as e:
    print(e.reason)
else:
    print("Request Successfully")


try:
    response = request.urlopen("https://www.baidu.com", timeout=0.01)
except error.URLError as e:
    print(type(e.reason))
    if isinstance(e.reason, socket.timeout):
        print("Timeout!")
