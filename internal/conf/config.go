package conf

import (
	"github.com/sptuan/stargazer/config"
	"github.com/sptuan/stargazer/pkg/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var cfg AllConfig

// Init Read config from file and parse. If file not exist, create it.
func Init(configFile string) {
	if configFile == "" {
		configFile = "config.yaml"
	}

	// if config not exist, create it
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		logger.Warn("Config file not exist, create it...")
		err = createConfigFile(configFile)
		if err != nil {
			logger.Panic("Create config file failed: ", err)
		}
	}

	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		logger.Panic("Read config file failed: ", err)
	}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		logger.Panic("Parse config file failed: ", err)
	}

	return
}

func createConfigFile(configFile string) error {
	return os.WriteFile(configFile, config.DefaultConfig, 0644)
}
