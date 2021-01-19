package mediator

import "fmt"

type Airplane interface {
	RequestArrival()
	Departure()
	PermitArrival()
}

type Airbus struct {
	Name string
	Mediator
}

func (g *Airbus) RequestArrival() {
	if g.Mediator.canLand(g) {
		fmt.Printf("Airbus %s: Landing. Requesting arrival permit\n", g.Name)
	} else {
		fmt.Printf("Airbus %s: Waiting\n", g.Name)
	}
}

func (g *Airbus) Departure() {
	fmt.Printf("Airbus %s: Leaving\n", g.Name)
	g.Mediator.notifyFree()
}

func (g *Airbus) PermitArrival() {
	fmt.Printf("Airbus: %s Arrival Permitted. Landing\n", g.Name)
}

type Boeing struct {
	Name     string
	Mediator Mediator
}

func (g *Boeing) RequestArrival() {
	if g.Mediator.canLand(g) {
		fmt.Printf("Boeing %s: Landing. Requesting arrival permit\n", g.Name)
	} else {
		fmt.Printf("Boeing %s: Waiting\n", g.Name)
	}
}

func (g *Boeing) Departure() {
	fmt.Printf("Boeing %s: Leaving\n", g.Name)
	g.Mediator.notifyFree()
}

func (g *Boeing) PermitArrival() {
	fmt.Printf("Boeing %s : Arrival Permitted. Landing\n", g.Name)
}
