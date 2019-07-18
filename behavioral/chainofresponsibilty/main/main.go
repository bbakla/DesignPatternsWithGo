package main

import (
	"fmt"
	chain "github.com/bbakla/DesignPatternsWithGo/behavioral/chainofresponsibilty"
)

func main() {

	professor := chain.ProfessorDoctor{}
	specilized := chain.SpecializedDoctor{&professor, "eye"}
	practitioner := chain.Practitioner{&specilized}

	sympthom := &chain.Sympthom{
		Area:       "eye",
		Difficulty: "easy",
	}
	practitioner.Next(sympthom)

	fmt.Println("------------------------------------")

	sympthom = &chain.Sympthom{
		Area:       "eye",
		Difficulty: "medium",
	}
	practitioner.Next(sympthom)

	fmt.Println("------------------------------------")

	sympthom = &chain.Sympthom{
		Area:       "nose",
		Difficulty: "difficult",
	}
	practitioner.Next(sympthom)

	fmt.Printf("\n")
	fmt.Println("------------------------------------")
	fmt.Println("With closure")
	fmt.Println("------------------------------------")

	sympthom = &chain.Sympthom{
		Area:       "nose",
		Difficulty: "medium",
	}
	closureDoctor := &chain.ClosureDoctor{
		Closure: func(s *chain.Sympthom) {
			fmt.Printf("Closure Doctor handling the issue: %s -> %s", s.Area, s.Difficulty)
		},
	}

	professor.NextDoctor = closureDoctor
	practitioner.Next(sympthom)

}
