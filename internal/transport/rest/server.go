package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/learningPlatform/internal/configs"
)

type Server struct {
	srv    *http.Server
	config *configs.Config
}

func NewServer(handler http.Handler, config *configs.Config) *Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &Server{
		srv:    srv,
		config: config,
	}
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
