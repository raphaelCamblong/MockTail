package tools

import (
  "fmt"
  "github.com/spf13/cobra"
  "os"
)

type Args struct {
  SourceDirectory string
  ConfigPath      string
  LogLevel        string
}

var ArgsVal Args

var RootCmd = &cobra.Command{
  Use:   "mockTail",
  Short: "MockTail is a mock server for testing APIs with mock data",
  RunE: func(cmd *cobra.Command, args []string) error {
    sourceDirectory, _ := cmd.Flags().GetString("source-directory")

    if err := VerifyFileExistence(sourceDirectory); err != nil {
      return err
    }
    return nil
  },
}

func init() {
  RootCmd.Flags().StringVarP(&ArgsVal.SourceDirectory, "source-directory", "s", "./", "Path to the source directory")
  RootCmd.Flags().StringVarP(&ArgsVal.ConfigPath, "config-path", "c", "./config.yml", "Path to the configuration file")
  RootCmd.Flags().StringVarP(&ArgsVal.LogLevel, "log-level", "l", "info", "Log level (debug, info, warn, error)")
}

func VerifyFileExistence(path string) error {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    return fmt.Errorf("entry point file does not exist: %s", path)
  }
  return nil
}
