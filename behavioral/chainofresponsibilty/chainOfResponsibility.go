package chainofresponsibilty

import (
	"fmt"
)

type DoctorChain interface {
	Next(*Sympthom)
}

type Sympthom struct {
	Difficulty
	Area string
}

type Difficulty string

const (
	Difficult Difficulty = "Difficult"
	Easy      Difficulty = "Easy"
	Medium    Difficulty = "Medium"
)

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

	if sympthom.Difficulty == Difficult {
		fmt.Printf("I am not enough. %s can handle %s problems\n", sympthom.Area, sympthom.Difficulty)
		p.NextDoctor.Next(sympthom)
	} else {
		fmt.Printf("Practitioner can handle %s %s problems\n", sympthom.Difficulty, sympthom.Area)
	}
}

type SpecializedDoctor struct {
	NextDoctor DoctorChain
	Area       string
}

func (s *SpecializedDoctor) Next(sympthom *Sympthom) {
	if sympthom.Area == s.Area {
		fmt.Printf("%s doctor is handling the issue as it is about %s\n", s.Area, sympthom.Area)

	} else {
		fmt.Printf("%s doctor is forwarding too as it is about %s\n", s.Area, sympthom.Area)
		sympthom.Difficulty = Difficult
		s.NextDoctor.Next(sympthom)
	}
}

type ProfessorDoctor struct {
	NextDoctor DoctorChain
}

func (p *ProfessorDoctor) Next(sympthom *Sympthom) {

	if sympthom.Difficulty == Difficult {
		fmt.Printf("Professor is handling %s\n", sympthom.Area)
	}
	if p.NextDoctor != nil && sympthom.Area == "nose" {
		fmt.Printf("Professor is also asking to closureDoctor as it is about %s\n", sympthom.Area)
		p.NextDoctor.Next(sympthom)
	}
}
