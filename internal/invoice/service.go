package internal

import (
	"context"
	"errors"

	"github.com/WalterMeLi/HackthonGo/internal/domain"
)

// Errors
var (
	ErrNotFound  = errors.New("invoice not found")
	ErrSellerCID = errors.New("invoice idCustomer not exists")
	ErrNotCID    = errors.New("Invoice idCustomer is requered")
)

type Service interface {
	Insert(ctx context.Context, i domain.Invoice) (domain.Invoice, error)
	Update(ctx context.Context, i domain.Invoice, id int) error
	Get(ctx context.Context, id int) (domain.Invoice, error)
	// Exists(ctx context.Context, idcustomer int) bool
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Get(ctx context.Context, id int) (domain.Invoice, error) {
	invoice, error := s.repository.Get(ctx, id)

	if error != nil {
		return domain.Invoice{}, error
	}

	return invoice, error
}

func (s *service) Insert(ctx context.Context, invoice domain.Invoice) (domain.Invoice, error) {

	if s.repository.Exists(ctx, invoice.IdCustomer) {
		return domain.Invoice{}, ErrSellerCID
	}

	if invoice.IdCustomer == 0 {
		return domain.Invoice{}, ErrNotCID
	}

	newId, error := s.repository.Insert(ctx, invoice)

	if error != nil {
		return domain.Invoice{}, error
	}

	invoice.ID = newId

	return invoice, nil
}

func (s *service) Update(ctx context.Context, invoice domain.Invoice, id int) (domain.Invoice, error) {
	invo, error := s.repository.Get(ctx, id)

	if error != nil {
		return domain.Invoice{}, error
	}

	if s.repository.Exists(ctx, invoice.IdCustomer) {
		return domain.Invoice{}, ErrSellerCID
	}

	invoice.ID = invo.ID

	if invoice.IdCustomer != 0 {
		invo.IdCustomer = invoice.IdCustomer
	}

	if invoice.DateTime != "" {
		invo.DateTime = invoice.DateTime
	}

	if invoice.Total != 0 {
		invo.Total = invoice.Total
	}

	error = s.repository.Update(ctx, invoice)

	if error != nil {
		return domain.Invoice{}, error
	}

	return invoice, nil
}
