package pattern

import (
	"fmt"
	"math/rand"
	"time"
)

// worker pool 模式
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d started job: %d\n", id, j)
		time.Sleep(time.Second * time.Duration(rand.Intn(3)))
		results <- j * 2
		fmt.Printf("worker %d finished job: %d\n", id, j)
	}
}
