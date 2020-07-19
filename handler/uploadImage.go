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

func UploadImage(context *gin.Context) {
	file, _ := context.FormFile("img")
	imgName := context.PostForm("imageName")
	imageClass := context.PostForm("imageClass")
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
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)
	filename := fmt.Sprintf("%s_%v_%v.png", imgName, uid, i)
	filepath := define.IMAGE_PATH + filename
	if err := context.SaveUploadedFile(file, filepath); err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"status":  0,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	err, pid := db.SaveImg(filename, imgName, imageClass, uid)
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"status":  0,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	err = db.UpdateUserImg(uid, pid)
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
		"message": "image upload successfully",
	})
}
