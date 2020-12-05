package main

import (
	"errors"
	"fmt"
)

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func BuildFactory(f int) (AbstractVehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return &CarFactory{}, nil
	case MotorbikeFactoryType:
		return &MotorbikeFactory{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))
	}
}

type AbstractVehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d is not recognized\n", v))
	}
}

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(LuxuryCar), nil
	case CruiseMotorbikeType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d is not recognized\n", v))
	}

}
