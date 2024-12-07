package internal

import (
  "errors"
  "github.com/sirupsen/logrus"
  "mockTail/mockTail/internal/dynamicCore"
  "mockTail/mockTail/internal/handlers"
  "mockTail/mockTail/internal/router"
)

type (
  Infra struct {
    Router router.Router
  }
  Server struct {
    Infra    Infra
    Core     dynamicCore.Core
    Handlers handlers.Handlers
  }
)

func NewServer() *Server {
  return &Server{}
}

func (server *Server) setInfra() {
  server.Infra.Router = router.NewRouter()
}

func (server *Server) setAll() {
  server.Handlers = *handlers.NewHandlers()
  server.Core = dynamicCore.NewCore(dynamicCore.NewRegistry())
  server.Handlers.Attach(server.Infra.Router, server.Core)
}

func (server *Server) HealthCheck() error {
  if server.Infra.Router == nil {
    return errors.New("failed to start router")
  }
  return nil
}

func (server *Server) init() {
  logrus.Info("Initializing infra...")
  server.setInfra()

  logrus.Info("Initializing dynamic core...")
  server.setAll()

  logrus.Info("Health check...")
  if err := server.HealthCheck(); err != nil {
    logrus.Error(err)
  }
}

func (server *Server) Run() {
  server.init()
  server.Infra.Router.Start()
}
