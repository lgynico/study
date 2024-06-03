package main

import (
	"log"
	"runtime"
	"time"
)

func readMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc: %d bytes, HeapIdle: %d bytes, HeapReleased: %d bytes", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func test() {
	s := make([]int, 8)

	log.Println("loop begin")
	for i := 0; i < 32*1000*1000; i++ {
		s = append(s, i)
		if i == 16*1000*1000 {
			readMemStats()
		}
	}
	log.Println("loop end")
}

func main() {
	log.Println("Start")

	readMemStats()
	test()
	readMemStats()

	log.Println("force GC")
	runtime.GC()
	log.Println("Done")

	readMemStats()

	go func() {
		for {
			readMemStats()
			time.Sleep(10 * time.Second)
		}
	}()

	time.Sleep(time.Hour)
}
