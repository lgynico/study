package main

import "fmt"

func deferFunc() {
	fmt.Println("Defer func")
}

func returnFunc() int {
	fmt.Println("Return func")
	return 0
}

func deferReturnFunc() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	deferReturnFunc()
}
