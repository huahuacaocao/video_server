/*
@Time : 07/04/2020 4:11 PM 
@Author : GC
*/
package utils

import (
	"encoding/json"
	"net/http"
	"video_server_api/defs/codes"
)

// 响应错误信息
func SendErrorResponse(w http.ResponseWriter, response codes.ErrorResponse) {
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(response.StatusCode)
	jsonBytes, _ := json.Marshal(&response.Error)
	w.Write(jsonBytes)
}

// 响应正确信息
func SendNormalResponse(w http.ResponseWriter, response codes.SuccessResponse) {
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(response.StatusCode)
	jsonBytes, _ := json.Marshal(&response.Data)
	w.Write(jsonBytes)
}
