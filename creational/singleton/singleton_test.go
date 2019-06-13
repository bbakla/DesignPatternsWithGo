package singleton

import (
	"testing"
)

func TestSee(t *testing.T) {

	s := GetInstance()
	infoLog := Log {INFO, "first log", 0}
	s.Log(infoLog)

	a := GetInstance()
	warnLog := Log {WARN, "second log", 0}
	a.Log(warnLog)

}

