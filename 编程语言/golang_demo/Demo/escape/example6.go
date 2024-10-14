package main

func foo(a *int) {
	return
}

func main() {
	data := 10
	f := foo
	f(&data)
	println(data)
}
