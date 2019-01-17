package bridge

import (
	"fmt"
)

type DrumBrake struct {
}

func (brake DrumBrake) release() {
	fmt.Println("drumBrake is disattached.")
}

func (brake DrumBrake) stop() {
	fmt.Println("drumBrake is attached to the wheel. Car will stop in 10 secs")
}

func (brake DrumBrake) slowDown(breakMagnitude int) {
	fmt.Printf("Car will be slowed down by the factor of %d by drumBrake\n", breakMagnitude)
}
