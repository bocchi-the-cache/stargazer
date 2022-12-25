package main

import (
	"flag"
	"github.com/sptuan/stargazer/internal/conf"
	"github.com/sptuan/stargazer/internal/db"
	"github.com/sptuan/stargazer/internal/service"
	"github.com/sptuan/stargazer/internal/task"
	"github.com/sptuan/stargazer/pkg/logger"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "config.yaml", "config file, default: {$pwd}/config.yaml")
}

func main() {
	flag.Parse()
	conf.Init(configFile)
	err := db.Init(conf.Cfg.Data.Database.Connection)
	if err != nil {
		logger.Panicf("failed to init database: %v", err)
	}
	task.Init()
	service.Init()

	logger.Info("Project init complete. Start to run web service...")

	if err := service.Run(); err != nil {
		panic(err)
	}
}
