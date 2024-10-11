import pymysql

db = pymysql.connect(host="localhost", port=3306, user="root", password="123456", db="spiders")
cursor = db.cursor()
sql = "SELECT * FROM students WHERE age >= 20"

try:
    cursor.execute(sql)
    print("Count:", cursor.rowcount)
    row = cursor.fetchone()
    while row:
        print("Row:", row)
        row = cursor.fetchone()
except:
    print("Error")

db.close()