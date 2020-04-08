/*
@Time : 07/04/2020 8:47 PM 
@Author : GC
*/
package codes

import "net/http"

// 定义正确 response
type SuccessData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type SuccessResponse struct {
	StatusCode int
	Data       SuccessData
}

var (
	SuccessDefault     = SuccessResponse{StatusCode: http.StatusOK, Data: SuccessData{Code: "10000", Msg: "success", Data: struct{}{}}}
	SuccessEmptyEntity = SuccessResponse{StatusCode: http.StatusOK, Data: SuccessData{Code: "10000", Msg: "success", Data: struct{}{}}}
)

func GenSuccessResponse(data interface{}) SuccessResponse {
	// @TODO 将 code 和 传进来
	return SuccessResponse{
		StatusCode: http.StatusOK,
		Data: SuccessData{
			Code: "10000",
			Msg:  "success",
			Data: data,
		},
	}
}
