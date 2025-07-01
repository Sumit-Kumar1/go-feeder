package handler

import (
	"feeder/internal/models"
	"html/template"
	"log/slog"
	"net/http"
)

type Handler struct {
	templ *template.Template
}

func New() *Handler {
	return &Handler{
		templ: models.NewTemplate(),
	}
}

func (h *Handler) Root(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := models.GetLoggerFromCtx(ctx)

	if err := h.templ.ExecuteTemplate(w, "index", nil); err != nil {
		logger.LogAttrs(ctx, slog.LevelError, "error while executing template: index", slog.String("error", err.Error()))

		http.Error(w, "error while executing template: index", http.StatusInternalServerError)

		return
	}
}
