package handlers

import (
	"github.com/gin-gonic/gin"
	"mockTail/mockTail/internal/dynamicCore"
	"net/http"
)

type Get struct {
	core dynamicCore.Core
}

func (method Get) Process(ctx *gin.Context) error {
	res, err := method.core.FindFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return nil
	}

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	//ctx.Header("Content-Disposition", "attachment; filename="+res.CompletePath)
	ctx.Header("Content-Type", "application/json")
	ctx.File(res.CompletePath)

	return nil
}
