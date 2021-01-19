package main

type Context struct {
	Strategy func()
}

func (context *Context) doIt() {
	context.Strategy()
}
