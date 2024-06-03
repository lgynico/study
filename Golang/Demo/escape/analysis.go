package main

func foo(val int) *int {
	var (
		val1 int = 11
		val2 int = 12
		val3 int = 13
		val4 int = 14
		val5 int = 15
	)

	// 说是为了防止 foo 被内联，但是测试不出来
	for i := 0; i < 5; i++ {
		println(&val, &val1, &val2, &val3, &val4, &val5)
	}

	return &val3
}

func main() {
	val := foo(666)
	println(*val, val)
}
