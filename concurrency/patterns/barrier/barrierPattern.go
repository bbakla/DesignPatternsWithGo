package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var timeoutMilliSeconds = 4000

type barrierResp struct {
	Err      error
	Response string
}

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)

	in := make(chan barrierResp, requestNumber)
	defer close(in)

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	for i := 0; i < requestNumber; i++ {
		response := <-in

		if response.Err != nil {
			fmt.Println("ERROR: ", response.Err)

		} else {
			fmt.Println(response.Response)
		}
	}

}

func makeRequest(out chan barrierResp, endpoint string) {
	barrierResponse := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(timeoutMilliSeconds) * time.Millisecond,
	}

	resp, err := client.Get(endpoint)
	if err != nil {
		barrierResponse.Err = err
		out <- barrierResponse
		return
	}

	responseAsByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		barrierResponse.Err = err
		out <- barrierResponse
		return
	}

	barrierResponse.Response = string(responseAsByte)
	out <- barrierResponse

}

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer

	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	writer.Close()
	temp := <-out

	return temp
}
