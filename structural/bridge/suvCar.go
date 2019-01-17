package bridge

import (
	"fmt"
)

type Car interface {
	StopTheCar()
	SlowDown(factor int)
	GoFaster()
}

type Suv struct {
	Brake Brake
}

func (suv Suv) StopTheCar() {
	fmt.Println("Suv stops")
	suv.Brake.stop()
}

func (suv Suv) SlowDown(factor int) {
	fmt.Println("Suv slows down")
	suv.Brake.slowDown(factor)
}

func (suv Suv) GoFaster() {
	fmt.Println("Suv goes faster")
	suv.Brake.release()
}
