package main

func main() {
	oldPrinter := &MyLegacyPrinter{}
	adapter := &Adapter{
		OldPrinter: oldPrinter,
		Message:    "print me",
	}

	adapter.PrintStored()
}
