package main

import (
	"fmt"
	"github.com/bbakla/DesignPatternsWithGo/behavioral/strategy"
)

func main() {

	strategyToAccessToFiles := strategy.FileReaderStrategy{DirectoryPath: "filePath"}

	context := &strategy.Context{
		Read: &strategyToAccessToFiles,
	}

	context.ReadFiles()
	context.ReadFile("read")
	context.GetFilesNameList()

	fmt.Println("----------------------------------------------------")
	context.Read = &strategy.S3ReaderStrategy{Password: "aa"}
	context.ReadFiles()
	context.ReadFile("read")
	context.GetFilesNameList()

}
