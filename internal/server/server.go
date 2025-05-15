package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Addr    string
	Handler http.Handler
}

// NewServer arma el http.Server con timeouts y un posible graceful-shutdown
func NewServer(handler http.Handler) *Server {
	return &Server{
		Addr:    ":8000",
		Handler: handler,
	}
}

// ListenAndServe inicia la escucha
func (s *Server) ListenAndServe() error {
	srv := &http.Server{
		Addr:         s.Addr,
		Handler:      s.Handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()

	return srv.ListenAndServe()
}
