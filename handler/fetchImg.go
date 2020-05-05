package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hku/wallpaper/db"
	"strconv"
)

func FetchImage(context *gin.Context) {
	imgClass := context.PostForm("imageClass")
	uid, err := strconv.ParseInt(context.PostForm("uid"), 10, 64)
	if err != nil {
		context.JSON(500, gin.H{
			"status": 0,
			"message": fmt.Sprintf("%v", err),
			"filename": "",
		})
		return
	}
	result, err := db.FetchImg(uid, imgClass)
	if err != nil {
		context.JSON(500, gin.H{
			"status": 0,
			"message": fmt.Sprintf("%v", err),
			"filename": "",
		})
		return
	}
	context.JSON(200, gin.H{
		"status": 1,
		"message": "",
		"filename": result,
	})
}
