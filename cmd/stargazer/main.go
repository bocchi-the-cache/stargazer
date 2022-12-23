package main

import (
	"flag"
	"github.com/sptuan/stargazer/internal/conf"
	"github.com/sptuan/stargazer/internal/service"
	"github.com/sptuan/stargazer/pkg/logger"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "config.yaml", "config file, default: {$pwd}/config.yaml")
}

func initConfig(configFilePath string) {
	conf.Init(configFilePath)
}

func initLogger() {
	// use default logger is enough
	//logger.Init()
}

func initService() {
	service.Init()
}

func main() {
	flag.Parse()

	initLogger()
	initConfig(configFile)
	initService()

	logger.Info("Project init complete. Start to run web service...")
	if err := service.Run(); err != nil {
		panic(err)
	}
}
