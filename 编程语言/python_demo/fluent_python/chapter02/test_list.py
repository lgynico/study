symbols = "$¢£¥€¤"
codes = []
for symbol in symbols:
    codes.append(ord(symbol))

print(codes)


codes2 = [ord(symbol) for symbol in symbols]
print(codes2)


x = "ABC"
codes3 = [ord(x) for x in x]
print(x)
print(codes3)

codes3 = [last := ord(c) for c in x]
print(last)
# print(c) # NameError: name 'c' is not defined


beyond_ascii = [ord(s) for s in symbols if ord(s) > 127]
print(beyond_ascii)
beyond_ascii2 = list(filter(lambda c: c > 127, map(ord, symbols)))
print(beyond_ascii2)


colors = ["black", "white"]
sizes = ["S", "M", "L"]
tshirts = [(color, size) for color in colors for size in sizes]
print(tshirts)

for color in colors:
    for size in sizes:
        print((color, size))
