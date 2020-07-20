package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hku/wallpaper/db"
	"strconv"
)


func FetchModel(context *gin.Context) {
	uid, err := strconv.ParseInt(context.PostForm("uid"), 10, 64)
	fmt.Println(uid)
	if err != nil {
		context.JSON(500, gin.H{
			"status": 0,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	models, err := db.FetchModel(uid)
	fmt.Printf("Finish")
	if err != nil {
		context.JSON(500, gin.H{
			"status": 0,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	context.JSON(200, gin.H{
		"status": 1,
		"message": "",
		"body": models,
	})
}
