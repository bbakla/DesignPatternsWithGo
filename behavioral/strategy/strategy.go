package strategy

import (
	"fmt"
	"os"
)

type Read interface {
	ReadFile(string) *os.File
	ReadFiles()
	GetFilesNameList() []string
}

type FileReaderStrategy struct {
	DirectoryPath string
}

func (fileReader *FileReaderStrategy) ReadFile(fileName string) *os.File {
	fmt.Println(fileName + " Reading file in local")

	return nil
}

func (fileReader *FileReaderStrategy) ReadFiles() {
	fmt.Println("Reading filesin local")
}

func (fileReader *FileReaderStrategy) GetFilesNameList() []string {
	fmt.Println("Getting names of the files in local")
	return []string{"a", "b"}
}

type S3ReaderStrategy struct {
	Password string
}

func (s3 *S3ReaderStrategy) ReadFile(fileName string) *os.File {
	fmt.Println(fileName + "Reading file in S3")

	return nil
}

func (s3 *S3ReaderStrategy) ReadFiles() {
	fmt.Println("Reading files in S3")
}

func (s3 *S3ReaderStrategy) GetFilesNameList() []string {
	fmt.Println("Getting names of the files in S3")
	return []string{"a", "b"}
}
