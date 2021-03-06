package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type TestArgument struct {
	Name      string
	Addresses []string
	expected  bool
}

type TestArguments []TestArgument

func TestTestBarrier(t *testing.T) {
	testarguments := TestArguments{
		{"correct inputs", []string{"http://httpbin.org/headers", "http://httpbin.org/user-agent"}, true},
		//{"mallformed linkds", []string{"http://malformed", "http://httpbin.org/User-Agent"}, false},
	}

	for _, a := range testarguments {
		t.Run(a.Name, func(t *testing.T) {
			/*result := captureBarrierOutput(a.Addresses...)

			contains := strings.Contains(result, "Accept-Encoding") && strings.Contains(result, "user-agent")
			assert.Equal(t, contains, a.expected)*/

			//MY IMPLEMENTATION
			result, err := combineResponses(a.Addresses)
			if a.expected {
				assert.Nil(t, err)
				contains := strings.Contains(result, "Accept-Encoding") && strings.Contains(result, "user-agent")
				assert.True(t, contains)
			} else {
				assert.NotNil(t, err)
				contains := strings.Contains(result, "Accept-Encoding") && strings.Contains(result, "user-agent")
				assert.False(t, contains)
			}
		})
	}
}
