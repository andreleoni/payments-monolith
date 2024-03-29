package main

import (
	"log/slog"
	"os"

	"payments/internal/interfaces/adapters/controllers"
	"payments/internal/interfaces/database/mongodb"
	"payments/internal/interfaces/gateways/persistence"

	"payments/pkg/middleware"
	"payments/pkg/queue"

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

	queueService := queue.NewQueue()
	go queueService.Consumer(logger)

	mongodb.MongoDBSetup()

	r := gin.New()

	r.Use(middleware.DefaultStructuredLogger())
	r.Use(gin.Recovery())

	paymentRepository := persistence.NewPaymentRepository(mongodb.MongoDB)

	v1 := r.Group("/api/v1")

	paymentsController := controllers.NewPaymentsController(
		logger, paymentRepository, queueService)

	v1.GET("/payments/:identifier", paymentsController.Get)

	v1.POST("/payments", paymentsController.Create)

	r.Run(":9090")
}
