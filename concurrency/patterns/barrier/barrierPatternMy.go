package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	Error    error
	Response string
}

func combineResponses(endpoints []string) (string, error) {
	numberOfRequests := len(endpoints)

	inputChannel := make(chan Response, numberOfRequests)
	defer close(inputChannel)

	for _, endpoint := range endpoints {
		go makeHttpRequest(inputChannel, endpoint)
	}

	var stringBuilder strings.Builder

	/*	for {
		select {
		case resp := <- inputChannel:
			if resp.Error != nil {
				return "", resp.Error
			}
			stringBuilder.WriteString(resp.Response)
		default:

			break
		}
	}*/

	for resp := range inputChannel {
		//go func() {}()
		if resp.Error != nil {
			return "", resp.Error
		}
		stringBuilder.WriteString(resp.Response)
	}

	/*for i := 0; i < numberOfRequests; i++ {
		resp := <- inputChannel

		if resp.Error != nil {
			return "", resp.Error
		}
		stringBuilder.WriteString(resp.Response)
	}*/

	return stringBuilder.String(), nil
}

func makeHttpRequest(channel chan Response, endpoint string) {
	httpClient := http.Client{
		//	Timeout: time.Duration(timeoutMilliSeconds) * time.Millisecond,
	}

	barrierResponse := Response{}

	response, err := httpClient.Get(endpoint)

	if err != nil {
		barrierResponse.Error = err
		channel <- barrierResponse

		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		barrierResponse.Error = err
		channel <- barrierResponse

		fmt.Printf("error is %s\n", barrierResponse.Error.Error())

		return
	}

	barrierResponse.Response = string(body)
	fmt.Printf("barrier response %s\n", barrierResponse.Response)
	channel <- barrierResponse

}
