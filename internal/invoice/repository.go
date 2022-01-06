package internal

import (
	"context"
	"database/sql"

	"github.com/WalterMeLi/HackthonGo/internal/domain"
)

type Repository interface {
	Insert(ctx context.Context, i domain.Invoice) (int, error)
	Update(ctx context.Context, i domain.Invoice) error
	Get(ctx context.Context, id int) (domain.Invoice, error)
	Exists(ctx context.Context, idcustomer int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id int) (domain.Invoice, error) {
	query := "SELECT * FROM invoices WHERE id=?;"
	row := r.db.QueryRow(query, id)
	i := domain.Invoice{}
	err := row.Scan(&i.ID, &i.DateTime, &i.IdCustomer, &i.Total)
	if err != nil {
		return domain.Invoice{}, err
	}

	return i, nil
}

func (r *repository) Exists(ctx context.Context, idcustomer int) bool {
	query := "SELECT idcustomer FROM invoices WHERE idcustomer=?;"
	row := r.db.QueryRow(query, idcustomer)
	err := row.Scan(&idcustomer)
	return err == nil
}

func (r *repository) Insert(ctx context.Context, i domain.Invoice) (int, error) {
	query := "INSERT INTO invoices (datetime, idcustomer, total) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(i.DateTime, i.IdCustomer, i.Total)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, i domain.Invoice) error {
	query := "UPDATE invoices SET datetime=?, idcustomer=?, total=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(i.DateTime, i.IdCustomer, i.Total, i.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
