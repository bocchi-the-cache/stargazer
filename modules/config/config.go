package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/sptuan/stargazer/modules/global"
	"path/filepath"
	"strings"
)

// Init Read config from file and parse it to ServerConfig
// TODO: provide zero config start up
func Init(configFile string) {
	path := filepath.Dir(configFile)
	fileWithoutExt := strings.TrimSuffix(filepath.Base(configFile), filepath.Ext(configFile))

	viper.SetConfigName(fileWithoutExt)
	if path != "" {
		fmt.Printf("start to read config from %s", path)
		viper.AddConfigPath(path)
	} else {
		fmt.Printf("start to read config from default path: config/config.yaml")
		viper.AddConfigPath("./")
	}

	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("read config failed: %s", err)
		panic(err)
	}
	err = viper.UnmarshalKey("server", &global.ServerConfig)
	if err != nil {
		fmt.Printf("parse sever config failed: %s", err)
		panic(err)

	}
	err = viper.UnmarshalKey("targets", &global.Targets)
	if err != nil {
		fmt.Printf("parse targets config failed: %s", err)
		panic(err)

	}

	fmt.Printf("config init success")
	return
}
