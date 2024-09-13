import json


str = """
[{
    "name": "Bob",
    "gender": "male",
    "birthday": "1992-10-18"
},{
    "name": "Selina",
    "gender": "female",
    "birthday": "1995-10-18"
}]
"""
print(type(str))
data = json.loads(str)
print(data)
print(type(data))


print(data[0]["name"])
print(data[0].get("name"))


print(data[0].get("age"))
print(data[0].get("age", 25))


with open("data.json", encoding="utf-8") as f:
    str = f.read()
    data = json.loads(str)
    print(data)

data = json.load(open("data.json", encoding="utf-8"))
print(data)


data1 = [
    {"name": "Bob", "gender": "male", "birthday": "1992-10-28"},
    {"name": "张三", "gender": "男", "birthday": "2000-10-01"},
]
with open("data1.json", "w", encoding="utf-8") as f:
    f.write(json.dumps(data1, indent=4, ensure_ascii=False))
json.dump(
    data1, open("data2.json", "w", encoding="utf-8"), indent=4, ensure_ascii=False
)
