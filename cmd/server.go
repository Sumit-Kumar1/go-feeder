package cmd

import (
	"context"
	"errors"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

const (
	envPath = "configs/.env"
)

func Run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := setupLogger(w)
	mux := http.NewServeMux()

	addRoutes(mux, logger)

	if err := godotenv.Load(envPath); err != nil {
		logger.LogAttrs(ctx, slog.LevelError, "error while loading env configs", slog.String("configs", err.Error()))
		return err
	}

	port := getEnvOrDefault("HTTP_PORT", "9001")
	logger.LogAttrs(ctx, slog.LevelInfo, "started server at address: "+":"+port)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.LogAttrs(ctx, slog.LevelError, "listen and serve returned err:"+err.Error())
		}
	}()

	<-ctx.Done()
	logger.LogAttrs(ctx, slog.LevelInfo, "got interruption signal")

	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Printf("server shutdown returned an err: %v\n", err)
		return err
	}

	log.Println("final")
	return nil
}

func getEnvOrDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	return val
}

func setupLogger(w io.Writer) *slog.Logger {
	var logLevel slog.Leveler

	switch getEnvOrDefault("LOG_LEVEL", "INFO") {
	case "DEBUG":
		logLevel = slog.LevelDebug
	case "WARN":
		logLevel = slog.LevelWarn
	case "ERROR":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		AddSource: false,
		Level:     logLevel,
	}))

	slog.SetDefault(logger)

	return logger
}
