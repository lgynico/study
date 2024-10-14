package main

import (
	"fmt"
	"time"
)

type Worker struct {
	id  int
	err error
}

func (p *Worker) Start(dieChan chan<- *Worker) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				p.err = err
			} else {
				p.err = fmt.Errorf("Panic happened with [%v]", r)
			}
		} else {
			p.err = err
		}

		dieChan <- p
	}()

	fmt.Println("Start Worker ... ID = ", p.id)

	// for i := 0; i < 5; i++ {
	// 	time.Sleep(time.Second)
	// }
	time.Sleep(5 * time.Second)

	panic("worker panic ...")

	return err
}

type WorkerManager struct {
	nWorker int
	dieChan chan *Worker
}

func NewWorkerManager(nWorker int) *WorkerManager {
	return &WorkerManager{
		nWorker: nWorker,
		dieChan: make(chan *Worker, nWorker),
	}
}

func (p *WorkerManager) Start() {
	for i := 0; i < p.nWorker; i++ {
		woker := &Worker{
			id: i,
		}

		go woker.Start(p.dieChan)
	}

	p.keepaliveWorkers()
}

func (p *WorkerManager) keepaliveWorkers() {
	for w := range p.dieChan {
		fmt.Printf("Worker %d stopped with err: [%v]\n", w.id, w.err)
		w.err = nil
		go w.Start(p.dieChan)
	}
}

// 一个工作线程池的例子
func main() {
	workerManager := NewWorkerManager(10)
	workerManager.Start()
}
