package dynamicCore

import (
	"errors"
	"mockTail/mockTail/internal/config"
	"mockTail/mockTail/internal/tools"
)

type Registry interface {
	Update()
	FindNode(request Request) (*Node, error)
	Find(request Request) (*FileResource, error)
}

type registry struct {
	root *Node
}

func NewRegistry() Registry {
	basePath := tools.ArgsVal.SourceDirectory + config.GetConfig().Api.EntryPoint
	return &registry{root: NewDirNode(basePath)}
}

func (r registry) Find(request Request) (*FileResource, error) {
	node, err := r.FindNode(request)
	if err != nil || node == nil {
		return nil, errors.New("resource not found")
	}
	return node.FindData(request)
}

func (r registry) FindNode(request Request) (*Node, error) {
	var node *Node
	for _, parentPath := range request.GetSegmentPath() {
		node = r.root.Find(parentPath)
		if node != nil && node.Info.Path != request.RawPath && !node.Info.IsDynamic {
			continue
		}
		if node != nil {
			// exact match / dynamic match
			break
		}
	}
	return node, nil
}

func (r registry) Update() {
	err := r.root.BuildTree()
	if err != nil {
		panic(err)
	}
}
