package main

import (
	"fmt"
	"sync"
)

func main() {
	bufferSize := 100
	var dispatcher = NewDispatcher(bufferSize)

	workers := 3
	for i := 0; i < workers; i++ {
		w := &PrefixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			id:      i,
		}

		dispatcher.LaunchWorker(w)
	}

	requests := 10

	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest(&wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()
	wg.Wait()

}
