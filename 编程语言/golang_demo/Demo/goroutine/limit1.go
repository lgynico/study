package main

import (
	"fmt"
	"math"
	"runtime"
)

func busi(ch chan bool, i int) {
	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch
}

func main() {
	task_cnt := math.MaxInt64
	// task_cnt := 10
	ch := make(chan bool, 3)

	for i := 0; i < task_cnt; i++ {
		ch <- true
		go busi(ch, i)
	}

	// 这里要阻塞等待
}
