package main

import (
	"fmt"
	chain "github.com/bbakla/DesignPatternsWithGo/behavioral/chainofresponsibilty"
)

func main() {

	closureDoctor := &chain.ClosureDoctor{
		Closure: func(s *chain.Sympthom) {
			fmt.Printf("Closure Doctor handling the issue: %s -> %s", s.Area, s.Difficulty)
		},
	}

	professor := chain.ProfessorDoctor{NextDoctor: closureDoctor}
	specilized := chain.SpecializedDoctor{NextDoctor: &professor, Area: "eye"}
	practitioner := chain.Practitioner{NextDoctor: &specilized}

	sympthom := &chain.Sympthom{
		Area:       "eye",
		Difficulty: chain.Difficult,
	}
	practitioner.Next(sympthom)

	fmt.Println("------------------------------------")

	sympthom = &chain.Sympthom{
		Area:       "eye",
		Difficulty: chain.Medium,
	}
	practitioner.Next(sympthom)

	fmt.Println("------------------------------------")

	sympthom = &chain.Sympthom{
		Area:       "nose",
		Difficulty: chain.Difficult,
	}
	practitioner.Next(sympthom)

	fmt.Printf("\n")
	fmt.Println("------------------------------------")
	fmt.Println("With closure")
	fmt.Println("------------------------------------")

	sympthom = &chain.Sympthom{
		Area:       "nose",
		Difficulty: chain.Medium,
	}

	practitioner.Next(sympthom)
}
