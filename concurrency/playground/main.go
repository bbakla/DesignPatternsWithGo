package main

import (
	"errors"
	"sync"
)

var repos = []string{"a", "b", "c", "d"}

func main() {

}

//https://www.youtube.com/watch?v=DqHb5KBe7qI&list=WL&index=182
func channelsIntroduction() {
	errors := restore([]string{"a", "afdf", "afdsdfsdfd", "b"})

	println(errors.Error())
}

func restore(repos []string) error {
	errChan := make(chan error, len(repos))
	sem := make(chan int, 4)

	var wg sync.WaitGroup
	wg.Add(len(repos))

	for _, repo := range repos {
		go worker(repo, sem, &wg, errChan)
	}

	wg.Wait()
	close(errChan)

	/*e := <- errChan
	println(e.Error())

	e = <- errChan
	println(e.Error())*/

	return <-errChan

}

func worker(repo string, sem chan int, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()
	sem <- 1

	/*if err := fetch(repo); err != nil {
		errChan <- err
	}*/

	if err := fetch(repo); err != nil {
		select {
		case errChan <- err:
			println("error")
		}
	}

	println(<-sem)
}

func fetch(string2 string) error {
	for _, s := range repos {
		if string2 == s {
			return nil
		}
	}

	return errors.New("nop")
}
