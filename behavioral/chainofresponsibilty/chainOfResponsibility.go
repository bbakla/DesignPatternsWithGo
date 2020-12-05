package chainofresponsibilty

import (
	"fmt"
)

type DoctorChain interface {
	Next(*Sympthom)
}

type Sympthom struct {
	Difficulty string
	Area       string
}

type ClosureDoctor struct {
	NextDoctor DoctorChain
	Closure    func(*Sympthom)
}

func (c *ClosureDoctor) Next(sympthom *Sympthom) {
	c.Closure(sympthom)

	//	c.Next(sympthom)
}

type Practitioner struct {
	NextDoctor DoctorChain
}

func (p *Practitioner) Next(sympthom *Sympthom) {
	fmt.Println("Practitioner")

	if sympthom.Difficulty == "easy" {
		fmt.Println("Practitioner can handle it")
		return
	} else {
		sympthom.Difficulty = "medium"
		p.NextDoctor.Next(sympthom)
	}
}

type SpecializedDoctor struct {
	NextDoctor DoctorChain
	Area       string
}

func (s *SpecializedDoctor) Next(sympthom *Sympthom) {
	fmt.Println("Specialized")
	if sympthom.Area == s.Area {
		fmt.Printf("%s doctor is handling\n", s.Area)
		return
	} else {
		sympthom.Difficulty = "difficult"
		s.NextDoctor.Next(sympthom)
	}
}

type ProfessorDoctor struct {
	NextDoctor DoctorChain
}

func (p *ProfessorDoctor) Next(sympthom *Sympthom) {
	fmt.Println("Professor")

	if sympthom.Difficulty == "difficult" {
		fmt.Println("Professor can handle it")
	}
	if p.NextDoctor != nil {
		p.NextDoctor.Next(sympthom)
	}
}
