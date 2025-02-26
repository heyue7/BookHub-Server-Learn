package srv

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Response struct {
	Code    int32         `json:"code"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data,omitempty"`
	Errors  []interface{} `json:"-"`
}

func (rsp *Response) String() string {
	buf, err := json.Marshal(rsp)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

func Success(data interface{}) *Response {
	if data == nil {
		data = map[string]interface{}{}
	}
	return &Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	}
}

func Error(code int32, message string, errs ...interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Errors:  errs,
	}
}

func ErrorWithValidate(err error) *Response {
	var v *json.UnmarshalTypeError
	if errors.As(err, &v) {
		return Error(7001, fmt.Sprintf("请求参数 %s 的类型是 %s, 不是 %s", v.Field, v.Type, v.Value))
	}
	return Error(7004, "参数错误", err.Error())
}
