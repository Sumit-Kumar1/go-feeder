package cmd

import (
	"context"
	"go-feeder/internal/handler"
	"go-feeder/internal/models"
	"go-feeder/internal/service"
	"go-feeder/internal/store"
	"log/slog"
	"net/http"
)

func addRoutes(mux *http.ServeMux, logger *slog.Logger) {
	s := store.New()
	svc := service.New(s)
	h := handler.New(svc)

	mdw := newMiddleware(logger)

	mux.HandleFunc("GET /", mdw(h.Root))
	mux.HandleFunc("POST /feeds", mdw(h.Feeds))
}

type middleware func(h http.HandlerFunc) http.HandlerFunc

func newMiddleware(logger *slog.Logger) middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newCtx := context.WithValue(r.Context(), models.LoggerKey, logger)

			h.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
