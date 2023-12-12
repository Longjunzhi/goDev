package config

import (
	"Img/databases"
	"Img/util"
	"errors"
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

type StorageConfig struct {
	Path         string `json:"path"`
	Dir          string `json:"dir"`
	AccessSecret string `json:"access_secret"`
}

type AppConfiguration struct {
	ENV         string
	BasePath    string
	ServerConf  ServerConfig
	MysqlConf   MysqlConfig
	StorageConf StorageConfig
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
	AppConf.StorageConf.Path = AppConf.BasePath + util.GetPathTag() + AppConf.StorageConf.Dir + util.GetPathTag()
	logrus.Infof("init success: conf = %+v", AppConf)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		AppConf.MysqlConf.User,
		AppConf.MysqlConf.Password,
		AppConf.MysqlConf.Host,
		AppConf.MysqlConf.Port,
		AppConf.MysqlConf.DB)
	err = databases.InitDatabase(dsn)

	if err != nil {
		logrus.Errorf("init mysql config: %v, err: %v", AppConf.MysqlConf, err)
		panic(errors.New("init mysql config fail"))
	}
}

func setBasePath() {
	basePath := ""
	if AppConf.ENV != "dev" {
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
