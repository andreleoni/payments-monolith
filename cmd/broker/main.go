package main

import (
	"log/slog"
	"os"

	controllersbroker "payments/internal/interfaces/adapters/controllers/broker"

	"payments/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	desiredLogLevel := os.Getenv("LOG_LEVEL")

	logLevel := slog.LevelInfo

	if desiredLogLevel == "DEBUG" {
		logLevel = slog.LevelDebug
	} else if desiredLogLevel == "WARN" {
		logLevel = slog.LevelWarn
	} else if desiredLogLevel == "ERROR" {
		logLevel = slog.LevelError
	}

	opts := &slog.HandlerOptions{Level: logLevel}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)

	slog.SetDefault(logger)

	r := gin.New()

	r.Use(middleware.DefaultStructuredLogger())
	r.Use(gin.Recovery())

	// This endpoint represents an enqueue endpoint for async messages processing
	v1 := r.Group("/api/v1")

	brokerController := controllersbroker.NewBrokerController(*logger)

	v1.POST("/enqueue", brokerController.Enqueue)

	go brokerController.Consumer()

	r.Run(":8081")
}
