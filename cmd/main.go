package main

import (
	"github.com/elem1092/gateway/internal/api/handlers"
	"github.com/elem1092/gateway/internal/api/server"
	"github.com/elem1092/gateway/internal/config"
	"github.com/elem1092/gateway/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Starting API gateway")

	cfg := config.GetConfiguration()

	router := handlers.InitializeRouter(logger, cfg)

	logger.Info("Creating server")
	srv := server.NewServer(cfg.GatewayCfg, router)

	logger.Fatalf("Got error from server: %v", srv.ListenAndServe())
}
