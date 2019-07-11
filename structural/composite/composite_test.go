package composite

import (
	"testing"
)

func TestCompositePattern(t *testing.T) {

	stringSensor := &Sensor{
		data: "sensor data",
	}

	floatSensor := &Sensor{
		data: 567.34,
	}
	sensors := []*Sensor{stringSensor, floatSensor}

	car := CarThing{
		Sensors: sensors,
	}
	car.publishData()

	lamb := LambThing{
		stringSensor,
	}

	lamb.publishData()

}
