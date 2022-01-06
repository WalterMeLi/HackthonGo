package internal

import (
	"context"
	"database/sql"

	"github.com/WalterMeLi/HackthonGo/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Sale, error)
	Insert(ctx context.Context, s domain.Sale) (int, error)
	Update(ctx context.Context, s domain.Sale) error
	Get(ctx context.Context, id int) (domain.Sale, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Sale, error) {
	query := "SELECT * FROM sales"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var sales []domain.Sale

	for rows.Next() {
		s := domain.Sale{}
		_ = rows.Scan(&s.ID, &s.IdInvoice, &s.IdProduct, &s.Quantity)
		sales = append(sales, s)
	}
	return sales, nil
}

func (r *repository) Insert(ctx context.Context, s domain.Sale) (int, error) {
	query := "INSERT INTO sales (idinvoice, idproduct, quantity) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(s.IdInvoice, s.IdProduct, s.Quantity)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Sale) error {
	query := "UPDATE sales SET idinvoice = ?, idproduct = ?, quantity = ? WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(s.IdInvoice, s.IdProduct, s.Quantity, s.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	query := "SELECT id FROM sales WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&id)
	return err == nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Sale, error) {
	query := "SELECT * FROM sales WHERE id=?;"
	row := r.db.QueryRow(query, id)
	s := domain.Sale{}
	err := row.Scan(&s.ID, &s.IdInvoice, &s.IdProduct, &s.Quantity)
	if err != nil {
		return domain.Sale{}, err
	}

	return s, nil
}
