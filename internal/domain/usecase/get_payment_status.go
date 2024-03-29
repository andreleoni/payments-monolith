package usecase

import (
	"log/slog"
	"payments/internal/domain/repository"
)

type GetPaymentStatusUseCase struct {
	logger            *slog.Logger
	paymentRepository repository.Payment
}

type GetPaymentStatusInput struct {
	Identifier string
}

type GetPaymentStatusOutput struct {
	ID     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

func NewGetPaymentStatusUseCase(
	logger *slog.Logger, paymentRepository repository.Payment) GetPaymentStatusUseCase {

	return GetPaymentStatusUseCase{logger: logger, paymentRepository: paymentRepository}
}

func (gpsuc GetPaymentStatusUseCase) Execute(
	input GetPaymentStatusInput) GetPaymentStatusOutput {
	paymentEntity, exists, err := gpsuc.paymentRepository.Get(input.Identifier)
	if exists {
		return GetPaymentStatusOutput{Status: paymentEntity.State}
	} else if err != nil {
		slog.Error("erro ao buscar repository",
			"error", err)

		return GetPaymentStatusOutput{Error: paymentEntity.Error}
	}

	return GetPaymentStatusOutput{ID: paymentEntity.ID, Status: paymentEntity.State}
}
