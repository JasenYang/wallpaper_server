package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hku/wallpaper/db"
)

func Register(context *gin.Context) {
	username := context.PostForm("name")
	password := context.PostForm("password")
	fmt.Println(username)
	fmt.Println(password)
	if db.CheckUser(username) {
		context.JSON(200, gin.H{
			"uid": -1,
			"status": 0,
			"message": "username is exist",
		})
		return
	}
	err, uid := db.InsertUser(username, password)
	if err != nil {
		fmt.Println("这儿也错了")
		fmt.Println(err)
		context.JSON(500, gin.H{
			"uid": -1,
			"status": 0,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	context.JSON(200, gin.H{
		"uid": uid,
		"status": 1,
		"message": "register successful",
	})
	return
}