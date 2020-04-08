/*
@Time : 07/04/2020 3:49 PM 
@Author : GC
*/
package user

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"video_server_api/dbops"
	"video_server_api/defs/codes"
	"video_server_api/utils"
)

type addReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户新增
func Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := &addReq{}
	if err := utils.ParseJsonRequest(r, user); err != nil {
		log.Println("user.Add:error:", err)
		utils.SendErrorResponse(w, codes.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUser(user.Username, user.Password); err != nil {
		log.Println("user.Add:error:", err)
		utils.SendErrorResponse(w, codes.ErrorDB)
		return
	}
	utils.SendNormalResponse(w, codes.SuccessDefault)
	return
}

type GetResp struct {
	Id int `json:"id"`
}

// 单用户
func Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("username")
	user, err := dbops.GetUser(name)
	if err != nil {
		log.Println("user.Get:error:", err)
		utils.SendErrorResponse(w, codes.ErrorDB)
		return
	}

	if user == nil {
		utils.SendNormalResponse(w, codes.SuccessEmptyEntity)
		return
	}

	utils.SendNormalResponse(w, codes.GenSuccessResponse(GetResp{Id: user.Id}))
	return
}
