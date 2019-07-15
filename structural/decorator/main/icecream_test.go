package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCanAddBall(t *testing.T) {
	icecream := &IcecreamDecorator{}

	fmt.Println(icecream.AddBall())

}

func TestCanAddHazelNussToTheIcecream(t *testing.T) {
	icecream := &IcecreamDecorator{}

	hazelNuss := &Hazelnuss{icecream}
	addBall := hazelNuss.AddBall()

	assert.True(t, strings.Contains(addBall, "Hazel"), "Result is "+addBall)

}
