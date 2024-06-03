package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Defer A")
	}()

	defer func() {
		fmt.Println("Defer B")
	}()

	defer func() {
		fmt.Println("Defer C")
	}()
}
