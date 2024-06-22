package dummy

import (
	"http_go_sample/webserver"
	"sync"
)

type InMemoryPlayerStore[T webserver.League] struct {
	store  map[string]int
	lock   sync.RWMutex
	league T
}

func (i *InMemoryPlayerStore[T]) GetLeague() T {
	return i.league
}

func (i *InMemoryPlayerStore[T]) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.store[name]
}

func (i *InMemoryPlayerStore[T]) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
}

func NewInMemoryPlayerStore[T webserver.League]() *InMemoryPlayerStore[T] {
	return &InMemoryPlayerStore[T]{
		map[string]int{},
		sync.RWMutex{},
		T{},
	}
}
