package dynamicCore

import (
  "os"
  "path/filepath"
)

type Registry interface {
  Find(path string) (*SourceResource, error)
  Add()
  Get()
  Remove()
}

type registry struct {
  routes []SourceResource
}

func NewRegistry() Registry {
  return &registry{}
}

func (r *registry) Find(path string) (*SourceResource, error) {
  var target SourceResource
  target.Path = path

  err := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
    if err != nil {
      return err
    }
    if !d.IsDir() {
      parser := NewParser(d)
      res, err := NewResource(parser)
      if err != nil {
        return err
      }
      target.Resources = append(target.Resources, *res)
    }
    return nil
  })
  if err != nil {
    return nil, err
  }
  r.routes = append(r.routes, target)
  return &target, nil
}

func (r *registry) Add() {
}

func (r *registry) Get() {
}

func (r *registry) Remove() {
}
