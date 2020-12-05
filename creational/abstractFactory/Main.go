package main

import (
	"fmt"
	//"github.com/bbakla/DesignPatternsWithGo/creational/abstractFactory"
)

func main() {

	fmt.Println("building the car factory")
	carFactory, _ := BuildFactory(1)

	fmt.Println("LuxuryCar is being created")
	luxuryCar, _ := carFactory.NewVehicle(LuxuryCarType)
	luxuryCar.GetProperties()
}
