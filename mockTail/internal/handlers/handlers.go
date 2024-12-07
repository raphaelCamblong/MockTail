package handlers

import (
  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
  "mockTail/mockTail/internal/dynamicCore"
  "mockTail/mockTail/internal/router"
)

type Handlers struct {
  methods map[string]IMethod
}

func NewHandlers() *Handlers {
  return &Handlers{}
}

func (handlers *Handlers) Attach(r router.Router, core dynamicCore.Core) {
  handlers.methods = make(map[string]IMethod)
  handlers.methods["GET"] = Get{core: core}
  handlers.methods["POST"] = Post{core: core}
  handlers.methods["PUT"] = Put{core: core}
  handlers.methods["DELETE"] = Delete{core: core}
  handlers.methods["OPTIONS"] = Default{}
  handlers.methods["DEFAULT"] = Default{}
  handlers.methods["PATCH"] = Default{}
  handlers.methods["HEAD"] = Default{}
  handlers.methods["CONNECT"] = Default{}
  handlers.methods["TRACE"] = Default{}

  r.Get().NoRoute(handlers.dispatch)
}

func (handlers *Handlers) findMethod(ctx *gin.Context) IMethod {
  method := ctx.Request.Method
  if _, ok := handlers.methods[method]; !ok {
    method = "DEFAULT"
  }
  return handlers.methods[method]
}

func (handlers *Handlers) dispatch(ctx *gin.Context) {
  method := handlers.findMethod(ctx)
  if err := method.Process(ctx); err != nil {
    logrus.Errorf("Error processing request: %v", err)
  }
}
