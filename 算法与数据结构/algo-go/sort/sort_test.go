package sort

import (
	"log"
	"math/rand"
	"sort"
	"testing"

	"github.com/lgynico/algo-go/utils/arrays"
)

func TestInsertionSort(t *testing.T) {
	testSort(Insertion)
}

func TestSelection(t *testing.T) {
	testSort(Selection)
}

func TestBubble(t *testing.T) {
	testSort(Bubble)
}

func testSort(sortFunc func([]int)) {
	log.Println("test begin")

	var (
		testCount = 10000
		maxSize   = 100
		max       = 10000
	)

	for i := 0; i < testCount; i++ {
		var (
			size = rand.Intn(maxSize)
			arr1 = arrays.RandomArray(0, max, size)
			arr2 = arrays.CopyArray(arr1)
		)

		sortFunc(arr1)
		sort.Ints(arr2)

		if !arrays.Equals(arr1, arr2) {
			log.Printf("%v\n", arr1)
			log.Printf("%v\n", arr2)
			log.Fatalln("test error")
			return
		}
	}

	log.Println("test success")
}
