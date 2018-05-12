package common

import "encoding/json"

type SelfResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

/**
* 错误的信息构造方法
* code: 错误码
* msg:  错误信息
 */
func NewResponseSimple(code int, msg interface{}) *SelfResponse {
	return &SelfResponse{code, msg, nil}
}

/**
* 成功的信息构造方法
* data: 数据
 */
func NewResponseData(data interface{}, msg string) *SelfResponse {
	return &SelfResponse{0, msg, data}
}

func (this *SelfResponse) Encode() []byte {
	s, _ := json.MarshalIndent(*this, " ", " ")
	return s
}
