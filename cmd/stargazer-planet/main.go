package main

import (
	"flag"
	"github.com/sptuan/stargazer/modules/config"
	"github.com/sptuan/stargazer/modules/logger"
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

func main() {
	flag.Parse()

	// TODO: Create default config file when not existed
	initConfig(configFile)
	initLogger()

	logger.Info("Test project init")
}
