package main

import "fmt"

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type Motorbike interface {
	GetMotorbikeType() int
}

type SportMotorbike struct{}

func (s *SportMotorbike) NumWheels() int {
	return 2
}

func (s *SportMotorbike) NumSeats() int {
	return 1
}

func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}

func (s *SportMotorbike) GetProperties() {
	fmt.Printf("Motorbike type :\t%d\nNumber of Wheels:\t%d\nNumber of Seats:\t%d\n", s.GetMotorbikeType(), s.NumWheels(), s.NumSeats())
}

type CruiseMotorbike struct{}

func (s *CruiseMotorbike) NumWheels() int {
	return 2
}

func (s *CruiseMotorbike) NumSeats() int {
	return 1
}

func (s *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
func (s *CruiseMotorbike) GetProperties() {
	fmt.Printf("Motorbike type :\t%d\nNumber of Wheels:\t%d\nNumber of Seats:\t%d\n", s.GetMotorbikeType(), s.NumWheels(), s.NumSeats())
}
