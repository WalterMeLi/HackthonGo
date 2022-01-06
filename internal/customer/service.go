package internal

import (
	"context"

	"github.com/WalterMeLi/HackthonGo/internal/domain"
)

type Service interface {
	Insert(ctx context.Context, c domain.Customer) (int, error)
	Update(ctx context.Context, c domain.Customer) error
	Get(ctx context.Context, id int) (domain.Customer, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (serv *service) Insert(ctx context.Context, p domain.Customer) (domain.Customer, error) {

	id, err := serv.repository.Insert(ctx, p)

	if err == nil {
		p.ID = id
		return p, nil
	}

	return domain.Customer{}, err
}

func (serv *service) Update(ctx context.Context, c domain.Customer) error {

	customer, err := serv.repository.Get(ctx, c.ID)

	if err != nil {
		return domain.Customer{}, err
	}
	if c.LastName == "" {
		c.LastName = customer.LastName
	}
	if c.FirstName == "" {
		c.FirstName = customer.FirstName
	}
	if c.Condition == "" {
		c.Condition = customer.Condition
	}

	errUpdate := serv.repository.Update(ctx, c)
	return c, errUpdate
}

func (serv *service) Get(ctx context.Context, id int) (domain.Customer, error) {
	customer, err := serv.repository.Get(ctx, id)

	if err != nil {
		return domain.Customer{}, err
	}

	return customer, nil
}
