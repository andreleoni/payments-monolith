package usecase

import (
	"log/slog"
	"payments/internal/domain/repository"
)

type ProcessPaymentRequestUseCase struct {
	paymentRepository repository.Payment
}

type ProcessPaymentRequestInput struct {
	PaymentID string
}

type ProcessPaymentRequestOutput struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewProcessPaymentRequestUseCase(paymentRepository repository.Payment) ProcessPaymentRequestUseCase {
	return ProcessPaymentRequestUseCase{paymentRepository: paymentRepository}
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

	return ProcessPaymentRequestOutput{ID: paymentEntity.ID}
}
