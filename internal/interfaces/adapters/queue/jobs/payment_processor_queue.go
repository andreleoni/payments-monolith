package jobs

import (
	"log/slog"
	"payments/internal/domain/usecase"
)

type PaymentProcessorJob struct {
	logger *slog.Logger
	id     string
}

func NewPaymentProcessorJob(
	logger *slog.Logger, id string) PaymentProcessorJob {
	return PaymentProcessorJob{logger: logger, id: id}
}

func (j PaymentProcessorJob) Logger() *slog.Logger {
	return j.logger
}

func (j PaymentProcessorJob) Process() error {
	slog.Debug("Processing job", "id", j.id)

	ppri := usecase.ProcessPaymentRequestInput{PaymentID: j.id}

	usecase.

	return nil
}
