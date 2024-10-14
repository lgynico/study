package main

import (
	"log"
	"runtime"
	"time"
)

func test() {
	s := make([]int, 8)
	log.Println("loop begin")
	for i := 0; i < 32*1000*1000; i++ {
		s = append(s, i)
	}
	log.Println("loop end")
}

func main() {
	log.Println("Start")

	test()

	log.Println("force GC")
	runtime.GC()
	log.Println("Done")

	time.Sleep(3600 * time.Second)
}
