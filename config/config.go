package config

import "github.com/spf13/viper"

type Config struct {
	Name string
}

func Init(name string) error {
	c := Config{Name: name}
	err := c.initConfig()
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigName(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}