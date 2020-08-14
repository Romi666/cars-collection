package Config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var (
	Logger   logrus.Logger
	CfgViper *viper.Viper
)

type Config struct {
	configFilePath string
}

func NewConfig(configFilePath string) *Config{
	return &Config{configFilePath: configFilePath}
}

func(c *Config) Config() *viper.Viper{
	viper.SetConfigFile(c.configFilePath)
	if err := viper.ReadInConfig(); err!=nil{
		panic(err)
	}

	cfg := viper.GetViper()
	return cfg
}
func (c *Config) Logger() {
	level := CfgViper.GetString("log_level")
	logFormat := new(logrus.JSONFormatter)
	var logLevel, err = logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	Logger = logrus.Logger{
		Out:       os.Stderr,
		Formatter: logFormat,
		Level:     logLevel,
	}
}

func (c *Config) GetConfigValue(key string) string{
	return c.Config().GetString(key)
}