package bridge

import (
	"fmt"
)

type drumBrake struct {
}

func (brake drumBrake) release() {
	fmt.Println("drumBrake is disattached.")
}

func (brake drumBrake) stop() {
	fmt.Println("drumBrake is attached to the wheel. Car will stop in 10 secs")
}

func (brake drumBrake) slowDown(breakMagnitude int) {
	fmt.Printf("Car will be slowed down by the factor of %d by drumBrake\n", breakMagnitude)
}
