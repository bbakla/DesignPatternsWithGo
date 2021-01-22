package main

import "fmt"

type Value int

type Array []Value // In practice this is more complicated

type Iterator struct {
	C      chan Value
	closer chan bool
}

func (i Array) GetIterator() *Iterator {
	c := make(chan Value)
	closer := make(chan bool)
	iter := &Iterator{c, closer}

	// Spew out the iterator in a goroutine
	go func() {
		for _, v := range i {
			select {
			case c <- v:
			case <-closer:
				close(c)
				return
			}
		}
		close(c)
		close(closer)
	}()

	return iter
}

func (i *Iterator) Break() {
	select {
	case _, _ = <-i.closer:
		// closer has been closed
	default:
		// closer is still open
		i.closer <- true
	}
}

func main() {
	iterable := Array([]Value{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	iter := iterable.GetIterator()
	for v := range iter.C {
		fmt.Println("Got", v)
		if v >= 8 {
			iter.Break()
			break
		}
	}
	fmt.Println("Hello, playground")

}
