package tree

import (
	"fmt"
	"testing"
)

func TestPrefix(t *testing.T) {
	prefix := NewPrefix()
	prefix.Insert("nico")
	prefix.Insert("Apple")
	prefix.Insert("about")
	prefix.Insert("banana")
	prefix.Insert("kerNel")

	fmt.Println(prefix.Contains("nico"))
	fmt.Println(prefix.ContainsPrefix("abo"))

	fmt.Println(prefix.Contains("apple"))
	fmt.Println(prefix.Contains("Apple"))

	prefix.Remove("Apple")
	fmt.Println(prefix.Contains("Apple"))
}
