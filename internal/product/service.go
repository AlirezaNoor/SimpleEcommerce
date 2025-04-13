package product

import "errors"

type Service interface {
	GetAll() ([]Product, error)
	GetByID(id int64) (*Product, error)
	Create(p *Product) error
	Update(p *Product) error
	Delete(id int64) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]Product, error) {
	return s.repo.FetchAll()
}

func (s *service) GetByID(id int64) (*Product, error) {
	return s.repo.FetchByID(id)
}

func (s *service) Create(p *Product) error {
	if p.Name == "" {
		return errors.New("نام محصول نمی‌تواند خالی باشد")
	}
	return s.repo.Store(p)
}

func (s *service) Update(p *Product) error {
	return s.repo.Update(p)
}

func (s *service) Delete(id int64) error {
	return s.repo.Delete(id)
}
