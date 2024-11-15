package thread

import "sync"

var (
	poolSize = 64
	Async    = make(async, poolSize)
	mutex    sync.Mutex
)

type async []*worker

func (p async) Submit(bindID int64, task, callback func()) {
	if task == nil {
		return
	}

	worker := p.getWorker(bindID)
	worker.submit(task, callback)
}

func (p async) getWorker(bindID int64) *worker {
	if bindID < 0 {
		bindID = -bindID
	}

	workerID := bindID % int64(len(p))
	worker := p[workerID]

	if worker == nil {
		worker = p.createWorker(workerID)
	}

	return worker
}

func (p async) createWorker(bindID int64) *worker {
	mutex.Lock()
	defer mutex.Unlock()

	worker := p[bindID]
	if worker != nil {
		return worker
	}

	worker = newWorker(64)
	p[bindID] = worker

	go worker.execute()

	return worker
}
