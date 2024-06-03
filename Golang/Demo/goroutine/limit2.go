package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

// 使用 sync 无法控制 goroutine 的量
var wg = sync.WaitGroup{}

func busi(i int) {
	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	wg.Done()
}

func main() {
	task_cnt := math.MaxInt64

	for i := 0; i < task_cnt; i++ {
		wg.Add(1)
		go busi(1)
	}

	wg.Wait()
}
