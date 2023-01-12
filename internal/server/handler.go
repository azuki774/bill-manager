package server

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func (s *Server) middlewareLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			s.Logger.Info("access", zap.String("url", r.URL.Path), zap.String("X-Forwarded-For", r.Header.Get("X-Forwarded-For")))
		}
		h.ServeHTTP(w, r)
	})
}

func (s *Server) getBillyyyymm(w http.ResponseWriter, r *http.Request) {
	yyyymm := chi.URLParam(r, "yyyymm")
	ctx := context.Background()
	bills, err := s.APISvc.GetBills(ctx, yyyymm)
	if err != nil {
		if errors.Is(err, model.ErrInvalidData) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: %s\n", err.Error())
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s\n", err.Error())
			return
		}
	}

	outputJson, err := json.Marshal(&bills)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(outputJson))
}
