package composite

import (
	"testing"
)

func TestCompositePattern(t *testing.T) {

	stringSensor := Sensor{
		name: "string sensor",
		data: "sensor data",
	}

	floatSensor := Sensor{
		name: "float sensor",
		data: 567.34,
	}
	lamb := LambThing{
		"lamb_thing",
		stringSensor,
	}

	lamb.PublishData()

	sensors := []Thing{stringSensor, floatSensor, lamb}

	car := CarThing{
		Sensors: sensors,
	}
	car.PublishData()

}
