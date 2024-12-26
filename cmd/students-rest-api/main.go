package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Kishan-Thanki/go-students-rest-api/internals/config"
	"github.com/Kishan-Thanki/go-students-rest-api/internals/http/handlers/student"
	"github.com/Kishan-Thanki/go-students-rest-api/internals/storage/sqlite"
)

func main() {
	// Load Config
	cfg := config.MustLoad()

	// Database Setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// Router Setup
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.Create(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students/", student.GetList(storage))

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
