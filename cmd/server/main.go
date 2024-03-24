package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"payments/internal/interfaces/adapters/dto"

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

	v1 := r.Group("/api/v1")

	v1.POST("/payments", func(c *gin.Context) {
		logCorrelationID, _ := c.Get("logCorrelationID")

		contextlogger := logger.With("correlation_id", logCorrelationID)

		var paymentRequest dto.PaymentRequest

		if err := c.BindJSON(&paymentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Make a GET request
		response, err := http.Get("http://localhost:8081/enqueue")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer response.Body.Close()

		// Read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		contextlogger.Info("Enqueued Successfully", "transaction_id", string(body))

		c.JSON(http.StatusOK, string(body))
	})

	r.Run()
}
