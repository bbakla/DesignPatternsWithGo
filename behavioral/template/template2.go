package main

import "fmt"

/*
This is again based on composition and polymorphism and little bit more close to classical template pattern approach.
Template method is however is not defined by the interface but the main object using the implementation of the
Algorithm.
*/

//algorithm interface
type start interface {
	insertKey()
	checkMirror()
	startEngine()
}

type Car struct {
	Start start
}

//template method
func (c *Car) startCar() {
	fmt.Println("let drive..")

	c.Start.checkMirror()
	c.Start.insertKey()
	c.Start.startEngine()
}

//Altorithm1
type AutonomousCar struct {
	car Car
}

func (a *AutonomousCar) insertKey() {
	fmt.Println("AutonomousCar: No need to insert key")
}

func (a *AutonomousCar) checkMirror() {
	fmt.Println("AutonomousCar: I dont have any mirrors")
}

func (a *AutonomousCar) startEngine() {
	fmt.Println("AutonomousCar: Starting the engine")
}

//Altorithm2
type ManualCar struct {
	car Car
}

func (a *ManualCar) insertKey() {
	fmt.Println("ManualCar: Inserting the key")
}

func (a *ManualCar) checkMirror() {
	fmt.Println("ManualCar: Checking the mirror")
}

func (a *ManualCar) startEngine() {
	fmt.Println("ManualCar: Starting the engine")
}

func main() {
	manualCar := &ManualCar{}
	autonomous := &AutonomousCar{}

	car := &Car{Start: manualCar}
	car.startCar()

	fmt.Println("-------------------------")

	car.Start = autonomous
	car.startCar()

}
