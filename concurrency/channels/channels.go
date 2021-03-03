package main

import (
	"time"
)

func main() {
	//directionalChannels_OnlyInsert()
	//bufferedChannel()

	rangeOverChannels()

}

func rangeOverChannels() {
	ch := make(chan int)
	//ch2 := make(chan int)

	go func() {
		ch <- 1
		//	ch2 <- 2
		//time.Sleep(time.Second)
		ch <- 2
		//	ch2 <- 4
		ch <- 3
		//	ch2 <- 6
		close(ch)
		//	close(ch2)
	}()

	for v := range ch {
		println(v)
	}

	/*for v := range ch2 {
		println(v)
	}*/
}

func directionalChannels_OnlyInsert() {
	channel := make(chan string, 1)

	go func(ch chan<- string) {
		ch <- "Hello World"
		println("Finishing goroutine")
	}(channel)

	time.Sleep(time.Second)
	message := <-channel
	println(message)
}

func bufferedChannel() {
	channel := make(chan string, 1)

	go func() {
		channel <- "Hello World"
		/*Since we add the second message to the channel, go routine will never finish before the main function exits.
		So "Finishing goroutine" will never print
		*/
		channel <- "Hello world 2"
		println("Finishing goroutine")
	}()

	time.Sleep(time.Second * 4)

	message := <-channel
	println(message)

	//If we don read the second message in the channel, go routine will never finish before the main function exits
	message = <-channel
	println(message)
}

func basicChannel() {
	println("starting")

	channel := make(chan string)
	go func() {
		channel <- "hello world"
		time.Sleep(time.Second * 3)
	}()

	message := <-channel
	println(message)
}
