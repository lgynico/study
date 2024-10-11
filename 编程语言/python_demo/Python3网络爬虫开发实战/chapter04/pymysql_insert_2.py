import pymysql

data = {
    "id": "20120002",
    "user": "Nico",
    "age": 30
}
table = "students"
keys = ", ".join(data.keys())
values = ", ".join(["%s"] * len(data))
sql = "INSERT INTO {table} ({keys}) VALUES ({values})".format(table=table, keys=keys, values=values)

print("keys:", keys)
print("values:", values)
print("sql:", sql)

db = pymysql.connect(host="localhost", port=3306, user="root", password="123456", db="spiders")
cursor = db.cursor()
sql = "INSERT INTO students (id, name, age) VALUES(%s, %s, %s)"
try:
    if cursor.execute(sql, tuple(data.values())):
        print("插入数据成功")
        db.commit()
except:
    db.rollback()
    print("插入数据失败")
db.close()