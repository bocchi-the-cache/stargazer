package service

import (
	"fmt"
	"github.com/bocchi-the-cache/stargazer/internal/conf"
	"github.com/bocchi-the-cache/stargazer/internal/model"
	"github.com/gin-gonic/gin"
)

var defaultRouter *gin.Engine

func Init() {
	defaultRouter = NewRouter()
}

func Run() error {
	// set gin debug mode
	if model.Level(conf.Cfg.Service.LogLevel) == model.DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	return defaultRouter.Run(fmt.Sprintf("%s:%d", conf.Cfg.Service.Http.Addr, conf.Cfg.Service.Http.Port))
}
