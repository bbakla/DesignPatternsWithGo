package command

type Command interface {
	execute()
}

type SaveCommand struct {
	*Editor
	DocumentName string
}

func (save *SaveCommand) execute() {
	save.Editor.Save(save.DocumentName)
}

type EditCommand struct {
	*Editor
	DocumentName string
}

func (edit *EditCommand) execute() {
	edit.Edit(edit.DocumentName)
}

type CopyCommand struct {
	*Editor
	CopiedText string
}

func (save *CopyCommand) execute() {
	save.Copy(save.CopiedText)
}
