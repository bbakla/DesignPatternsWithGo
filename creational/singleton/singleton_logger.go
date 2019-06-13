package singleton

import (
	. "github.com/logrusorgru/aurora"
	"fmt"
	"text/template"
	"bytes"
)


type Logger interface {
	Log(Log)
}

type logger struct {
	count int
}

type Log struct {
	LogLevel
	Message string
	CorrelationID int
}

var instance *logger

func GetInstance() Logger {
	if instance == nil {
		instance = new(logger)
	}
	return instance
}

func (s *logger) Log(log Log) {

	s.count++
	log.CorrelationID = s.count

	const levelString = "[{{.LogLevel}}]-{{.CorrelationID}}:   {{.Message}}"
	
	buffer := &bytes.Buffer{}
	t, err := template.New("t1").Parse(levelString)
	if err != nil {
		panic(err)
	}


	t.Execute(buffer, log)

	 switch l := log.LogLevel; l {

	 case INFO: 
		 fmt.Println(Bold(Cyan(buffer.String())))
	 case DEBUG:
		fmt.Println(Bold(Green(buffer.String())))
	 case WARN:
		fmt.Println(Bold(Yellow(buffer.String())))
	 case ERROR:
		fmt.Println(Bold(Red(buffer.String())))
	 }
}

type LogLevel string

const (
	INFO LogLevel = "INFO"
	DEBUG LogLevel = "DEBUG"
	WARN LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
)

