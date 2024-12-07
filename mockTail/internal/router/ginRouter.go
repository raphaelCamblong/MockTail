package router

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
  "mockTail/mockTail/internal/config"
  "mockTail/mockTail/internal/router/middleware"
)

type GinRouter struct {
  *gin.Engine
}

func NewRouter() Router {
  r := gin.Default()
  r.Use(middleware.CORSMiddleware())

  return &GinRouter{r}
}

func (r *GinRouter) Start() {
  c := config.GetConfig()
  addr := fmt.Sprintf("%s:%d", c.Host, c.Port)

  err := r.Run(addr)
  if err != nil {
    logrus.Errorf("failed to initiate gin router: %v", err)
  }
}

func (r *GinRouter) Get() *gin.Engine {
  return r.Engine
}
