package thread

import (
	"sync"
)

type logic chan func()

var (
	taskSize = 2048
	Logic    = make(logic, taskSize)
	once     sync.Once
)

func (p logic) Submit(task func()) {
	if task == nil {
		return
	}

	p <- task

	once.Do(func() {
		go p.execute()
	})
}

func (p logic) execute() {
	for task := range p {
		task()
	}
}
