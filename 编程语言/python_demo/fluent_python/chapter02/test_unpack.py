lax_coordinates = (33.9425, -118.408056)
latitude, longitude = lax_coordinates
print(latitude)
print(longitude)


a = 1
b = 2
b, a = a, b
print(a, b)


print(divmod(20, 8))
t = (20, 8)
print(divmod(*t))

quotient, remainder = divmod(*t)
print(quotient, remainder)


import os

_, filename = os.path.split("/home/nico/.ssh/id_rsa.pub")
print(filename)


a, b, *rest = range(5)
print(a, b, rest)

a, b, *rest = range(3)
print(a, b, rest)

a, b, *rest = range(2)
print(a, b, rest)


a, *body, c, d = range(5)
print(a, body, c, d)
*head, b, c, d = range(5)
print(head, b, c, d)


def fun(a, b, c, d, *rest):
    return a, b, c, d, rest


print(fun(*[1, 2], 3, *range(4, 7)))


print((*range(4), 4))
print([*range(4), 4])
print({*range(4), 4, *(5, 6, 7)})


metro_areas = [
    ("Tokyo", "JP", 36.933, (35.689722, 139.691667)),
    ("Delhi NCR", "IN", 21.935, (28.613889, 77.208889)),
    ("Mexico City", "MX", 20.142, (19.433333, -99.133333)),
    ("New York-Newark", "US", 20.104, (40.808611, -74.020386)),
    ("São Paulo", "BR", 19.649, (-23.547778, -46.635833)),
]
print(f'{"":15} | {"latitude":>9} | {"longitude":>9}')
for name, _, _, (lat, lon) in metro_areas:
    if lon <= 0:
        print(f"{name:15} | {lat:9.4f} | {lon:9.4f}")
