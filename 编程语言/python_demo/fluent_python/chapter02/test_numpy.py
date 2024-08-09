import numpy
from time import perf_counter as pc

a = numpy.arange(12)
print(a)
print(type(a))
print(a.shape)

a.shape = 3, 4
print(a)
print(a[2])
print(a[2, 1])
print(a[:, 1])
print(a.transpose())


floats = numpy.loadtxt("floats-10M-lines.txt")
print(floats[-3:])
floats *= 0.5
print(floats[-3:])

t0 = pc()
floats /= 3
print(pc() - t0)

numpy.save("floats-10M", floats)
floats2 = numpy.load("floats-10M.npy", "r+")
floats2 *= 6
print(floats2[-3:])
