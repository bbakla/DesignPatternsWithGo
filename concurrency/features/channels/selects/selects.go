package main

import "time"

func main() {
	helloCh := make(chan string, 1)
	goodbyCh := make(chan string, 1)
	quitCh := make(chan bool)

	go receiver(helloCh, goodbyCh, quitCh)
	go sendString(helloCh, "hello Africa")

	time.Sleep(time.Second * 1)

	go sendString(goodbyCh, "goodbye")
	<-quitCh

}

func sendString(ch chan<- string, s string) {
	ch <- s
}

/*
goodbyeCh and quitCh are receiving channel
helloCh can only be used to send messages
*/
func receiver(helloCh, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			println(msg)
		case msg := <-goodbyeCh:
			println(msg)
		case <-time.After(time.Second * 2):
			println("Nothing receivd in 2 seconds. Exiting")
			quitCh <- true
			break

		}
	}
}
