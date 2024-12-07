package handlers

import (
  "github.com/gin-gonic/gin"
  "mockTail/mockTail/internal/config"
  "mockTail/mockTail/internal/dynamicCore"
  "mockTail/mockTail/internal/tools"
)

type Get struct {
  core dynamicCore.Core
}

func (method Get) Process(ctx *gin.Context) error {
  reqPath := ctx.Request.URL.Path
  basePath := tools.ArgsVal.SourceDirectory + config.GetConfig().Api.EntryPoint
  completePath := basePath + reqPath

  _, err := method.core.GetResource(completePath)
  if err != nil {
    ctx.JSON(404, gin.H{"error": err.Error()})
    return nil
  }
  ctx.JSON(200, gin.H{"data": "success"})
  return nil
}
