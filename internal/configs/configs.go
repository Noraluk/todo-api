package configs

import "github.com/spf13/viper"

type Config struct {
	Port     string   `mapstructure:"port"`
	Database database `mapstructure:"database"`
}

type (
	database struct {
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		Port     string `mapstructure:"port"`
	}
)

const (
	configType = "yaml"
	configName = "configs"
)

var Conf *Config

func (c *Config) New(path string) error {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigType(configType)
	v.SetConfigName(configName)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&Conf); err != nil {
		return err
	}

	return nil
}
