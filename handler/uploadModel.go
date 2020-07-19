package handler

import (
	"fmt"
	"hku/wallpaper/db"
	"hku/wallpaper/define"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadModel(context *gin.Context) {
	model, _ := context.FormFile("model")
	modelName := context.PostForm("name")
	modelClass := context.PostForm("classify")
	uid, err := strconv.ParseInt(context.PostForm("uid"), 10, 64)
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"status":  -1,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)
	name := fmt.Sprintf("%v-%v-%v-%v.stl", uid, modelClass, modelName, i)
	modelPath := define.MODEL_PATH + name
	if err := context.SaveUploadedFile(model, modelPath); err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"status":  0,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	filePath := ""
	err = db.InsertModel(uid, modelPath, filePath, modelClass, modelName)
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"status":  0,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	context.JSON(200, gin.H{
		"status":  1,
		"message": "model upload successfully",
	})
	return
}
