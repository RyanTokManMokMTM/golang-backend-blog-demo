package errcode

import (
	"fmt"
	"net/http"
)

//Error Response format ,All error Message using this format
type Error struct {
	code   int      `json:"code"`
	msg    string   `json:"msg"`
	detail []string `json:"detail"`
}

var codes = map[int]string{}

//NewError Custom Error Code
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		//Not allow duplicating Code
		panic(fmt.Sprintf("Error %d is already exist, please try another error", code))
	}

	codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

//Error Implement Error interface
func (e *Error) Error() string {
	return fmt.Sprintf("Error:%d,message:%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Detail() []string {
	return e.detail
}

//WithDetail set detail to error
func (e *Error) WithDetail(details ...string) *Error {
	newErr := *e
	newErr.detail = []string{}
	for _, d := range details {
		newErr.detail = append(newErr.detail, d)
	}
	return &newErr
}

//StatusCode According to custom status code return related HTTP Status Code,easy to manager the server
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code():
		fallthrough //next case
	case UnauthorizedTokenError.Code():
		fallthrough //next case
	case UnauthorizedTokenGenerateError.Code():
		fallthrough //next case
	case UnauthorizedTokenTimeOut.Code():
		return http.StatusUnauthorized
	case TooManyRequest.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError //not such case
}
