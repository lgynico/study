import pymongo
from bson.objectid import ObjectId

# client = pymongo.MongoClient(host="localhost", port=27017)
client = pymongo.MongoClient("mongodb://localhost:27017")

db = client.test
# db = client["test"]

collection = db.students
# collection = db["students"]

student = {"id": "20170101", "name": "Jordan", "age": 20, "gender": "male"}
# result = collection.insert(student)
result = collection.insert_one(student)
print("insert one sucess:", result, result.inserted_id)


student1 = {"id": "20170102", "name": "Nico", "age": 30, "gender": "male"}
student2 = {"id": "20170103", "name": "XiaoMei", "age": 18, "gender": "female"}
# resule = collection.insert([student1, student2])
result = collection.insert_many([student1, student2])
print("insert many success:", result, result.inserted_ids)


result = collection.find_one({"name": "Nico"})
print(type(result))
print(result)

result = collection.find_one({"_id": ObjectId("670906f41f0ebc51392b03f1")})
print(result)


results = collection.find({"age": {"$gt": 20}})
print(results)
for result  in results:
    print(result)


# count  = collection.find().count()
count = collection.count_documents({})
print(count)


results = collection.find().sort("name", pymongo.ASCENDING)
print([result["name"]] for result in results)


results = collection.find().sort("name", pymongo.ASCENDING).skip(2).limit(2)
print([result["name"]] for result in results)



condition = {"name": "Nico"}
student = collection.find_one(condition)
student["age"] = 18
# result = collection.update(condition, student)
# result = collection.update(condition, {"$set": student})
result = collection.update_one(condition, {"$set": student})
print(result, result.matched_count, result.modified_count)


# result = collection.remove({"name": "Nico"})
result = collection.delete_one({"name": "Nico"})
print(result, result.deleted_count)

result = collection.delete_many({"age": {"$gt": 0}})
print(result, result.deleted_count)