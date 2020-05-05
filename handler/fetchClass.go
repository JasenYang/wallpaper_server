package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hku/wallpaper/db"
	"strconv"
)

func FetchClass(context *gin.Context)  {
	uid, err := strconv.ParseInt(context.PostForm("uid"), 10, 64)
	if err != nil {
		context.JSON(500, gin.H{
			"status": 0,
			"message": fmt.Sprintf("%v", err),
			"classify": make([]string, 0),
		})
		return
	}
	result, err := db.FetchClass(uid)
	if err != nil {
		context.JSON(500, gin.H{
			"status": 0,
			"message": fmt.Sprintf("%v", err),
			"classify": make([]string, 0),
		})
		return
	}
	context.JSON(200, gin.H{
		"status": 1,
		"message": "",
		"classify": result,
	})
}
