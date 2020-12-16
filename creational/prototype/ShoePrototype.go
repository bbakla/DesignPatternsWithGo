package main

type ShoePrototype interface {
	clone() ShoePrototype
	getColor() ShoeColor
	getGender() Gender
	getName() string
}

type ShoeColor string

const (
	White ShoeColor = "White"
	Black ShoeColor = "Black"
	Mixed ShoeColor = "Mixed"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
	Unisex Gender = "U"
)

type femaleShoe struct {
	name   string
	color  ShoeColor
	gender Gender
}

func (ws *femaleShoe) getColor() ShoeColor {
	return ws.color
}

func (ws *femaleShoe) getGender() Gender {
	return ws.gender
}

func (ws *femaleShoe) getName() string {
	return ws.name
}

func (ws *femaleShoe) clone() ShoePrototype {
	return &femaleShoe{
		name:   ws.name + "_clone",
		color:  ws.color,
		gender: Female,
	}
}

type maleShoe struct {
	name   string
	color  ShoeColor
	gender Gender
}

func (ws *maleShoe) getColor() ShoeColor {
	return ws.color
}

func (ws *maleShoe) getGender() Gender {
	return ws.gender
}

func (ws *maleShoe) getName() string {
	return ws.name
}

func (ws *maleShoe) clone() ShoePrototype {
	return &maleShoe{
		name:   ws.name + "_clone",
		color:  ws.color,
		gender: Male,
	}
}
