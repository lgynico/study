import pymysql

data = {
    "id": "20120001",
    "name": "Bob",
    "age": 21
}
table = "students"
keys = ", ".join(data.keys())
values = ", ".join(["%s"] * len(data))
sql = "INSERT INTO {table} ({keys}) VALUES ({values}) ON DUPLICATE KEY UPDATE ".format(table=table, keys=keys, values=values)
update = ", ".join(["{key} = %s".format(key=key) for key in data])
sql += update

print("keys:", keys)
print("values:", values)
print("sql:", sql)

db = pymysql.connect(host="localhost", port=3306, user="root", password="123456", db="spiders")
cursor = db.cursor()
try:
    if cursor.execute(sql, tuple(data.values()) * 2):
        db.commit()
        print("插入或更新数据成功")
except:
    db.rollback()
    print("插入或更新数据失败")

db.close()