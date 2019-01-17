package bridge

import (
	"fmt"
)

type Hatchback struct {
	Brake Brake
}

func (hatchback Hatchback) StopTheCar() {
	fmt.Println("Hatchback stops")
	hatchback.Brake.stop()
}

func (hatchback Hatchback) SlowDown(factor int) {
	fmt.Println("Hatchback slows down")
	hatchback.Brake.slowDown(factor)
}

func (hatchback Hatchback) GoFaster() {
	fmt.Println("Hatchback goes faster")
	hatchback.Brake.release()
}
