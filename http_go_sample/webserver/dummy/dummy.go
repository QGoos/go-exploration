package dummy

import (
	"http_go_sample/webserver"
	"sync"
)

type InMemoryPlayerStore[T webserver.Player] struct {
	store map[string]int
	lock  sync.RWMutex
}

func (i *InMemoryPlayerStore[T]) GetLeague() []T {
	var league []T
	for name, wins := range i.store {
		player := T{Name: name, Wins: wins}
		league = append(league, player)
	}
	return league
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

func NewInMemoryPlayerStore[T webserver.Player]() *InMemoryPlayerStore[T] {
	return &InMemoryPlayerStore[T]{
		map[string]int{},
		sync.RWMutex{},
	}
}
