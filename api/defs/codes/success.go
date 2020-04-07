/*
@Time : 07/04/2020 8:47 PM 
@Author : GC
*/
package codes

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
	SuccessDefault = SuccessResponse{StatusCode: 200, Data: SuccessData{Code: "10000", Msg: "success", Data: struct{}{}}}
)
