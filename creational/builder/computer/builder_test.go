package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	computerBuilder := ComputerBuilder{}

	computer := computerBuilder.ChooseBrand("IBM").ChooseMonitor("14 Inch").ChooseProcessor("Intel").ChooseStorageSize(456).Build()

	fmt.Println(computer.Brand)

	//by keeping a computer instance in the builder we violate the immutability. The builder returns always the same object.
	/*	c := computerBuilder.ChooseBrand("Dell").Build()
		fmt.Println(c.Brand)
		fmt.Println(c.Monitor)*/

	assert.Equal(t, "IBM", computer.Brand)

}
