package handlers

import "github.com/gin-gonic/gin"

type IMethod interface {
  Process(c *gin.Context) error
}
