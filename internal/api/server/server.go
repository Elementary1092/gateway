package server

import (
	"fmt"
	"github.com/elem1092/gateway/internal/config"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func NewServer(cfg config.GatewayConfig, router *mux.Router) http.Server {
	return http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Address, cfg.Port),
		Handler:      router,
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
}
