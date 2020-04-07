/*
@Time : 07/04/2020 6:24 PM 
@Author : GC
*/
package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// json 解析
func ParseJsonRequest(r *http.Request, req interface{}) error {
	// @I: 查了一些源码, 发现在校验 json 时, go 偏向于不校验 Content-Type
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(buf, req); err != nil {
		return err
	}
	return nil
}
