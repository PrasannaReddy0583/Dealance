package httpx

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type Config struct {
	Addr string
}

func NewServer(cfg Config) *Server {
	router := applyMiddleware(NewRouter())

	return &Server{
		httpServer: &http.Server{
			Addr:         cfg.Addr,
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
