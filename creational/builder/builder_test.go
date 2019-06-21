package builder

import (
	"testing"
	"github.com/stretchr/testify/assert"
) 

func TestBuilderPattern(t *testing.T) {
	computerBuilder := &ComputerBuilder{}
	
	computer := computerBuilder.ChooseBrand("IBM").ChooseMonitor("14 Inch").ChooseProcessor("Intel").ChooseStorageSize(456).Build()

	assert.Equal(t, "IBM", computer.Brand)
}