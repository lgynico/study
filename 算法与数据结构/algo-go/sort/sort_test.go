package sort

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"testing"

	"github.com/lgynico/algo-go/utils/arrays"
	"github.com/lgynico/algo-go/utils/rands"
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

func TestMerge(t *testing.T) {
	// testSort(Merge)
	testSort(MergeAdvance)
}

func TestQuick(t *testing.T) {
	// testSort(QuickV1)
	// testSort(Quick)
	testSort(Quick)
}

func TestNetherlandFlag(t *testing.T) {
	var (
		arr    = arrays.RandomArray(0, 10, 10)
		x      = rands.Random(0, 10)
		p      = netherlandFlag1(arr, x)
		pl, pr = netherlandFlag2(arr, x)
	)

	fmt.Println(arr)
	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(pl, pr)
}

func testSort(sortFunc func([]int)) {
	log.Println("test begin")

	var (
		testCount = 10000
		maxSize   = 100
		max       = 1000
	)

	for i := 0; i < testCount; i++ {
		var (
			size = rand.Intn(maxSize)
			arr  = arrays.RandomArray(0, max, size)
			arr1 = arrays.CopyArray(arr)
			arr2 = arrays.CopyArray(arr)
		)

		sortFunc(arr1)
		sort.Ints(arr2)

		if !arrays.Equals(arr1, arr2) {
			log.Printf("%v\n", arr)
			log.Printf("%v\n", arr1)
			log.Printf("%v\n", arr2)
			log.Fatalln("test error")
			return
		}
	}

	log.Println("test success")
}
