package main

import (
  "github.com/sirupsen/logrus"
  "mockTail/mockTail/internal"
  "mockTail/mockTail/internal/tools"
  "os"
)

func main() {
  if err := tools.RootCmd.Execute(); err != nil {
    logrus.Errorf("Error executing command: %v", err)
    os.Exit(-1)
  }
  internal.NewServer().Run()
}
