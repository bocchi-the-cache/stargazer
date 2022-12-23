package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sptuan/stargazer/internal/conf"
)

var defaultRouter *gin.Engine

func Init() {
	defaultRouter = NewRouter()
}

func Run() error {
	return defaultRouter.Run(fmt.Sprintf("%s:%d", conf.Cfg.Service.Http.Addr, conf.Cfg.Service.Http.Port))
}
