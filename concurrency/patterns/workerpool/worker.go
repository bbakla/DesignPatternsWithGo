package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(wg *sync.WaitGroup) Request {
	request := Request{
		Data: "Hello",
		Handler: func(i interface{}) {
			defer wg.Done()

			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}

			fmt.Println(s)
		},
	}

	return request
}

type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

const (
	PREFIX = "prefix"
	SUFFIX = "suffix"
)

type PrefixSuffixWorker struct {
	prefixS string
	id      int
}

func (w *PrefixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.upperCase(in)))
}

func (w *PrefixSuffixWorker) upperCase(inputCh <-chan Request) <-chan Request {
	outputChannel := make(chan Request)

	go func() {
		for msg := range inputCh {
			s, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			outputChannel <- msg
		}
		close(outputChannel)
	}()

	return outputChannel

	/*message := <- inputCh
	outputChannel <- strings.ToUpper(message)

	close(outputChannel)

	return outputChannel*/
}

func (w *PrefixSuffixWorker) append(inputCh <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range inputCh {
			s, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Data = fmt.Sprintf("%s%s", s, SUFFIX)
			out <- msg
		}
		close(out)
	}()

	return out
}

func (w *PrefixSuffixWorker) prefix(inputCh <-chan Request) {

	go func() {
		for msg := range inputCh {
			s, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Handler(fmt.Sprintf("%s%s", w.prefixS, s))
		}
	}()
}
