package main

import "fmt"

func function(a, b int) int {
	fmt.Println(a, b)
	return a + b
}

func main() {
	defer function(1, function(2, 3))
	defer function(4, function(5, 6))
}
