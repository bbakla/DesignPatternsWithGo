package main

import (
	"github.com/bbakla/DesignPatternsWithGo/behavioral/strategy"
)

func main() {

	strategyToAccessToFiles := strategy.FileReaderStrategy{"filePath"}

	context := &strategy.Context{
		&strategyToAccessToFiles,
	}

	context.ReadFiles()
	context.ReadFile("read")
	context.GetFilesNameList()

}
