package handlers

import (
  "github.com/gin-gonic/gin"
  "mockTail/mockTail/internal/dynamicCore"
)

type Post struct {
  core dynamicCore.Core
}

func (method Post) Process(c *gin.Context) error {
  c.JSON(200, gin.H{
    "error": "method not available",
  })
  return nil
}
