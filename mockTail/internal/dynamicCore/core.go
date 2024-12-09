package dynamicCore

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mockTail/mockTail/internal/config"
	"mockTail/mockTail/internal/tools"
)

type Core interface {
	GetRegistry() Registry
	Find(path string) (*FileResource, error)
	FindFromContext(ctx *gin.Context) (*FileResource, error)
}

type core struct {
	registry Registry
}

func NewCore(registry Registry) Core {
	return &core{
		registry: registry,
	}
}

func (c *core) GetRegistry() Registry {
	return c.registry
}

func (c *core) Find(path string) (*FileResource, error) {
	basePath := tools.ArgsVal.SourceDirectory + config.GetConfig().Api.EntryPoint
	completePath := basePath + path
	node, _ := c.registry.Find(NewRequest(completePath))

	if node == nil {
		return node, errors.New("resource not found")
	}
	return node, nil
}

func (c *core) FindFromContext(ctx *gin.Context) (*FileResource, error) {
	request := NewRequestFromContext(ctx)
	res, _ := c.registry.Find(request)

	if res == nil {
		return nil, errors.New("resource not found")
	}
	return res, nil
}
