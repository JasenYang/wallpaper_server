package main

import (
	"fmt"
	"hku/wallpaper/db"
	"hku/wallpaper/define"
	"hku/wallpaper/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitSQLiteDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	engine.Any("/", WebRoot)
	engine.POST("/user/register", handler.Register)
	engine.POST("/user/login", handler.Login)
	engine.POST("/image/upload", handler.UploadImage)
	engine.POST("/image/fetch", handler.FetchImage)
	engine.POST("/class/fetch", handler.FetchClass)
	// 上传stl文件
	engine.POST("/model/upload", handler.UploadModel)
	// 上传图片压缩包用于训练模型
	engine.POST("/model/train", handler.TrainModel)
	engine.POST("/model/fetch", handler.FetchModel)
	engine.StaticFS("/static", http.Dir(define.PATH))
	// 绑定端口，然后启动应用
	engine.Run("0.0.0.0:6789")
}

/**
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
 */
func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
