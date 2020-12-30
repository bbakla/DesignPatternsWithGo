package composite

import (
	"fmt"
)

//Component
type Thing interface {
	PublishData()
}

//Leaf
type Sensor struct {
	name string
	data interface{}
}

func (sensor Sensor) PublishData() {
	fmt.Printf("%v is published by %s\n", sensor.data, sensor.name)
}

type CarThing struct {
	Sensors []Thing
}

func (car *CarThing) PublishData() {
	for _, sensor := range car.Sensors {
		sensor.PublishData()
	}
}

type LambThing struct {
	name   string
	Sensor Thing
}

func (lamb LambThing) PublishData() {
	fmt.Printf("\nSensor of %s\n", lamb.name)
	lamb.Sensor.PublishData()
}
