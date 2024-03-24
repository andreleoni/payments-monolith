package usecase

import "payments/internal/interfaces/adapters/dto"

type CreatePaymentRequestUseCase struct {
}

type CreatePaymentRequestInput struct {
	PaymentRequest dto.PaymentRequest
}

type CreatePaymentRequestOutput struct {
}

func NewCreatePaymentRequestUseCase() CreatePaymentRequestUseCase {
	return CreatePaymentRequestUseCase{}
}
