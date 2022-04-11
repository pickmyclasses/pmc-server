package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"go.uber.org/zap"
)

// JWTConfig defines the configuration for JWT
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"` // Signing Key is the key for signing the JWT signature (PK)
}

// JaegerConfig defines the configuration for Jaeger
type JaegerConfig struct {
	Host string `mapstructure:"host" json:"host"` // Host of the server
	Port int    `mapstructure:"port" json:"port"` // Port of the server
	Name string `mapstructure:"name" json:"name"` // Name of the server
}

// ConsulConfig defines the configuration for Consul distributed config center
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"` // Host of Consul
	Port int    `mapstructure:"port" json:"port"` // Port of Consul
}

// ServerConfig defines the configuration for this server
// Basically this should include everything that we will be using
// This should be removed in the future to be replaced command line for security concerns
type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`     // Name of this sever
	Host       string       `mapstructure:"host" json:"host"`     // Host of this server
	Port       int          `mapstructure:"port" json:"port"`     // Port of this server
	Tags       []string     `mapstructure:"tags" json:"tags"`     // Tags used in Consul (will only be in Consul)
	JWTInfo    JWTConfig    `mapstructure:"jwt" json:"jwt"`       // JWT config of the server
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"` // Consul configurations
	JaegerInfo JaegerConfig `mapstructure:"consul" json:"jaeger"` // Jaeger configurations
}

// NacosConfig defines the configuration for Nacos
type NacosConfig struct {
	Host      string `mapstructure:"host"`      // Host of the Nacos server
	Port      uint64 `mapstructure:"port"`      // Port of the Nacos server
	Namespace string `mapstructure:"namespace"` // Namespace of this sever in Nacos
	User      string `mapstructure:"user"`      // Username we will be using for Nacos
	Password  string `mapstructure:"password"`  // Password we will be using for Nacos
	DataId    string `mapstructure:"dataid"`    // DataId for fetching certain database ID
	Group     string `mapstructure:"group"`     // The group of the database server
}

// Init initialize Viper for reading configuration files from Consul or local file
// This should be replaced by command line in the future for security purpose
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
