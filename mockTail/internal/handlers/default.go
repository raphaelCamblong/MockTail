package handlers

import (
  "github.com/gin-gonic/gin"
)

type Default struct{}

func (method Default) Process(c *gin.Context) error {
  return nil
}
