package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type EnvConfig struct {
	Home       string
	Os         string
	ConfigPath string
}

type ServerConfig struct {
	Port        uint32 `mapstructure:"port"`
	AppName     string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	RoutePrefix string `mapstructure:"route-prefix"`
}

type TransportConfig struct {
	ServerPort int `mapstructure:"server-port"`
	ClientPort int `mapstructure:"client-port"`
}

type Config struct {
	Env       EnvConfig
	Server    ServerConfig    `mapstructure:"server"`
	Log       LogConfig       `mapstructure:"log"`
	SysMysql  SysMysql        `mapstructure:"sys_mysql"`
	Transport TransportConfig `mapstructure:"transport"`
}

func InitConfig() *Config {
	config := &Config{}

	config.Env.Home = getHome()
	config.Env.Os = runtime.GOOS

	configFilePath := flag.String("c", "", "Path to the configuration file.")

	if *configFilePath != "" {
		if ok, err := isFile(*configFilePath); err != nil {
			panic(err)
		} else if !ok {
			panic(fmt.Errorf("%s is not a file", *configFilePath))
		}
	} else {
		*configFilePath = filepath.Join(config.Env.Home, "system.toml")
		//fmt.Printf("load default config: %s\n", *configFilePath)
	}
	config.Env.ConfigPath = *configFilePath

	v := viper.New()

	v.SetConfigFile(*configFilePath)
	v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//TODO 监听配置文件变化

	err := v.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unmarshal config err: %v\n", err)
		panic(err)
	}

	return config
}

func getHome() string {
	dir, _ := os.Getwd()
	return dir
}

func isFile(path string) (bool, error) {
	file, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, errors.New("file is not exist")
		}
		return false, err
	}
	return !file.IsDir(), nil
}
