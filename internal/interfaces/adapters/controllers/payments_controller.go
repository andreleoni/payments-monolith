package controllers

import (
	"fmt"
	"log/slog"
	"net/http"
	"payments/internal/domain/repository"
	"payments/internal/domain/usecase"
	"payments/pkg/queue"

	"payments/internal/interfaces/adapters/dto"

	"github.com/gin-gonic/gin"
)

type PaymentsController struct {
	logger            *slog.Logger
	paymentRepository repository.Payment
	queueService      *queue.Queue
}

func NewPaymentsController(
	logger *slog.Logger, paymentRepository repository.Payment, queueService *queue.Queue) PaymentsController {

	return PaymentsController{logger: logger, paymentRepository: paymentRepository, queueService: queueService}
}

func (ppc PaymentsController) Create(c *gin.Context) {
	createPaymentRequestUseCase := usecase.NewCreatePaymentRequestUseCase(
		ppc.paymentRepository, ppc.queueService)

	logCorrelationID, _ := c.Get("logCorrelationID")

	// Tips: With will clone the logger, so, you will not change the
	//   global slog state
	contextlogger := ppc.logger.With("correlation_id", logCorrelationID)

	var paymentRequest dto.PaymentRequest

	if err := c.BindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cpri := usecase.CreatePaymentRequestInput{PaymentRequest: paymentRequest}

	output := createPaymentRequestUseCase.Execute(ppc.logger, cpri)

	if output.Error != "" {
		fmt.Println(output)
		c.JSON(http.StatusUnprocessableEntity, output)

		return
	}

	contextlogger.Info("Created transaction ID sucessfully", "id", output.ID)

	c.JSON(http.StatusOK, output)
}

func (ppc PaymentsController) Get(c *gin.Context) {
	logCorrelationID, _ := c.Get("logCorrelationID")

	contextlogger := ppc.logger.With("correlation_id", logCorrelationID)

	contextlogger.Info("Get from database")

	data := map[string]string{
		"hello": "test",
	}

	c.JSON(http.StatusOK, data)
}
