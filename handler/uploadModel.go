package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hku/wallpaper/db"
	"strconv"
)

const (
	//IMAGE_PATH = "/Users/bytedance/TEMP/Image"
	MODEL_PATH = "/static/model/"
	// IMAGE_PATH = ""
	// MODEL_PATH = ""
)

func UploadModel(context *gin.Context) {
	file, _ := context.FormFile("model")
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

	//filename := file.Filename
	//zipName := fmt.Sprintf("%s_%v_%v.zip", modelName, uid, i)
	//zipPath := IMAGE_PATH + zipName
	modelPath := MODEL_PATH + modelName
	if err := context.SaveUploadedFile(file, modelPath); err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"status":  0,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	//filePath := IMAGE_PATH + "/" + strconv.FormatInt(uid, 10) + "/" + strconv.Itoa(i)
	//err = DeCompress(zipPath, filePath)
	//if err != nil {
	//	fmt.Println(err)
	//	context.JSON(500, gin.H{
	//		"status":  0,
	//		"message": "Depress error",
	//	})
	//	return
	//}

	err = db.InsertModel(uid, modelPath, "", modelClass, modelName)
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
