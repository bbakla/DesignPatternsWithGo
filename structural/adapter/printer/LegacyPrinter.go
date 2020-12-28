package main

import "fmt"

//adaptee
type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (message string) {
	message = fmt.Sprintf("Legacy Printer: %s\n", s)

	fmt.Println(message)

	return
}
