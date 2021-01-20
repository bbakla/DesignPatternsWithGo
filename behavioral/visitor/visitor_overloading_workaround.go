package main

import "fmt"

//Since there is no method overloading in Go, we solved the problem with polymorphism. That is why
//we have one more interface than classical definition of the pattern.

//Visitor interface
type Visitor interface {
	visit(operation DocumentElement)
}

type CopyOperationVisiter struct{}

func (v *CopyOperationVisiter) visit(element DocumentElement) {
	element.copy()
}

type CutOperationsVisiter struct{}

func (v *CutOperationsVisiter) visit(element DocumentElement) {
	element.cut()
}

//Element interface
type DocumentElement interface {
	copy()
	cut()
}

//Visitable. Another element interface
type Visitable interface {
	accept(visitor Visitor)
}

type ButtonVisitable struct {
	pressed bool
	name    string
}

func (v *ButtonVisitable) accept(visitor2 Visitor) {
	visitor2.visit(v)
}

func (v *ButtonVisitable) copy() {
	fmt.Printf("Copying the button, %s\n", v.name)
}

func (v *ButtonVisitable) cut() {
	fmt.Printf("Cutting the button, %s\n", v.name)
}

type TextVisitable struct {
	value string
	name  string
}

func (v *TextVisitable) accept(visitor Visitor) {
	visitor.visit(v)
}

func (v *TextVisitable) copy() {
	fmt.Printf("Copying the text, %s\n", v.name)
}

func (v *TextVisitable) cut() {
	fmt.Printf("Cutting the text, %s\n", v.name)
}

func main() {
	textVisitable := &TextVisitable{
		value: "test",
		name:  "testText",
	}

	buttonVisible := &ButtonVisitable{
		pressed: false,
		name:    "buttonTest",
	}

	copyVisitor := &CopyOperationVisiter{}
	cutVisitor := &CutOperationsVisiter{}

	buttonVisible.accept(copyVisitor)
	textVisitable.accept(copyVisitor)

	fmt.Println("----------------------------")

	buttonVisible.accept(cutVisitor)
	textVisitable.accept(cutVisitor)

}
