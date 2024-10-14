package main

import "fmt"

func deferCall() {
	defer func() { fmt.Println("defer A before panic") }()
	defer func() { fmt.Println("defer B before panic") }()

	panic("panic")

	defer func() { fmt.Println("defer C after panic") }()
}

func main() {
	deferCall()

	fmt.Println("main end")
}
