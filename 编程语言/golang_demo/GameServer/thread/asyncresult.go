package thread

import "sync"

type AsyncResult[T any] struct {
	onComplete func(*T)
	once       sync.Once
}

func (p *AsyncResult[T]) Complete(data *T) {
	if data == nil {
		return
	}

	p.once.Do(func() {
		if p.onComplete != nil {
			Logic.Submit(func() {
				p.onComplete(data)
			})
		}
	})
}

func (p *AsyncResult[T]) OnComplete(completeFunc func(data *T)) {
	p.onComplete = completeFunc
}
