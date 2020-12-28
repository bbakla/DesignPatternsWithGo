package main

import "fmt"

type Adapter struct {
	OldPrinter LegacyPrinter
	Message    string
}

//adapts the old printer here.
func (p *Adapter) PrintStored() (message string) {
	message = fmt.Sprintf("Adapter: %s\n", p.Message)

	return p.OldPrinter.Print(message)
}
