package models

import (
	"context"
	"html/template"
	"log/slog"
	"sync"
)

type contextKey string

const (
	LoggerKey contextKey = "logger"
)

func NewTemplate() *template.Template {
	var (
		init  sync.Once
		templ *template.Template
	)

	init.Do(func() {
		templ = template.Must(template.ParseGlob("templates/*.html"))
	})

	return templ
}

func GetLoggerFromCtx(ctx context.Context) *slog.Logger {
	val := ctx.Value(LoggerKey)

	if val == nil {
		slog.LogAttrs(ctx, slog.LevelError, "nil logger found in ctx")
		return slog.Default()
	}

	logger, ok := val.(*slog.Logger)
	if !ok {
		slog.LogAttrs(ctx, slog.LevelDebug, "not a type of expected logger")
		return slog.Default()
	}

	return logger
}
