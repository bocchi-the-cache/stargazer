package main

import (
	"flag"
	"github.com/sptuan/stargazer/modules/config"
	"github.com/sptuan/stargazer/modules/logger"
	"github.com/sptuan/stargazer/modules/service"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "config.yaml", "config file, default: {$pwd}/config.yaml")
}

func initConfig(configFilePath string) {
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}
	config.Init(configFilePath)
}

func initLogger() {
	logger.Init()
}

func initService() {
	service.Init()
}

func main() {
	flag.Parse()

	initConfig(configFile)
	initLogger()
	initService()

	logger.Info("Project init complete. Start to run web service...")
	if err := service.Run(); err != nil {
		panic(err)
	}
}
