package handler

import (
	"github.com/gin-gonic/gin"
	"hku/wallpaper/db"
)

func Login(context *gin.Context) {
	username := context.PostForm("name")
	password := context.PostForm("password")
	uid := db.Validate(username, password)
	if uid == -1 {
		context.JSON(200, gin.H{
			"uid": -1,
			"status": 0,
			"message": "user not exists",
		})
		return
	}
	context.JSON(200, gin.H{
		"uid": uid,
		"status": 1,
		"message": "login successfully",
	})
}
