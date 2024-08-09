l = [1, 2, 3]
print(l)
print(l * 5)
print(5 * "abcd")


board = [["_"] * 3 for _ in range(3)]
print(board)
board[1][2] = "X"
print(board)


weird_board = [["_"] * 3] * 3
print(weird_board)
weird_board[1][2] = "O"
print(weird_board)


l = [1, 2, 3]
print(l)
print(id(l))
l *= 2
print(l)
print(id(l))
t = (1, 2, 3)
print(t)
print(id(t))
t *= 2
print(t)
print(id(t))

# list puzzle: run in shell
# t = (1, 2, [30, 40])
# t[2] += [50, 60] # TypeError: 'tuple' object does not support item assignment
# print(t) # (1, 2, [30, 40, 50, 60])
