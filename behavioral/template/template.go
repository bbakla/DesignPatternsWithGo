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

type Part1Impl struct {
}

func (p *Part1Impl) part1() {
	fmt.Println("part1")
}

type Part2Impl struct {
}

func (p *Part2Impl) part2() {
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
		Part1: &Part1Impl{},
		Part2: &Part2Impl{},
	}

	template.doIt()
}
