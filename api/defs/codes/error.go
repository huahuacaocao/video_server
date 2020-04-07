/*
@Time : 07/04/2020 4:13 PM 
@Author : GC
*/
package codes

// 定义错误 response
type Error struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	StatusCode int
	Error      Error
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{StatusCode: 400, Error: Error{Error: "Request body is not correct.", ErrorCode: "10001"}}
	ErrorNotAuthUser            = ErrorResponse{StatusCode: 401, Error: Error{Error: "User anthentication failed.", ErrorCode: "10002"}}
	// @I: 思路1: 定义多种错误类型, 在 http response 时候转换和统一,将 http 状态码统一(500),将 subcode 和 message 统一(不提示 DB 错误), 如下 eg.
	//ErrorDB = ErrorResponse{StatusCode: 500, Error: Error{Error: "Internal service error", ErrorCode: "004"}}
	// @I: 思路2: 定义多种错误类型, 在 http response 时候转换和统一,将 http 状态码统一(500), subcode 和 message 提示具体信息 eg.
	ErrorDB             = ErrorResponse{StatusCode: 500, Error: Error{Error: "DB ops failed", ErrorCode: "10003"}}
	ErrorInternalFaults = ErrorResponse{StatusCode: 500, Error: Error{Error: "Internal service error", ErrorCode: "10004"}}
)
