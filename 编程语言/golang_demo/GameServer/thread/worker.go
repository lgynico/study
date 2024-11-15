package thread

type worker struct {
	taskC chan func()
}

func newWorker(size int32) *worker {
	return &worker{taskC: make(chan func(), size)}
}

func (p *worker) submit(task, callback func()) {
	if task == nil {
		return
	}

	p.taskC <- func() {
		task()
		if callback != nil {
			Logic.Submit(callback)
		}
	}
}

func (p *worker) execute() {
	for task := range p.taskC {
		task()
	}
}
