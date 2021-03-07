package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

/*
It is not ideal to do it with waiting groups but I wanted to try it
PROBLEMS:
* returning and argument from goroutine was problematic. I had to make responses struct
global

* executing HTTP request was a problem. GET Request didnt return anything. Moving wg.Done() into that
method solved the problem but this time I had to inject waitingGroup into makeRequestWithNoChannel function.
Channels are the best
*/

var responses []Response

func withWaitingGroups(endpoints []string) (string, error) {

	var wg sync.WaitGroup
	wg.Add(len(endpoints))

	for i, endpoint := range endpoints {
		go func() {
			makeRequestWithNoChannel(endpoint, i, wg)
		}()

	}
	var stringBuilder strings.Builder

	for _, response := range responses {
		stringBuilder.WriteString(response.Response)
	}

	wg.Wait()

	return stringBuilder.String(), nil
}

func makeRequestWithNoChannel(endpoint string, i int, wg sync.WaitGroup) {
	httpClient := http.Client{
		Timeout: time.Duration(timeoutMilliSeconds) * time.Millisecond,
	}

	barrierResponse := Response{}

	response, err := httpClient.Get(endpoint)

	if err != nil {
		barrierResponse.Error = err
		responses[i] = barrierResponse

		return

	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		barrierResponse.Error = err
		responses[i] = barrierResponse
		fmt.Printf("error is %s\n", barrierResponse.Error.Error())

		return
	}

	barrierResponse.Response = string(body)
	fmt.Printf("barrier response %s\n", barrierResponse.Response)

	responses[i] = barrierResponse

	wg.Done()
}
