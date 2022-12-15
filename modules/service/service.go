package service

import "github.com/gin-gonic/gin"

var defaultRouter *gin.Engine

func Init() {
	defaultRouter = gin.Default()
	defaultRouter.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
}

func Run() error {
	return defaultRouter.Run(":8080")
}
