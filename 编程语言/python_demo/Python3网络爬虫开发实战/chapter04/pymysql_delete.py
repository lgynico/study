import pymysql

table = "students"
condition = "age > 20"
sql = "DELETE FROM {table} WHERE {condition}".format(table=table, condition=condition)

db = pymysql.connect(host="localhost", port=3306, user="root", password="123456", db="spiders")
cursor = db.cursor()
try:
    cursor.execute(sql)
    db.commit()
    print("删除数据成功")
except:
    db.rollback()
    print("删除数据失败")

db.close()
