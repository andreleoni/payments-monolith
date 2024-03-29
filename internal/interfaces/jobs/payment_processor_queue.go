package jobs

import (
	"fmt"
	"log/slog"
	"payments/internal/domain/repository"
	"payments/internal/domain/usecase"
)

type PaymentProcessorJob struct {
	logger            *slog.Logger
	id                string
	paymentRepository repository.Payment
}

func NewPaymentProcessorJob(
	logger *slog.Logger, id string, paymentRepository repository.Payment) PaymentProcessorJob {

	return PaymentProcessorJob{logger: logger, id: id, paymentRepository: paymentRepository}
}

func (j PaymentProcessorJob) Logger() *slog.Logger {
	return j.logger
}

func (j PaymentProcessorJob) Process() error {
	slog.Debug("Processing job", "id", j.id)

	ppruc := usecase.NewProcessPaymentRequestUseCase(j.logger, j.paymentRepository)

	ppri := usecase.ProcessPaymentRequestInput{PaymentID: j.id}

	output := ppruc.Execute(ppri)

	return fmt.Errorf(output.Error)
}
