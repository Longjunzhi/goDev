package config

import (
	"Img/util"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type MysqlConfig struct {
	User     string `json:"user"`
	DB       string `json:"db"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type storageConfig struct {
	Path         string `json:"path"`
	AccessSecret string `json:"access_secret"`
}

type AppConfiguration struct {
	ENV         string
	BasePath    string
	ServerConf  ServerConfig
	MysqlConf   MysqlConfig
	storageConf storageConfig
}

var AppConf AppConfiguration

func init() {
	env := os.Getenv("ENV")
	AppConf.ENV = env
	setBasePath()
	viper.SetConfigName("config." + AppConf.ENV)
	viper.AddConfigPath(AppConf.BasePath + util.GetPathTag() + "config")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading config, %s", err))
	}
	err := viper.Unmarshal(&AppConf)
	if err != nil {
		panic(fmt.Errorf("unable to decode into appConf, %v", err))
	}
	logrus.Infof("init success: conf = %+v", AppConf)
}

func setBasePath() {
	basePath := ""
	if AppConf.ENV == "prod" {
		path, err := os.Executable()
		if err != nil {
			panic(err)
		}
		basePath = path
		basePath = filepath.Dir(basePath)
	}
	if AppConf.ENV == "dev" {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		basePath = path
	}
	fmt.Println("basePath", basePath)
	AppConf.BasePath = basePath
}
