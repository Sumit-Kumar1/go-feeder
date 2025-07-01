package models

import (
	"context"
	"html/template"
	"log/slog"
	"sync"
)

type loggerKey string

const (
	MiddlwareLogger loggerKey = "logger"
)

func (l loggerKey) String() string {
	return "logger"
}

func GetLoggerFromCtx(ctx context.Context) *slog.Logger {
	val := ctx.Value(MiddlwareLogger)
	if val == nil {
		return slog.Default()
	}

	logger, ok := val.(*slog.Logger)
	if !ok {
		return slog.Default()
	}

	return logger
}

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
