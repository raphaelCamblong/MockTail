package config

import (
  "github.com/sirupsen/logrus"
  "github.com/spf13/viper"
  "os"
)

type (
  Api struct {
    EntryPoint string `mapstructure:"entryPoint"`
  }
  AppConfig struct {
    Host       string `mapstructure:"host"`
    Port       int    `mapstructure:"port"`
    Version    string `mapstructure:"version"`
    ConfigPath string `mapstructure:"configPath"`
    Api        Api    `mapstructure:"api"`
  }
)

func LoadYamlConfig(path string) *AppConfig {
  var config AppConfig

  viper.SetConfigFile(path)
  if err := viper.ReadInConfig(); err != nil {
    logrus.Errorf("Error config, %v", err)
    os.Exit(1)
  }
  if err := viper.Unmarshal(&config); err != nil {
    logrus.Errorf("Error config, %v", err)
    os.Exit(1)
  }
  return &config
}
