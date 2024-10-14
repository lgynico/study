package main

func foo(val int) *int {
	var (
		val1 *int = new(int)
		val2 *int = new(int)
		val3 *int = new(int)
		val4 *int = new(int)
		val5 *int = new(int)
	)

	for i := 0; i < 5; i++ {
		println(&val, val1, val2, val3, val4, val5)
	}

	return val3
}

func main() {
	val := foo(666)
	println(*val, val)
}
