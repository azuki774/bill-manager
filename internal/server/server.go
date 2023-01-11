package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type APIService interface{}
type Server struct {
	Port   string
	Logger *zap.Logger
	APISvc APIService
}

func (s *Server) addRouting(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { // GET /
			w.Write([]byte("OK"))
		})
	})
	r.Use(s.middlewareLogging)
}

func (s *Server) Start(ctx context.Context) (err error) {
	s.Logger.Info("server start")

	addr := fmt.Sprintf(":%s", s.Port)

	r := chi.NewRouter()
	s.addRouting(r)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	ctxIn, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var errCh = make(chan error)
	go func() {
		errCh <- server.ListenAndServe()
	}()

	<-ctxIn.Done()
	if nerr := server.Shutdown(ctx); nerr != nil {
		s.Logger.Error("failed to shutdown server", zap.Error(nerr))
		return nerr
	}

	err = <-errCh
	if err != nil && err != http.ErrServerClosed {
		s.Logger.Error("failed to close server", zap.Error(err))
		return err
	}

	s.Logger.Info("http server close gracefully")
	return nil
}
