package config

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type JaegerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	Port       int          `mapstructure:"port" json:"port"`
	Tags       []string     `mapstructure:"tags" json:"tags"`
	JWTInfo    JWTConfig    `mapstructure:"jwt" json:"jwt"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
	JaegerInfo JaegerConfig `mapstructure:"consul" json:"jaeger"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error %s \n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		zap.L().Info("Config file changed")
	})

	return nil
}
