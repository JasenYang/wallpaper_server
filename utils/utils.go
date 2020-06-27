package utils

import "log"

type CommonResponse struct {
	Code int64                  `json:"code"`
	Msg  string                 `json:"msg"`
	Data interface{} 			`json:"data"`
}

var (
	// code
	CODE_SUCCESS        int64 = int64(1)
	//CODE_FAILED         int64 = int64(1)
	//CODE_PARAMERR       int64 = int64(2)
	//CODE_NOTEXIST       int64 = int64(3)
	//CODE_FORBIDDEN      int64 = int64(4)
	//CODE_LOGIN_REQUIRED int64 = int64(5)
	//CODE_TIMEOUT        int64 = int64(6)
)

var (
	msgMap map[int64]string = make(map[int64]string)
)

func NewCommonResponse(code int64) *CommonResponse {
	msg, ok := msgMap[code]
	if ok != true {
		log.Printf("code not installed code:%d", code)
		return nil
	}
	rsp := &CommonResponse{
		Code: code,
		Msg:  msg,
		Data: make(map[string]interface{}),
	}

	return rsp
}