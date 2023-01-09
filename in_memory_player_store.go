package main

func NewInMemoryPlayerStore() *InMemPlayerStore {
	return &InMemPlayerStore{map[string]int{}}
}

type InMemPlayerStore struct {
	store map[string]int
}

func (i *InMemPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemPlayerStore) GetLeague() []Player {
	return nil
}
