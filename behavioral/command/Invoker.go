package command

type Invoker struct {
	Commands []Command
}

func (invoker *Invoker) Execute(command Command) {
	invoker.Commands = append(invoker.Commands, command)
	command.execute()
}

type Invoker2 struct {
}

func (invoker *Invoker2) Execute(fn func()) {
	fn()
}
