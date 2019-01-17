package bridge

import (
	"fmt"
)

type Brake interface {
	release()
	stop()
	slowDown(breakMagnitude int)
}

type DiscBrake struct {
}

func (brake DiscBrake) release() {
	fmt.Println("Diskbrake is disattached.")
}

func (brake DiscBrake) stop() {
	fmt.Println("Diskbrake is attached to the wheel. Car will stop in 3 secs")
}

func (brake DiscBrake) slowDown(breakMagnitude int) {
	fmt.Printf("Car will be slowed down by the factor of %d\n", breakMagnitude)
}
