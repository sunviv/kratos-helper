package errors

import (
	"strconv"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Code int64 // 业务响应状态码

var codeMessageMap map[Code]string // 业务状态码映射表

func (c Code) Message() string {
	return codeMessageMap[c]
}

func (c Code) String() string {
	return strconv.FormatInt(int64(c), 10)
}

// 实现 error 接口
func (c Code) Error() string {
	return c.Message()
}

func (c Code) ToCodeError(errs ...error) CodeError {
	if len(errs) == 0 {
		return CodeError{c, c.Message(), c}
	}
	return CodeError{
		c,
		c.Message(),
		errs[0],
	}
}

// CodeError 同时具备业务状态码，和原始错误
type CodeError struct {
	Code           // 业务状态码（供前端使用）
	message string // 自定义 Message，默认为 Code 对应的 codeMessageMap 中的信息
	Err     error  // 原始错误（真正的错误，排错和日志记录）
}

func (ce CodeError) Message() string {
	if len(ce.message) > 0 {
		return ce.message
	}
	return ce.Code.Message()
}

// 自定义 Code Message
func (ce CodeError) WithMessage(msg string) CodeError {
	ce.message = msg
	return ce
}

// 实现 error 接口，返回原始错误信息
func (ce CodeError) Error() string {
	return ce.Err.Error()
}

func (ce CodeError) String() string {
	return ce.Code.Message()
}

// 实现 GRPCStatus() 接口
func (ce CodeError) GRPCStatus() *status.Status {
	s, _ := status.New(codes.Code(ce.Code), ce.Message()).WithDetails(&errdetails.ErrorInfo{
		Reason: ce.Error(),
	})
	return s
}

func Init(messageMap map[Code]string) {
	codeMessageMap = messageMap
}
