package main

import "fmt"

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type Car interface {
	NumDoors() int
}

type LuxuryCar struct{}

func (l *LuxuryCar) NumDoors() int {
	return 4
}

func (l *LuxuryCar) NumWheels() int {
	return 4
}

func (l *LuxuryCar) NumSeats() int {
	return 5
}
func (l *LuxuryCar) GetProperties() {
	fmt.Printf("Number of Doors:\t%d\nNumber of Wheels:\t%d\nNumber of Seats:\t%d\n", l.NumDoors(), l.NumWheels(), l.NumSeats())
}

type FamilyCar struct{}

func (f *FamilyCar) NumDoors() int {
	return 5
}

func (f *FamilyCar) NumWheels() int {
	return 4
}

func (f *FamilyCar) NumSeats() int {
	return 5
}
func (f *FamilyCar) GetProperties() {
	fmt.Printf("Number of Doors:\t%d\nNumber of Wheels:\t%d\nNumber of Seats:\t%d\n", f.NumDoors(), f.NumWheels(), f.NumSeats())
}
