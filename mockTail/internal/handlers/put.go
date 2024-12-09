package handlers

import (
	"github.com/gin-gonic/gin"
	"mockTail/mockTail/internal/dynamicCore"
)

type Put struct {
	core dynamicCore.Core
}

func (method Put) Process(ctx *gin.Context) error {
	ctx.JSON(200, gin.H{"data": "success"})
	return nil
}
