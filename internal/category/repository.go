package category

import "database/sql"

type Repository interface {
	GetAll() ([]Category, error)
	GetByID(Id int32) (*Category, error)
	Create(c *Category) error
	Update(c *Category) error
	Delete(id int32) error
}

type repository struct {
	db *sql.DB
}

// Create implements Repository.
func (r *repository) Create(c *Category) error {
	panic("unimplemented")
}

// Delete implements Repository.
func (r *repository) Delete(id int32) error {
	panic("unimplemented")
}

// GetAll implements Repository.
func (r *repository) GetAll() ([]Category, error) {
	panic("unimplemented")
}

// GetByID implements Repository.
func (r *repository) GetByID(Id int32) (*Category, error) {
	panic("unimplemented")
}

// Update implements Repository.
func (r *repository) Update(c *Category) error {
	panic("unimplemented")
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
