package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer A")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("panic")
}
