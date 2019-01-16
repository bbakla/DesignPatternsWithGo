package bridge

import (
	"fmt"
)

type Car interface {
	stopTheCar()
	slowDown()
	goFaster()
}

type Suv struct {
}
