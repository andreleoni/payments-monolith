package usecase

import (
	"log/slog"
	"payments/internal/domain/entity"
	"payments/internal/domain/repository"
	"payments/internal/interfaces/adapters/dto"
	"payments/pkg/queue"
)

type CreatePaymentRequestUseCase struct {
	paymentRepository repository.Payment
	queueService      *queue.Queue
}

type CreatePaymentRequestInput struct {
	PaymentRequest dto.PaymentRequest
}

type CreatePaymentRequestOutput struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewCreatePaymentRequestUseCase(
	paymentRepository repository.Payment, queueService *queue.Queue) CreatePaymentRequestUseCase {

	return CreatePaymentRequestUseCase{paymentRepository: paymentRepository, queueService: queueService}
}

func (cpruc CreatePaymentRequestUseCase) Execute(
	logger *slog.Logger, cpri CreatePaymentRequestInput) CreatePaymentRequestOutput {

	_, exists, err := cpruc.paymentRepository.Get(cpri.PaymentRequest.Identifier)
	if exists {
		return CreatePaymentRequestOutput{Error: "identifier já existe na base"}
	} else if err != nil {
		slog.Error("erro ao buscar repository",
			"error", err)

		return CreatePaymentRequestOutput{Error: "erro ao criar request"}
	}

	paymentEntity := cpruc.dtoToEntityConverter(cpri.PaymentRequest)

	err = cpruc.paymentRepository.Create(&paymentEntity)
	if err != nil {
		slog.Error("erro ao buscar payment",
			"error", err)

		return CreatePaymentRequestOutput{Error: "erro ao criar request"}
	}

	return CreatePaymentRequestOutput{ID: paymentEntity.ID}
}

func (CreatePaymentRequestUseCase) dtoToEntityConverter(
	pr dto.PaymentRequest) entity.Payment {

	return entity.Payment{
		Identifier:        pr.Identifier,
		ValueCents:        pr.ValueCents,
		UserFullName:      pr.User.FullName,
		UserCPF:           pr.User.CPF,
		CreditCardNumber:  pr.CreditCard.Number,
		CreditCardCVV:     pr.CreditCard.CVV,
		CreditCardExpires: pr.CreditCard.Expires,
		AddressStreet:     pr.User.Address.Street,
		AddressZipcode:    pr.User.Address.Zipcode,
		AddressNumber:     pr.User.Address.Number,
		AddressComplement: pr.User.Address.Complement,
	}
}
