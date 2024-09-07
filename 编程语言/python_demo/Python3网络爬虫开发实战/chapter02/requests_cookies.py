import requests
import requests.cookies

r = requests.get("https://www.baidu.com")
print(r.cookies)
for key, value in r.cookies.items():
    print(key, "=", value)


# 应该是 github 升级了认证的原因所以不行了
try:
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
        "Cookie": "_octo=GH1.1.192791326.1722825212; _device_id=d154643def0f3b6ff44aa88e670c700a; saved_user_sessions=2893568%3ANpRnoFpgXFH58dKf3n57f-owYqMixm51GoqVDBEisb2G8mv3; user_session=NpRnoFpgXFH58dKf3n57f-owYqMixm51GoqVDBEisb2G8mv3; __Host-user_session_same_site=NpRnoFpgXFH58dKf3n57f-owYqMixm51GoqVDBEisb2G8mv3; logged_in=yes; dotcom_user=lgynico; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22light%22%2C%22color_mode%22%3A%22light%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark%22%2C%22color_mode%22%3A%22dark%22%7D%7D; preferred_color_mode=light; tz=Asia%2FShanghai; _gh_sess=epIalNfc1wCkxzRt%2FgHzLfNDT%2Bd6Q60DFShgMxWtxchX1d04n5Aza9hBqC9Eug3C29jnX3vlmOKVy3l%2Bd9HrMsaKqYhxUvgXVka3C6wKJ5uTNu6l6ltqir4eB769yagxl05xhdNd5Fviw30SIgGGBBOX3Aex49lBN8NiVbqXDNWYNDBj1A%2BIFqEeI3hZHhWOoiWHnOkuzlRKaJEXcpHT6jfy%2F7Lt%2FC9Lic%2FbLCNXa6bLm0iPiz2Cm30UXZk%2B9sUjfq01UIY46PuwjYaTsfbyr8R1Jnh7YkIxJ3pp8%2FS9AqrSCPl64cgSjbU%2FxGf7GbVHDD3PMnzKba4%3D--WiCZ%2BDyaagKPkLOe--patHXXFysYBRdFK4RF97IQ%3D%3D",
    }
    r = requests.get("https://github.com/", headers=headers)
    print(r.text)
except requests.exceptions.SSLError as e:
    print(e)


try:
    cookies = "_octo=GH1.1.192791326.1722825212; _device_id=d154643def0f3b6ff44aa88e670c700a; saved_user_sessions=2893568%3ANpRnoFpgXFH58dKf3n57f-owYqMixm51GoqVDBEisb2G8mv3; user_session=NpRnoFpgXFH58dKf3n57f-owYqMixm51GoqVDBEisb2G8mv3; __Host-user_session_same_site=NpRnoFpgXFH58dKf3n57f-owYqMixm51GoqVDBEisb2G8mv3; logged_in=yes; dotcom_user=lgynico; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22light%22%2C%22color_mode%22%3A%22light%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark%22%2C%22color_mode%22%3A%22dark%22%7D%7D; preferred_color_mode=light; tz=Asia%2FShanghai; _gh_sess=epIalNfc1wCkxzRt%2FgHzLfNDT%2Bd6Q60DFShgMxWtxchX1d04n5Aza9hBqC9Eug3C29jnX3vlmOKVy3l%2Bd9HrMsaKqYhxUvgXVka3C6wKJ5uTNu6l6ltqir4eB769yagxl05xhdNd5Fviw30SIgGGBBOX3Aex49lBN8NiVbqXDNWYNDBj1A%2BIFqEeI3hZHhWOoiWHnOkuzlRKaJEXcpHT6jfy%2F7Lt%2FC9Lic%2FbLCNXa6bLm0iPiz2Cm30UXZk%2B9sUjfq01UIY46PuwjYaTsfbyr8R1Jnh7YkIxJ3pp8%2FS9AqrSCPl64cgSjbU%2FxGf7GbVHDD3PMnzKba4%3D--WiCZ%2BDyaagKPkLOe--patHXXFysYBRdFK4RF97IQ%3D%3D"
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"
    }
    jar = requests.cookies.RequestsCookieJar()
    for cookie in cookies.split(";"):
        key, value = cookie.split("=", 1)
        jar.set(key, value)
    r = requests.get("https://github.com", cookies=jar, headers=headers, verify=False)
    print(r.text)
except requests.exceptions.SSLError as e:
    print(e)
