package main

import "fmt"

type (
	Empty interface{}

	EmptyImpl struct{}

	NonEmpty interface {
		F()
	}

	NonEmptyImpl struct{}
)

func (p *NonEmptyImpl) F() {}

func EmptyData() Empty {
	var e *EmptyImpl
	return e
}

func NonEmptyData() NonEmpty {
	return nil
}

func main() {
	var i = EmptyData()
	if i == nil {
		fmt.Println("i == nil")
	} else {
		fmt.Println("i != nil")
	}

	var j = EmptyData()
	if j == nil {
		fmt.Println("j == nil")
	} else {
		fmt.Println("j != nil")
	}

	var k Empty
	if k == nil {
		fmt.Println("k == nil")
	} else {
		fmt.Println("k != nil")
	}
}
