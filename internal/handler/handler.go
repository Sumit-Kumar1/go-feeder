package handler

import (
	"context"
	"go-feeder/internal/models"
	"html/template"
	"log/slog"
	"net/http"
)

type Servicer interface {
	FetchFeeds(context.Context, string) (*models.RSS, error)
}

type Handler struct {
	templ   *template.Template
	Service Servicer
}

func New(svc Servicer) *Handler {
	return &Handler{
		templ:   models.NewTemplate(),
		Service: svc,
	}
}

func (h *Handler) Root(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := models.GetLoggerFromCtx(ctx)

	if err := h.templ.ExecuteTemplate(w, "index", nil); err != nil {
		logger.LogAttrs(ctx, slog.LevelError, "error while executing template: index",
			slog.String("error", err.Error()))

		http.Error(w, "error while executing template: index", http.StatusInternalServerError)

		return
	}
}

func (h *Handler) Feeds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := models.GetLoggerFromCtx(ctx)

	url := r.FormValue("url")

	resp, err := h.Service.FetchFeeds(ctx, url)
	if err != nil {
		logger.LogAttrs(ctx, slog.LevelError, "handler: error while fetching feed", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp == nil {
		logger.LogAttrs(ctx, slog.LevelError, "nil feed response")
	}

	logger.LogAttrs(ctx, slog.LevelInfo, "successfully fetched feeds",
		slog.String("feed-title", resp.Channel.Title),
		slog.Int("no. of items", len(resp.Channel.Items)))

	if err := h.templ.ExecuteTemplate(w, "feeds", resp); err != nil {
		logger.LogAttrs(ctx, slog.LevelError, "error while executing template: feeds",
			slog.String("error", err.Error()))

		http.Error(w, "error while executing template: feeds", http.StatusInternalServerError)
	}
}
