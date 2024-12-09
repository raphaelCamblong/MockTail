package dynamicCore

import (
	"github.com/gin-gonic/gin"
	"mockTail/mockTail/internal/tools"
	"mockTail/pkg/myArray"
	"mockTail/pkg/utils"
)

type Request struct {
	RawPath     string
	SegmentPath []string
	Method      string
	Name        string
}

func NewRequest(rawPath string) Request {
	return Request{
		RawPath: utils.CleanPath(rawPath),
	}
}

func NewRequestFromContext(ctx *gin.Context) Request {
	reqPath := ctx.Request.URL.Path
	basePath := tools.ArgsVal.SourceDirectory
	return Request{
		RawPath: utils.CleanPath(basePath, reqPath),
	}
}

func (r Request) GetSegmentPath() []string {
	if r.SegmentPath == nil {
		r.SegmentPath = utils.GetParentPaths(r.RawPath)
		r.SegmentPath = myArray.Reverse(r.SegmentPath)
	}
	return r.SegmentPath
}
