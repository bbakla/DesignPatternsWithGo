package strategy

import (
	"os"
)

type Context struct {
	Read
}

func (context *Context) ReadFile(fileName string) *os.File {
	return context.Read.ReadFile(fileName)
}

func (context *Context) ReadFiles() {
	context.Read.ReadFiles()
}

func (context *Context) ListFileNames() []string {
	return context.Read.GetFilesNameList()
}
