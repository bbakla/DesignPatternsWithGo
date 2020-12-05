package main

import (
	// "bytes"
	// "strings"
	"fmt"
)

type BallAdd interface {
	AddBall() string
}

type Icecream struct {
}

// type IcecreamDecorator struct {
// 	Decorators BallAdd
// }

func (decorator *Icecream) AddBall() string {

	return "icecream with "
	// var buffer bytes.Buffer
	// buffer.WriteString("Icecream with\n")
	// for _, d := range decorator.Decorators {
	// 	buffer.WriteString(d.AddBall())
	// 	buffer.WriteString("\n")
	// }

	// return buffer.String()
}

type Gurke struct {
	Ball BallAdd
}

func (gurke *Gurke) AddBall() string {
	return gurke.Ball.AddBall() + "Gurke "
}

type Hazelnuss struct {
	Ball BallAdd
}

func (hazelnus *Hazelnuss) AddBall() string {
	return hazelnus.Ball.AddBall() + "Hazelnus "
}

func main() {

	iceCream := &Icecream{}

	gurke := &Gurke{iceCream}
	hazelnuss := &Hazelnuss{gurke}

	fmt.Println(hazelnuss.AddBall())

}
