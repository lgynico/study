import pymysql

db = pymysql.connect(host="localhost", port=3306, user="root", password="123456", db="spiders")
cursor = db.cursor()
sql = "SELECT * FROM students WHERE age >= 20"

try:
    cursor.execute(sql)
    print("Count:", cursor.rowcount)
    one = cursor.fetchone()
    print("One:", one)
    results = cursor.fetchall()
    print("Results:", results)
    print("Results Type:", type(results))
    for row in results:
        print(row)
except:
    print("Error")

db.close()