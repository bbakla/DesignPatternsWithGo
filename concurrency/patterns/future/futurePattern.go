package future

import "fmt"

type MaybeString struct {
	successFunction SuccessFunc
	failFunc        FailFunc
}

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

func (s *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunction(str)
		}
	}(s)
}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunction = f

	return s
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%s Closure!\n", msg)

	return func() (string, error) {
		return msg, nil
	}
}
