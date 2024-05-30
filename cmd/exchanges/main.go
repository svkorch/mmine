package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/svkorch/mmine/internal/api"
	"github.com/svkorch/mmine/internal/config"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	_ = godotenv.Load()

	cfg := config.GetConfigInstance()
	config.LogConfig()

	logrus.Info("`Exchanges' service is starting...")

	mux := http.NewServeMux()
	mux.HandleFunc("/exchange", api.Exchange)

	srv := http.Server{
		Addr:    cfg.HTTPServerAddr,
		Handler: mux,
	}

	go func() {
		ctx := context.Background()
		ctx, _ = signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)

		select {
		case <-ctx.Done():
		}

		logrus.Info("stopping http server...")
		srv.Shutdown(context.Background())
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
		logrus.Warn("http server has not stopped gracefully")
	} else {
		logrus.Info("http server has stopped gracefully")
	}
}
