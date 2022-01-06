package product

import (
	"context"
	"errors"

	"github.com/WalterMeLi/HackthonGo/internal/domain"
)

var (
	ErrNotFound = errors.New("product not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Insert(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
	LoadData(ctx context.Context) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

//Calls the GetAll function from the repository to show all the registry
func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	ps, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

//Calls the Get function from the repository to obtain one item from the registry
func (s *service) Get(ctx context.Context, id int) (domain.Product, error) {
	ps, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return ps, nil
}

//Calls the Save function from the Repository to store the data
func (s *service) Insert(ctx context.Context, p domain.Product) (int, error) {
	product, err := s.repository.Insert(ctx, p)
	if err != nil {
		return 0, err
	}
	return product, nil
}

//Calls the Update function from the repository to update a registry item
func (s *service) Update(ctx context.Context, p domain.Product) error {

	_, err := s.repository.Get(ctx, p.ID)

	if err != nil {
		return ErrNotFound
	}

	return s.repository.Update(ctx, p)
}

func (s *service) LoadData(ctx context.Context) error {
	err := s.repository.LoadData(ctx)
	if err != nil {
		return err
	}
	return nil
}
