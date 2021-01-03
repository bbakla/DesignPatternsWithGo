package main

import (
	// "bytes"
	// "strings"
	"fmt"
)

type IceCream interface {
	AddBall() string
}

//Component
type IceCreamWithCone struct {
}

func (component *IceCreamWithCone) AddBall() string {

	return "icecream with "
	// var buffer bytes.Buffer
	// buffer.WriteString("IceCreamWithCone with\n")
	// for _, d := range decorator.Decorators {
	// 	buffer.WriteString(d.AddBall())
	// 	buffer.WriteString("\n")
	// }

	// return buffer.String()
}

type Gurke struct {
	iceCream IceCream
}

func (gurkeDecorator *Gurke) AddBall() string {
	return gurkeDecorator.iceCream.AddBall() + "Gurke "
}

type Hazelnuss struct {
	IceCream
}

func (hazelnus *Hazelnuss) AddBall() string {
	return hazelnus.IceCream.AddBall() + "Hazelnus "
}

func main() {

	iceCream := &IceCreamWithCone{}

	gurke := &Gurke{iceCream}
	hazelnuss := &Hazelnuss{gurke}

	fmt.Println(hazelnuss.AddBall())

}
