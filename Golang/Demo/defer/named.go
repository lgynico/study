package main

import "fmt"

func returnDefer() (i int) {
	defer func() {
		i += 10
	}()

	return 10
}

func main() {
	fmt.Println(returnDefer())
}
