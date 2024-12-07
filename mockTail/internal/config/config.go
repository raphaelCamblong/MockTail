package config

import (
  "fmt"
  "github.com/sirupsen/logrus"
  "mockTail/mockTail/internal/tools"
  "sync"
)

type (
  Configuration struct {
    *AppConfig
  }
)

var (
  once           sync.Once
  configInstance *Configuration
)

func findConfigFile() (string, error) {
  path1 := fmt.Sprintf("%s/%s", tools.ArgsVal.SourceDirectory, tools.ArgsVal.ConfigPath)
  if err := tools.VerifyFileExistence(path1); err == nil {
    return path1, nil
  }
  path2 := fmt.Sprintf("%s", tools.ArgsVal.ConfigPath)
  if err := tools.VerifyFileExistence(path2); err == nil {
    return path2, nil
  }
  return "", fmt.Errorf("config file not found in %s or %s", path1, path2)
}

func GetConfig() *Configuration {
  once.Do(
    func() {
      path, err := findConfigFile()
      if err != nil {
        logrus.Errorf("Error finding config file: %v", err)
      }
      config := LoadYamlConfig(path)
      configInstance = &Configuration{
        config,
      }
    },
  )
  return configInstance
}
