package composite

import (
	"fmt"
)

type Thing interface {
	publishData()
}

type Sensor struct {
	data interface{}
}

func (sensor *Sensor) publishData() {
	fmt.Printf("%v is published by \n", sensor.data)
}

type CarThing struct {
	Sensors []*Sensor
}

func (car *CarThing) publishData() {
	for _, sensor := range car.Sensors {
		sensor.publishData()
	}
}

type LambThing struct {
	*Sensor
}

func (lamb *LambThing) publishData() {
	lamb.Sensor.publishData()
}
