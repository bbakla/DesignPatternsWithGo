package main

type Singleton struct {
}

var singleton Singleton

var addCh = make(chan bool)
var getCountCh = make(chan chan int)
var quitCh = make(chan bool)

func init() {
	var count int

	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return
			}
		}
	}(addCh, getCountCh, quitCh)

}

func GetInstance() Singleton {
	//if singleton == nil {
	//singleton = &Singleton{}
	//}
	return singleton
}

func (s Singleton) AddOne() {
	addCh <- true
}

func (s Singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)

	getCountCh <- resCh

	return <-resCh
}

func (s Singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}
