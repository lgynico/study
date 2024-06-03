package main

import "fmt"

func main() {
	defer1()
	println()
	defer2()
}

func defer1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func defer2() {
	for i := 10; i < 20; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}
