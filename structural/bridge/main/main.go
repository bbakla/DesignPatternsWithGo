package main

import (
	"fmt"
	bridge "github.com/bbakla/DesignPatternsWithGo/structural/bridge"
)
/**
Bridge patterns decouples the abstraction from its implementation.
As you can see in the code, Brake and car are abstractions and none of the implementations of
Car or Brake know each others implementation. Any change in Brake wont affect 
Car object using Brake implementation 
*/

func main() {
	discBrake := new(bridge.DiscBrake)
	drumBrake := new(bridge.DrumBrake)

	suvCar := bridge.Suv{discBrake}
	hatchBack := bridge.Hatchback{drumBrake}

	suvCar.GoFaster()
	suvCar.SlowDown(5)
	suvCar.StopTheCar()

	fmt.Println("---------------------")

	hatchBack.GoFaster()
	hatchBack.SlowDown(10)
	hatchBack.StopTheCar()
}
