package main

import "sync"

type singletonWithMutex struct {
	count int
	sync.RWMutex
}

var instance singletonWithMutex

func GetSingletonWithMutex() *singletonWithMutex {
	return &instance
}

func (s *singletonWithMutex) AddOne() {
	s.Lock()
	defer s.Unlock()

	s.count++
}

func (s *singletonWithMutex) GetCount() int {
	s.RLock()
	defer s.RUnlock()

	return s.count
}
