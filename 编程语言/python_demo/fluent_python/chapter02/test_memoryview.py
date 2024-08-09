from array import array

octets = array("B", range(6))
m1 = memoryview(octets)
print(m1.tolist())

m2 = m1.cast("B", [2, 3])
print(m2.tolist())

m3 = m1.cast("B", [3, 2])
print(m3.tolist())

m2[1, 1] = 22
m3[1, 1] = 33
print(octets)


numbers = array("h", [-2, -1, 0, 1, 2])
memv = memoryview(numbers)
print(len(memv))
print(memv[0])
memv_oct = memv.cast("B")
print(memv_oct.tolist())
memv_oct[5] = 4
print(numbers)
