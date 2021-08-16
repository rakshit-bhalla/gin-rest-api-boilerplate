package errors

import (
	"net/http"
	"runtime"
)

type Error struct {
	Msg  string
	File string
	Func string
	Line int
}

func whereWasI(depthList ...int) (string, string, int) {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return file, runtime.FuncForPC(function).Name(), line
}

func CreateError(msg string) Error {
	function, file, line := whereWasI(2)
	return Error{
		Msg:  msg,
		File: file,
		Func: function,
		Line: line,
	}
}

var errors = initErrors()

func initErrors() map[UserError]int {
	m := make(map[UserError]int)
	m[ErrUserNotFound] = http.StatusNotFound
	return m
}

func GetHttpStatus(err *Error) *int {
	if nil != err {
		for key, value := range errors {
			if err.Msg == key.Message() {
				return &value
			}
		}
	}
	return nil
}
