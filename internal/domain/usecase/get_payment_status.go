package usecase

import (
	"log/slog"
	"payments/internal/domain/repository"
)

type GetPaymentStatusUseCase struct {
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

func NewGetPaymentStatusUseCase() GetPaymentStatusUseCase {
	return GetPaymentStatusUseCase{}
}

func (gpsuc GetPaymentStatusUseCase) Execute(
	input GetPaymentStatusInput) GetPaymentStatusOutput {
	paymentEntity, exists, err := gpsuc.paymentRepository.Get(input.Identifier)
	if exists {
		return GetPaymentStatusOutput{Error: "identifier já existe na base"}
	} else if err != nil {
		slog.Error("erro ao buscar repository",
			"error", err)

		return GetPaymentStatusOutput{Error: "erro ao criar request"}
	}

	return GetPaymentStatusOutput{ID: paymentEntity.ID, Status: paymentEntity.State}
}
