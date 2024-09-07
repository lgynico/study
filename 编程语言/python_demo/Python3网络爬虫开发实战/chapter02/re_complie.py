import re


content1 = "2024-09-05 12:00"
content2 = "1992-10-01 02:16"
content3 = "2015-09-01 08:00"

pattern = re.compile(r"\d{2}:\d{2}")

print(re.sub(pattern, "", content1))
print(re.sub(pattern, "", content2))
print(re.sub(pattern, "", content3))
