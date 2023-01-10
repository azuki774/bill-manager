package server

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type APIService interface{}
type Server struct {
	Port   string
	Logger *zap.Logger
	APISvc APIService
}

func (s *Server) Start(ctx context.Context) (err error) {
	s.Logger.Info("server start")
	addr := fmt.Sprintf(":%s", s.Port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}

	return nil
}
