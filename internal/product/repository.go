package product

import (
	"context"
	"database/sql"

	"github.com/WalterMeLi/HackthonGo/internal/domain"
)

const (
	GetAllProduct   = "SELECT * FROM products WHERE id=?;"
	GetProduct      = "SELECT * FROM products WHERE id=?;"
	StoreProduct    = "INSERT INTO products (id, description, price) VALUES (?, ?, ?);"
	UpdateProduct   = "UPDATE products SET description=?, price=?  WHERE id=?"
	LoadDataProduct = "LOAD DATA LOCAL INFILE '/Users/wcastillo/Documents/Bootcamp/HackthonGo/HackthonGo/datos/products.txt'  INTO TABLE products FIELDS TERMINATED BY ';' LINES TERMINATED BY '\n' "
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Insert(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
	LoadData(ctx context.Context) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

//Gets all the items in the registry
func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	query := GetAllProduct
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		p := domain.Product{}
		_ = rows.Scan(&p.ID, &p.Description, &p.Price)
		products = append(products, p)
	}

	return products, nil
}

//Get one item from the registry
func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	query := GetProduct
	row := r.db.QueryRow(query, id)
	p := domain.Product{}
	err := row.Scan(&p.ID, &p.Description, &p.Price)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

//Saves the data in the registry
func (r *repository) Insert(ctx context.Context, p domain.Product) (int, error) {
	query := StoreProduct
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&p.Description, p.Price)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

//Update an item from the registry
func (r *repository) Update(ctx context.Context, p domain.Product) error {
	query := UpdateProduct
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.Description, p.Price, p.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) LoadData(ctx context.Context) error {
	query := LoadDataProduct
	_, err := r.db.Query(query)
	if err != nil {
		return err
	}
	return nil
}
