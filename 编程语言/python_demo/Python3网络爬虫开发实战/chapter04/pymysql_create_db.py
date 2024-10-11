import pymysql

db = pymysql.connect(host="localhost", user="root", passwd="123456", port=3306)
cursor = db.cursor()
cursor.execute("SELECT version()")
data = cursor.fetchone()
print("Database version:", data)
cursor.execute("CREATE DATABASE spiders DEFAULT CHARACTER SET utf8mb4")
db.close()