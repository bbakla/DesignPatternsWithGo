package main

import (
	"fmt"
	"github.com/bbakla/DesignPatternsWithGo/behavioral/command"
)

func main() {
	receiver := &command.Editor{}
	editCommand := command.EditCommand{
		Editor:       receiver,
		DocumentName: "main.txt"}

	saveCommand := command.SaveCommand{
		Editor:       receiver,
		DocumentName: "main.txt",
	}

	copyCommand := command.CopyCommand{
		Editor:     receiver,
		CopiedText: "copied text",
	}

	invoker := &command.Invoker{Commands: []command.Command{}}

	invoker.Execute(&editCommand)
	invoker.Execute(&copyCommand)
	invoker.Execute(&saveCommand)

	fmt.Println("-----------------------------------")
	fmt.Println("Shorter version for commands")
	fmt.Println("-----------------------------------")

	editCommand2 := func() {
		receiver.Edit("main2.txt")
	}

	saveCommand2 := func() {
		receiver.Save("main2.txt")
	}

	copyCommand2 := func() {
		receiver.Copy("copy text2")
	}

	invoker2 := &command.Invoker2{}
	invoker2.Execute(editCommand2)
	invoker2.Execute(copyCommand2)
	invoker2.Execute(saveCommand2)

}
