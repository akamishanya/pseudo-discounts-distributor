package dbrepository

import (
	"database/sql"
	"pseudo-discounts-distributor/internal/db/dbentity"
)

type ProductRepository interface {
	GetAllProducts() []dbentity.ProductEntity

	openConnection() *sql.DB
	closeConnection(db *sql.DB)
}
