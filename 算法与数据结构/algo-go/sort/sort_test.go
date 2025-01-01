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
	testSort(Insertion, 10000, 1000, 100)
}

func TestSelection(t *testing.T) {
	testSort(Selection, 10000, 1000, 100)
}

func TestBubble(t *testing.T) {
	testSort(Bubble, 1000, 1090, 100)
	// Bubble([]int{0, 2, -8, 8, 1})
}

func TestMerge(t *testing.T) {
	// testSort(Merge)
	testSort(MergeAdvance, 10000, 1000, 100)
}

func TestQuick(t *testing.T) {
	// testSort(QuickV1)
	// testSort(Quick)
	testSort(Quick, 10000, 1000, 100)
}

func TestHeap(t *testing.T) {
	testSort(Heap, 10000, 1000, 100)
}

func TestCount(t *testing.T) {
	testSort(Count, 10000, 1000, 10)
}

func TestBucket(t *testing.T) {
	testSort(Bucket, 10000, 1000, 100)
	testSort(BucketAdvance, 10000, 1000, 100)
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

func testSort(sortFunc func([]int), testCount, maxSize, max int) {
	log.Println("test begin")

	for i := 0; i < testCount; i++ {
		var (
			size = rand.Intn(maxSize)
			arr  = arrays.RandomArray(-10, max, size)
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
