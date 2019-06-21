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

type Hatchback struct {
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
