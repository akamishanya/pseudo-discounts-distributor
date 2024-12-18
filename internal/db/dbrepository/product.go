package dbrepository

import (
	"database/sql"
	"pseudo-discounts-distributor/internal/db/dbentity"
)

type ProductRepository interface {
	openConnection()
	closeConnection(db *sql.DB)

	GetAllProducts() []dbentity.ProductEntity
}
