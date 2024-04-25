package global

import "fmt"

const (
	SUCCESS = 0
	FAIL    = 500
)

var (
	_codes    = make(map[int]struct{}) // 空 set
	_messages = make(map[int]string)
)

type Result struct {
	code int
	msg  string
}

func GetMsg(code int) string {
	return _messages[code]
}

func (e Result) Msg() string {
	return e.msg
}

func (e Result) Code() int {
	return e.code
}

func RegisterResult(code int, msg string) Result {
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在， 请更换一个", code))
	}
	if msg == "" {
		panic("错误码不能为空")
	}

	_codes[code] = struct{}{}
	_messages[code] = msg

	return Result{
		code: code,
		msg:  msg,
	}
}

var (
	OkResult   = RegisterResult(SUCCESS, "OK")
	FailResult = RegisterResult(FAIL, "Fail")
)

var (
	ErrNotFound = RegisterResult(4003, "资源不存在")
	ErrDbOp     = RegisterResult(4004, "数据库操作异常")
	ErrRequest  = RegisterResult(4005, "请求参数异常")
)
