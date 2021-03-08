package future

import (
	"errors"
	"sync"
	"testing"
	"time"
)

type TestArgument struct {
	Name     string
	Input    string
	Expected string
}

type TestArguments []TestArgument

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout")

	t.Fail()
	wg.Done()
}

func TestStringOrError_Execute(t *testing.T) {
	testArguments := TestArguments{
		{"Success", "", ""},
		{"Error", "", ""},
	}

	future := &MaybeString{}

	t.Run(testArguments[0].Name, func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "HEllo world", nil
		})

		wg.Wait()
	})

	t.Run(testArguments[1].Name, func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(err error) {
			t.Log(err.Error())
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "", errors.New("error occurred")
		})

		wg.Wait()
	})
}

func TestContext(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	go timeout(t, &wg)

	future := &MaybeString{}
	future.Success(func(s string) {
		t.Log(s)
		wg.Done()
	}).Fail(func(err error) {
		t.Fail()
		wg.Done()
	})

	future.Execute(setContext("Hello"))
	wg.Wait()
}
