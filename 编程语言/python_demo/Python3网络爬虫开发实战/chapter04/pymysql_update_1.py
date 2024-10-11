import pymysql

db = pymysql.connect(host="localhost", port=3306, user="root", password="123456", db="spiders")
cursor = db.cursor()
sql = "UPDATE students SET age = %s WHERE name = %s"
try:
    cursor.execute(sql, (25, "Bob"))
    db.commit()
    print("修改数据成功")
except:
    db.rollback()
    print("修改数据失败")

db.close()