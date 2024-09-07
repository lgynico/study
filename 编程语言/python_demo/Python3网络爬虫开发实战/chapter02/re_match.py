import re


content = "Hello 123 4567 World_This is a Regex Demo"
print(len(content))

result = re.match(r"^Hello\s\d\d\d\s\d{4}\s\w{10}", content)
print(result)
print(result.group())
print(result.span())


content = "Hello 1234567 World_This is a Regex Demo"
result = re.match(r"^Hello\s(\d+)\sWorld", content)
print(result)
print(result.group())
print(result.group(1))
print(result.span())


content = "Hello 123 4567 World_This is a Regex Demo"
result = re.match(r"^Hello.*Demo$", content)
print(result)
print(result.group())
print(result.span())


content = "Hello 1234567 World_This is a Regex Demo"
result = re.match(r"^He.*(\d+).*Demo$", content)
print(result)
print(result.group(1))


result = re.match(r"^He.*?(\d+).*Demo$", content)
print(result)
print(result.group(1))


content = "http://weibo.com/comment/kEraCN"
result1 = re.match(r"http.*?comment/(.*?)", content)
result2 = re.match(r"http.*?comment/(.*)", content)
print("result1", result1.group(1))
print("result2", result2.group(1))


content = """Hello 1234567 World_This
is a Regex Demo
"""
result = re.match(r"^He.*?(\d+).*?Demo$", content, re.S)
print(result.group(1))


content = "(百度) www.baidu.com"
result = re.match("\(百度\) www\.baidu\.com", content)
print(result)
