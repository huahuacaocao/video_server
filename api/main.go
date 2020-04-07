/*
@Time : 06/04/2020 10:23 PM 
@Author : GC
*/
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server_api/handlers/user"
)
import (
	_ "github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("starting:")
	router := registerHandlers()
	http.ListenAndServe(":8888", router)
}

func registerHandlers() *httprouter.Router {
	router := httprouter.New()
	// user
	router.POST("/users", user.Add)

	//router.GET("/users/:username") // GetUserInfo
	//router.POST("/login")          // Login
	//router.GET("/users/:username") // GetUserInfo

	// user-videoinfo
	////router.POST("/user/:username/videos") // AddNewVideo
	//router.GET("/user/:username/videos")             // ListAllVideos
	//router.DELETE("/user/:username/videos/:videoid") // DeleteVideo

	// videoinfo
	//router.GET("/videos")
	//router.GET("/videos/videoid")

	// video-comment
	//router.POST("/videos/:vidoeid/comments") // PostComment
	//router.GET("/videos/:videoid/comments")  // ShowComments

	return router
}
