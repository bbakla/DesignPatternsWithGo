package main

import "fmt"

type ElementVisitor interface {
	visitButton(element button)
	visitText(element text)
}

type CopyVisitor struct {
}

func (v *CopyVisitor) visitButton(element button) {
	fmt.Printf("Copying the button, %s\n", element.name)
}

func (v *CopyVisitor) visitText(element text) {
	fmt.Printf("Copying the text %s\n", element.name)
}

type CutVisitor struct {
}

func (v *CutVisitor) visitButton(element button) {
	fmt.Printf("Cutting the button, %s\n", element.name)
}

func (v *CutVisitor) visitText(element text) {
	fmt.Printf("Cutting the text %s\n", element.name)
}

type Element interface {
	accept(visitor ElementVisitor)
}

type button struct {
	pressed bool
	name    string
}

func (c *button) accept(visitor ElementVisitor) {
	visitor.visitButton(*c)
}

type text struct {
	value string
	name  string
}

func (c *text) accept(visitor ElementVisitor) {
	visitor.visitText(*c)
}

func main1() {
	button := &button{
		pressed: false,
		name:    "testButton",
	}

	text := &text{
		value: "this is is test",
		name:  "testText",
	}

	copyVisitor := &CopyVisitor{}
	cutVisitor := &CutVisitor{}

	button.accept(copyVisitor)
	text.accept(copyVisitor)

	fmt.Println("--------------------------------")

	button.accept(cutVisitor)
	text.accept(cutVisitor)

}
