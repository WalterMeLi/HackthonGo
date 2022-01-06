package internal

import (
	"context"

	"github.com/WalterMeLi/HackthonGo/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Sale, error)
	Insert(ctx context.Context, s domain.Sale) (int, error)
	Update(ctx context.Context, s domain.Sale) (domain.Sale, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
func (s *service) GetAll(ctx context.Context) ([]domain.Sale, error) {
	sales, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (s *service) Insert(ctx context.Context, sale domain.Sale) (int, error) {
	id, err := s.repository.Insert(ctx, sale)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (serv *service) Update(ctx context.Context, s domain.Sale) (domain.Sale, error) {
	sale, err := serv.repository.Get(ctx, s.ID)
	if err != nil {
		return domain.Sale{}, err
	}

	if s.IdInvoice == 0 {
		s.IdInvoice = sale.IdInvoice
	}

	if s.IdProduct == 0 {
		s.IdProduct = sale.IdProduct
	}

	if s.Quantity == 0 {
		s.Quantity = sale.Quantity
	}

	error := serv.repository.Update(ctx, s)

	return s, error
}
