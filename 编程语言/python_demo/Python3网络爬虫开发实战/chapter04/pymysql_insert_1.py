import pymysql

id = "20120001"
user = "Bob"
age = 20

db = pymysql.connect(host="localhost", port=3306, user="root", password="123456", db="spiders")
cursor = db.cursor()
sql = "INSERT INTO students (id, name, age) VALUES(%s, %s, %s)"
try:
    cursor.execute(sql, (id, user, age))
    db.commit()
    print("插入数据成功")
except:
    db.rollback()
    print("插入数据失败")
db.close()