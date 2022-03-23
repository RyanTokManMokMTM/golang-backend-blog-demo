package errcode

import (
	"fmt"
	"net/http"
)

//Error Response format ,All error Message using this format
type Error struct {
	Code   int      `json:"ErrorCode"`
	Msg    string   `json:"ErrorMsg"`
	Detail []string `json:"ErrorDetail"`
}

var codes = map[int]string{}

//NewError Custom Error ErrorCode
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		//Not allow duplicating ErrorCode
		panic(fmt.Sprintf("Error %d is already exist, please try another error", code))
	}

	codes[code] = msg
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

//Error Implement Error interface
func (e *Error) Error() string {
	return fmt.Sprintf("Error:%d,message:%s", e.ErrorCode(), e.ErrorMsg())
}

func (e *Error) ErrorCode() int {
	return e.Code
}

func (e *Error) ErrorMsg() string {
	return e.Msg
}

func (e *Error) ErrorMsgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *Error) ErrorDetail() []string {
	return e.Detail
}

//WithDetail set ErrorDetail to error
func (e *Error) WithDetail(details ...string) *Error {
	newErr := *e
	newErr.Detail = []string{}
	for _, d := range details {
		newErr.Detail = append(newErr.Detail, d)
	}
	return &newErr
}

//StatusCode According to custom status ErrorCode return related HTTP Status ErrorCode,easy to manager the server
func (e *Error) StatusCode() int {
	switch e.ErrorCode() {
	case Success.ErrorCode():
		return http.StatusOK
	case ServerError.ErrorCode():
		return http.StatusInternalServerError
	case InvalidParams.ErrorCode():
		return http.StatusBadRequest
	case NotFound.ErrorCode():
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.ErrorCode():
		fallthrough //next case
	case UnauthorizedTokenError.ErrorCode():
		fallthrough //next case
	case UnauthorizedTokenGenerateError.ErrorCode():
		fallthrough //next case
	case UnauthorizedTokenTimeOut.ErrorCode():
		return http.StatusUnauthorized
	case TooManyRequest.ErrorCode():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError //not such case
}
