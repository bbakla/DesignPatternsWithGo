package main

/*import "sync"

var waitGroup *sync.WaitGroup

type Match struct {
	observers []chan string
	//observers []chan int
	mu *sync.Mutex
}

func (m *Match) Subscribe(name chan string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.observers = append(m.observers, name)
}


func (m *Match) Unsubscribe(name chan string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, observer := range m.observers {
		if observer == name {
			m.observers = append(m.observers[:i], m.observers[i+1:]...)
			return
		}
	}
}

func (m *Match) notify(event string)  {
	for _, observer := range m.observers {
		 observer <- event
	}
}



*/

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup

type Observable struct {
	observers []chan int
	mu        *sync.Mutex
}

func (o *Observable) Attach(c chan int) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.observers = append(o.observers, c)
}

func (o *Observable) Detach(c chan int) {
	o.mu.Lock()
	defer o.mu.Unlock()
	for i, v := range o.observers {
		if v == c {
			o.observers = append(o.observers[:i], o.observers[i+1:]...)
			return
		}
	}
}

func (o *Observable) Notify(evt int) {
	for _, v := range o.observers {
		v <- evt
	}
}

type Observer struct {
	Food string
	ch   chan int
}

func (obs *Observer) Observe() {
	evt := <-obs.ch
	fmt.Println("Food:", obs.Food, "Event:", evt)
	wg.Done()
}

func main() {
	o := &Observable{}
	o.mu = &sync.Mutex{}
	obs := []Observer{
		{"Eggs", make(chan int, 2)},
		{"Bacon", make(chan int, 2)},
	}
	wg = &sync.WaitGroup{}
	wg.Add(len(obs))
	for _, v := range obs {
		o.Attach(v.ch)
		go v.Observe()
	}

	go func() {
		<-time.After(1 * time.Second)
		o.Notify(3)
	}()
	go func() {
		<-time.After(3 * time.Second)
		o.Notify(5)
	}()

	wg.Wait()
}
