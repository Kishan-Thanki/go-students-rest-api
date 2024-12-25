package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Kishan-Thanki/go-students-rest-api/internals/config"
)

func main() {
	// Load Config
	cfg := config.MustLoad()

	// Database Setup

	// Router Setup
	router := http.NewServeMux()
	router.HandleFunc("GET /api/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Students API"))
	})

	// Server Setup
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("Address:", cfg.Addr))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	signal := <-done
	slog.Info("Received signal, shutting down", slog.String("signal", signal.String()))

	slog.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error during server shutdown", slog.String("error", err.Error()))
	} else {
		slog.Info("Server shut down gracefully")
	}
}
