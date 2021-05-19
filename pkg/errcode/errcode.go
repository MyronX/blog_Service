package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code	int			`json:"code"`
	msg		string		`json:"msg"`
	details	[]string	`json:"details"`
}

//map 对象，里面是一个string的数组 应该是放多个json文件
var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

//表示这个函数可以被e这个类型的调用
func (e *Error) Error()	string {
	return fmt.Sprintf("错误码：%d,错误信息：：%s", e.Code(), e.Msg())
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

func (e *Error) Details() []string {
	return e.details
}

//更新这个细节  实际上就是补充说明
func (e *Error) WithDetails(details ...string) *Error {
	NewError := *e
	NewError.details = []string{}
	for _,d := range details {
		NewError.details = append(NewError.details, d)
	}
	return &NewError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}