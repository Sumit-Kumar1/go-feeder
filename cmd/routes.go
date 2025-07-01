package cmd

import (
	"context"
	"feeder/internal/handler"
	"feeder/internal/models"
	"log/slog"
	"net/http"
)

func addRoutes(mux *http.ServeMux, logger *slog.Logger) {
	h := handler.New()
	mdw := newMiddleware(logger)

	mux.HandleFunc("GET /", mdw(h.Root))
}

type middleware func(h http.HandlerFunc) http.HandlerFunc

func newMiddleware(logger *slog.Logger) middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newCtx := context.WithValue(r.Context(), models.MiddlwareLogger, logger)

			h.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
