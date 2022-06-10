package server

import (
	"fmt"
	"github.com/elem1092/gateway/internal/config"
	"net/http"
	"time"
)

func NewServer(cfg config.GatewayConfig, router http.Handler) http.Server {
	return http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Address, cfg.Port),
		Handler:      router,
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
}
