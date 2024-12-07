package dynamicCore

type Core interface {
  GetResource(path string) (*SourceResource, error)
}

type core struct {
  registry Registry
}

func NewCore(registry Registry) Core {
  return &core{
    registry: registry,
  }
}

func (c *core) GetResource(path string) (*SourceResource, error) {
  return c.registry.Find(path)
}
