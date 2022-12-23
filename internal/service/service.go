package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sptuan/stargazer/internal/conf"
	"github.com/sptuan/stargazer/internal/constant"
)

var defaultRouter *gin.Engine

func Init() {
	defaultRouter = NewRouter()
}

func Run() error {
	// set gin debug mode
	if constant.Level(conf.Cfg.Service.LogLevel) == constant.DebugLevel {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	return defaultRouter.Run(fmt.Sprintf("%s:%d", conf.Cfg.Service.Http.Addr, conf.Cfg.Service.Http.Port))
}
