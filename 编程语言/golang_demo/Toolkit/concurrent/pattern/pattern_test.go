package pattern

import (
	"fmt"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	var (
		jobCount    = 10
		workerCount = 5
		jobs        = make(chan int, 8)
		results     = make(chan int, 8)
	)

	for w := 1; w <= workerCount; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= jobCount; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= jobCount; a++ {
		<-results
	}
}

func TestFadeInFadeOut(t *testing.T) {
	gen := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}

	in := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// FadeOut
	c1 := sq(in)
	c2 := sq(in)

	// FadeIn
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func TestPipeline(t *testing.T) {
	pipe := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			pipe <- i
		}
		close(pipe)
	}()

	out := pipeline(pipe)
	for n := range out {
		fmt.Println(n)
	}
}

func TestContext(t *testing.T) {
	cancelByContext()
}

func TestErrGroup(t *testing.T) {
	errGroup()
}
