package usecase

import (
	"log/slog"
	"payments/internal/domain/repository"
	"payments/pkg/externalpaymentservice"
)

type ProcessPaymentRequestUseCase struct {
	logger            *slog.Logger
	paymentRepository repository.Payment
}

type ProcessPaymentRequestInput struct {
	PaymentID string
}

type ProcessPaymentRequestOutput struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewProcessPaymentRequestUseCase(
	logger *slog.Logger, paymentRepository repository.Payment) ProcessPaymentRequestUseCase {

	return ProcessPaymentRequestUseCase{logger: logger, paymentRepository: paymentRepository}
}

func (ppruc ProcessPaymentRequestUseCase) Execute(
	input ProcessPaymentRequestInput) ProcessPaymentRequestOutput {
	paymentEntity, exists, err := ppruc.paymentRepository.Get(input.PaymentID)
	if exists {
		return ProcessPaymentRequestOutput{Error: "identifier já existe na base"}
	} else if err != nil {
		slog.Error("erro ao buscar repository",
			"error", err)

		return ProcessPaymentRequestOutput{Error: "erro ao criar request"}
	}

	externalServiceIdentifier, err := externalpaymentservice.Pay(input.PaymentID)

	if err != nil {
		ppruc.paymentRepository.SetError(input.PaymentID, err)
	} else {
		ppruc.paymentRepository.SetApproved(input.PaymentID, externalServiceIdentifier, "approved")
	}

	ppruc.logger.Info(
		"Payment processed with success!",
		"payment_id", paymentEntity)

	return ProcessPaymentRequestOutput{ID: paymentEntity.ID}
}
