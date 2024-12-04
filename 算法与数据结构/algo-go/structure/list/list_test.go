package list

import (
	"fmt"
	"testing"
)

func TestSingle(t *testing.T) {
	list := Single[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(5)

	fmt.Println(list.String())

	list.Remove(1)
	fmt.Println(list.String())

	list.Remove(5)
	fmt.Println(list.String())

	list.Add(6)
	fmt.Println(list.String())

	list.RemoveAt(1)
	fmt.Println(list.String())

	list.RemoveAt(3)
	fmt.Println(list.String())

	list.Add(10)
	fmt.Println(list.String())

	fmt.Println(list.Size())

	list.Reverse()
	fmt.Println(list.String())
}

func TestDouble(t *testing.T) {
	list := Double[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(5)

	fmt.Println(list.String())

	list.Remove(1)
	fmt.Println(list.String())

	list.Remove(5)
	fmt.Println(list.String())

	list.Add(6)
	fmt.Println(list.String())

	list.Add(10)
	fmt.Println(list.String())

	fmt.Println(list.Size())

	list.Reverse()
	fmt.Println(list.String())
}

func TestArray(t *testing.T) {
	list := Array[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(5)

	fmt.Println(list.String())

	list.Remove(1)
	fmt.Println(list.String())

	list.Remove(5)
	fmt.Println(list.String())

	list.Add(6)
	fmt.Println(list.String())

	list.RemoveAt(1)
	fmt.Println(list.String())

	list.RemoveAt(3)
	fmt.Println(list.String())

	list.Add(10)
	fmt.Println(list.String())

	fmt.Println(list.Size())

	list.Reverse()
	fmt.Println(list.String())
}

func TestDefault(t *testing.T) {
	list := Default[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(5)

	fmt.Println(list.String())

	list.Remove(1)
	fmt.Println(list.String())

	list.Remove(5)
	fmt.Println(list.String())

	list.Add(6)
	fmt.Println(list.String())

	list.RemoveAt(1)
	fmt.Println(list.String())

	list.RemoveAt(3)
	fmt.Println(list.String())

	list.Add(10)
	fmt.Println(list.String())

	fmt.Println(list.Size())

	list.Reverse()
	fmt.Println(list.String())
}
