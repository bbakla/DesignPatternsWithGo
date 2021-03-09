package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Subscriber interface {
	Notify(interface{}) error
	Close()
}

type subscriber struct {
	in     chan interface{}
	id     int
	Writer io.Writer
}

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	if out == nil {
		out = os.Stdout
	}

	s := &subscriber{
		in:     make(chan interface{}),
		id:     id,
		Writer: out,
	}

	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d) : %v\n", s.id, msg)
		}
	}()

	return s
}

func (s *subscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v#v", rec)
		}
	}()

	select {
	case s.in <- msg:
	case <-time.After(time.Second):
		err = fmt.Errorf("Timeout\n")
	}

	return
}

func (s *subscriber) Close() {
	close(s.in)
}
