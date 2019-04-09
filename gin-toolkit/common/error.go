package common

import "fmt"

type Error struct {
	ErrorCode    int
	ErrorMsg     string
	ErrorUserMsg string
}

func NewError(code int, message string, userMessage string) Error {
	return Error{
		ErrorCode:    code,
		ErrorMsg:     message,
		ErrorUserMsg: userMessage,
	}
}

func (err Error) Error() string {
	return err.ErrorMsg
}

func SetErrPrintfMsg(err *Error, v ...interface{}) {
	err.ErrorMsg = fmt.Sprintf(err.ErrorMsg, v...)
}

func SetErrPrintfUserMsg(err *Error, v ...interface{}) {
	err.ErrorUserMsg = fmt.Sprintf(err.ErrorMsg, v...)
}

var HttpDecodeError = Error{1001, "http response code is not 0", ""}
var PanicError = Error{1002, "server have panic error", ""}
var MQTTPubError = Error{1003, "mqtt pub msg error", ""}
var HttpStatusError = Error{1004, "http status code is not 200", ""}
