package main

/*import (
	"fmt"
	"sync"
	"time"
)

type Observer struct {
	//Id string
	Name chan string
}

func (o *Observer) Observe() {
	evt := <- o.Name
	fmt.Println("Observer is: ", o.Name, "Event is: ", evt)
	waitGroup.Done()
}


func main() {
	m := &Match{}
	m.mu = &sync.Mutex{}
	obs := []Observer{
		{make(chan string, 1)},
		{ make(chan string, 2)},
	}

	waitGroup = &sync.WaitGroup{}
	waitGroup.Add(len(obs))

	for _, observer := range obs {
		m.Subscribe(observer.Name)
		go observer.Observe()
	}

	go func() {
		<-time.After(1 * time.Second)
		m.notify("2-5")
	}()
	go func() {
		<-time.After(3 * time.Second)
		m.notify("2-4")
	}()

	waitGroup.Wait()
}
*/
