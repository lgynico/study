import csv


with open("data.csv", "w") as f:
    w = csv.writer(f)
    w.writerow(["id", "name", "age"])
    w.writerow(["10001", "Mike", 20])
    w.writerow(["10002", "Bob", 22])
    w.writerow(["10003", "Jordan", 21])
