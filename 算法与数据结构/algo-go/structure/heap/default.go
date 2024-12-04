package heap

import (
	"container/heap"
	"sync"

	"github.com/lgynico/algo-go/utils/constraints"
)

type Default[K constraints.Comparable, V any] struct {
	impl     heapImpl[K]
	elements map[K]V
	once     sync.Once
}

func (p *Default[K, V]) Init() {
	p.once.Do(func() {
		p.elements = make(map[K]V)
	})
}

func (p *Default[K, V]) Put(key K, value V) {
	heap.Push(&p.impl, key)
	p.elements[key] = value
}

func (p *Default[K, V]) Min() (key K, value V, ok bool) {
	if !p.IsEmpty() {
		key = heap.Pop(&p.impl).(K)
		value = p.elements[key]
		ok = true

		delete(p.elements, key)
	}
	return
}

func (p *Default[K, V]) Size() int {
	return p.impl.Len()
}

func (p *Default[K, V]) IsEmpty() bool {
	return p.Size() == 0
}
