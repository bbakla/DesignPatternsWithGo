package main

import "fmt"

/*
Since inheritance is not supported in golang, we implement it based on composition.
Therefore, it is very similar to Strategy pattern.
*/

type Part1 interface {
	part1()
}

type Part2 interface {
	part2()
}

type Part1S struct {
}

func (p *Part1S) part1() {
	fmt.Println("part1")
}

type Part2S struct {
}

func (p *Part2S) part2() {
	fmt.Println("part2")
}

type Template struct {
	Part1 Part1
	Part2 Part2
}

func (t *Template) doIt() {
	fmt.Println("Default implementation")

	t.Part1.part1()
	t.Part2.part2()
}

func main2() {

	template := &Template{
		Part1: &Part1S{},
		Part2: &Part2S{},
	}

	template.doIt()
}
