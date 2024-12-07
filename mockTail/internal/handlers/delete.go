package handlers

import (
  "github.com/gin-gonic/gin"
  "mockTail/mockTail/internal/dynamicCore"
)

type Delete struct {
  core dynamicCore.Core
}

func (method Delete) Process(c *gin.Context) error {
  c.JSON(200, gin.H{
    "error": "method not available",
  })
  return nil
}
