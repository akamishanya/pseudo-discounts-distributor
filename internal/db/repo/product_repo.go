package repo

import "pseudo-discounts-distributor/internal/db/entity"

type ProductRepo interface {
	GetAll() (products []entity.Product, err error)
	UpdateById(product entity.Product) (rowsAffected int64, err error)
}
