package main

import "github.com/bbakla/DesignPatternsWithGo/behavioral/mediator"

func main() {
	tower := mediator.NewAirTower(2)

	plane1 := mediator.Airbus{Name: "Airbus1", Mediator: tower}
	plane2 := mediator.Boeing{Name: "Boeing1", Mediator: tower}
	plane3 := mediator.Airbus{Name: "Airbus2", Mediator: tower}
	plane4 := mediator.Boeing{Name: "Boeing2", Mediator: tower}

	plane1.RequestArrival()
	plane2.RequestArrival()
	plane3.RequestArrival()
	plane4.RequestArrival()

	plane1.Departure()
	plane2.Departure()

}
