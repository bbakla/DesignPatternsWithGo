package main

import "fmt"

func main() {
	manShoesPrototype := &maleShoe{
		name:   "Classic",
		color:  White,
		gender: Male,
	}

	femaleShoePrototype := femaleShoe{
		name:   "Sports",
		color:  Black,
		gender: Female,
	}

	fmt.Printf("male shoe name:%s\n", manShoesPrototype.getName())
	fmt.Printf("female shoe name:%s\n", femaleShoePrototype.getName())

	manShoesPrototype2 := manShoesPrototype.clone()
	fmt.Printf("male shoe color:%s\n", manShoesPrototype2.getName())

	femaleShoesPrototype2 := manShoesPrototype.clone()
	fmt.Printf("male shoe color:%s\n", femaleShoesPrototype2.getName())

}
