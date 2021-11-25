package rest

import (
	"context"
	"fibonachi/internal/config"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	cfg := config.New()

	s.httpServer = &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      handler,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	fmt.Println("Rest server started successfully")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
