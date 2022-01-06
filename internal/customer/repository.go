package internal

import (
	"context"
	"database/sql"

	"github.com/WalterMeLi/HackthonGo/internal/domain"
)

type Repository interface {
	Insert(ctx context.Context, c domain.Customer) (int, error)
	Update(ctx context.Context, c domain.Customer) error
	Get(ctx context.Context, id int) (domain.Customer, error)
	Exists(ctx context.Context, id int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(ctx context.Context, c domain.Customer) (int, error) {
	query := "INSERT INTO customers(last_name,first_name,condition) VALUES (?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&c.LastName, &c.FirstName, &c.Condition)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, c domain.Customer) error {
	query := "UPDATE buyers SET last_name=?, first_name=?, condition=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&c.ID, &c.LastName, &c.FirstName, &c.Condition)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Customer, error) {
	query := "SELECT * FROM customers WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	c := domain.Customer{}
	err := row.Scan(&c.ID, &c.LastName, &c.FirstName, &c.Condition)
	if err != nil {
		return domain.Customer{}, err
	}

	return c, nil
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	query := "SELECT id FROM customers WHERE id=?;"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&id)
	return err == nil
}
