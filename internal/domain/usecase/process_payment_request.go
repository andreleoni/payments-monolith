package usecase

import "payments/internal/interfaces/adapters/dto"

type ProcessPaymentRequestUseCase struct {
}

type ProcessPaymentRequestInput struct {
	PaymentRequest dto.PaymentRequest
}

type ProcessPaymentRequestOutput struct {
}

func NewProcessPaymentRequestUseCase() ProcessPaymentRequestUseCase {
	return ProcessPaymentRequestUseCase{}
}
