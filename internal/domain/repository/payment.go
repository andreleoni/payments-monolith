package repository

import "payments/internal/domain/entity"

type Payment interface {
	Get(identifier string) (*entity.Payment, bool, error)
	Create(*entity.Payment) error
}
