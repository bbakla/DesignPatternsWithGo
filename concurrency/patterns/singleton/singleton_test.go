package main

import (
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	for i := 0; i < n; i++ {
		singCount := singleton.GetCount()
		sing2Count := singleton2.GetCount()
		println(singCount)
		println(sing2Count)
	}

	//assert.Equal(t, 0, singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(100 * time.Millisecond)
	}

	/*assert.NotEqual(t, singCount, 0)
	assert.NotEqual(t, sing2Count, 0)*/

	singleton.Stop()
}
