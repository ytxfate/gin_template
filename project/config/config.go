package config

import (
	"github.com/spf13/viper"
)

var Cfg *Config

type Web struct {
	IsProdEnv     bool   `yaml:"isProdEnv"`
	Title         string `yaml:"title"`
	Description   string `yaml:"description"`
	Addr          string `yaml:"addr"`
	SecretKey     string `yaml:"secretKey"`
	Version       string `yaml:"version"`
	ApiPrefixPath string `yaml:"apiPrefixPath"`
}

type Config struct {
	Web Web `yaml:"web"`
}

func InitConfig() (err error) {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("../../") // 查找配置文件所在的路径
	viper.AddConfigPath(".")      // 还可以在工作目录中查找配置
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Cfg)
	return err
}
