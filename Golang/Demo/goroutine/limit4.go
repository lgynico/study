package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func busi(ch chan int) {
	for t := range ch {
		fmt.Println("go task ", t, ", goroutine count = ", runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(task int, ch chan int) {
	wg.Add(1)
	ch <- task
}

func main() {
	// 无缓冲 channel 分离任务
	ch := make(chan int)
	goCnt := 3
	for i := 0; i < goCnt; i++ {
		go busi(ch)
	}

	taskCnt := math.MaxInt64
	for t := 0; t < taskCnt; t++ {
		sendTask(t, ch)
	}

	wg.Wait()
}
