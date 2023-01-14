package rest

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func NewServer(port int, handler http.Handler) *Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	return &Server{
		srv: srv,
	}
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
