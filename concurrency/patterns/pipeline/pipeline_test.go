package pipeline

import (
	"github.com/likexian/gokit/assert"
	"testing"
)

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}

	var res int
	for _, test := range tableTest {
		res = LaunchPipeline(test[0])
		t.Logf("%d == %d\n", res, test[1])

		assert.Equal(t, res, test[1])

	}
}
