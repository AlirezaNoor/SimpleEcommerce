package product

import (
	"database/sql"
)

type Repository interface {
	FetchAll() ([]Product, error)
	FetchByID(id int64) (*Product, error)
	Store(p *Product) error
	Update(p *Product) error
	Delete(id int64) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FetchAll() ([]Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *repository) FetchByID(id int64) (*Product, error) {
	var p Product
	err := r.db.QueryRow("SELECT id, name, description, price, stock FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) Store(p *Product) error {
	return r.db.QueryRow(
		"INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Name, p.Description, p.Price, p.Stock,
	).Scan(&p.ID)
}

func (r *repository) Update(p *Product) error {
	_, err := r.db.Exec(
		"UPDATE products SET name = $1, description = $2, price = $3, stock = $4 WHERE id = $5",
		p.Name, p.Description, p.Price, p.Stock, p.ID,
	)
	return err
}

func (r *repository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
