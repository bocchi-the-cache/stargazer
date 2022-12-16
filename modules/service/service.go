package service

import "github.com/gin-gonic/gin"

var defaultRouter *gin.Engine

func Init() {
	defaultRouter = NewRouter()
}

func Run() error {
	// TODO: add config port here
	return defaultRouter.Run(":8080")
}
