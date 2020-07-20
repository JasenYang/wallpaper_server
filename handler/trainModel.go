package handler

import (
	"archive/zip"
	"fmt"
	"hku/wallpaper/db"
	"hku/wallpaper/define"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func TrainModel(context *gin.Context) {
	file, _ := context.FormFile("file")
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
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)
	zipName := fmt.Sprintf("%s_%v_%v.zip", modelName, uid, i)
	zipPath := define.IMAGE_PATH + zipName
	if err := context.SaveUploadedFile(file, zipPath); err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"status":  0,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	filePath := define.IMAGE_PATH + "/" + strconv.FormatInt(uid, 10) + "/" + strconv.Itoa(i)
	err = DeCompress(zipPath, filePath)
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"status":  0,
			"message": "Depress error",
		})
		return
	}
	modelPath := GetModelPath(filePath)
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

func GetModelPath(filepath string) string {
	modelName := ""
	return modelName
}

//解压
func DeCompress(zipFile, dest string) (err error) {
	//目标文件夹不存在则创建
	if _, err = os.Stat(dest); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dest, 0755)
		}
	}

	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}

	defer reader.Close()

	for _, file := range reader.File {
		//    log.Println(file.Name)

		if file.FileInfo().IsDir() {

			err := os.MkdirAll(dest+"/"+file.Name, 0755)
			if err != nil {
				log.Println(err)
			}
			continue
		} else {

			err = os.MkdirAll(getDir(dest+"/"+file.Name), 0755)
			if err != nil {
				return err
			}
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		filename := dest + "/" + file.Name
		//err = os.MkdirAll(getDir(filename), 0755)
		//if err != nil {
		//    return err
		//}

		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()

		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		//w.Close()
		//rc.Close()
	}
	return
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
