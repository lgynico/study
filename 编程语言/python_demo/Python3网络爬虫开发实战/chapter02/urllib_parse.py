from urllib.parse import (
    urlparse,
    urlunparse,
    urlsplit,
    urlunsplit,
    urljoin,
    urlencode,
    parse_qs,
    parse_qsl,
    quote,
    unquote,
)


result = urlparse("https://www.baidu.com/index.html;user?id=5#comment")
print(result)


result = urlparse("https://www.baidu.com/index.html;user?id=5#comment", scheme="https")
print(result)


result = urlparse("http://www.baidu.com/index.html;user?id=5#comment", scheme="https")
print(result)


result = urlparse(
    "https://www.baidu.com/index.html;user?id=5#comment", allow_fragments=False
)
print(result)


result = urlparse("https://www.baidu.com/index.html#comment", allow_fragments=False)
print(result)
print(result.scheme, result[0], result.netloc, result[1], sep="\n")


data = ["https", "www.baidu.com", "index.html", "user", "a=6", "comment"]
print(urlunparse(data))


result = urlsplit("https://www.baidu.com/index.html;user?id=5#comment")
print(result)
print(result.scheme, result[0])


data = ["https", "www.baidu.com", "index.html", "a=6", "comment"]
print(urlunsplit(data))  # 只能 5 个长度


print(urljoin("https://www.baidu.com", "FAQ.html"))
print(urljoin("https://www.baidu.com", "https://cuiqingcai.com/FAQ.html"))
print(urljoin("https://www.baidu.com/about.html", "https://cuiqingcai.com/FAQ.html"))
print(
    urljoin(
        "https://www.baidu.com/about.html", "https://cuiqingcai.com/FAQ.html?question=2"
    )
)
print(urljoin("https://www.baidu.com?wd=abc", "https://cuiqingcai.com/index.php"))
print(urljoin("https://www.baidu.com", "?category=2#comment"))
print(urljoin("www.baidu.com", "?category=2#comment"))
print(urljoin("www.baidu.com#comment", "?category=2"))


params = {"name": "nico", "age": "30"}
base_url = "https://www.baidu.com?"
url = base_url + urlencode(params)
print(url)


query = "name=nico&age=30"
print(parse_qs(query))
print(parse_qsl(query))


keyword = "游戏"
url = "https://www.baidu.com/s?wd=" + quote(keyword)
print(url)
print(unquote(url))
