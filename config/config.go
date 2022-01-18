package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file %s \n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		fmt.Println("Config file changed")
	})

	return nil
}
