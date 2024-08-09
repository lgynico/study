p = (4, 5)
x, y = p
print("x = %d, y = %d" % (x, y))


data = ["ACME", 50, 91.1, (2012, 12, 21)]
name, shares, price, date = data
print(name, date)

name, shares, price, (year, mon, day) = data
print(name, year, mon, day)

# ValueError: not enough values to unpack (expected 3, got 2)
# x, y, z = p


s = "Hello"
a, b, c, d, e = s
print(a, b, e)


_, shares, price, _ = data
print(shares, price)
