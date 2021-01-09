package command

import "fmt"

type Editor struct {
}

func (editor *Editor) Copy(text string) {
	fmt.Printf("copying %s\n", text)
}

func (editor *Editor) Edit(documentName string) {
	fmt.Printf("editing %s\n", documentName)
}

func (editor *Editor) Save(documentName string) {
	fmt.Printf("saving %s\n", documentName)
}
