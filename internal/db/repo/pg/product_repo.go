package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"pseudo-discounts-distributor/internal/db/entity"
	"pseudo-discounts-distributor/internal/db/repo/pg/config"
)

type ProductRepo struct {
	isInitialized bool
	db            *sql.DB
}

func (productRepo *ProductRepo) GetAll() (products []entity.Product, err error) {
	if !productRepo.isInitialized {
		err = fmt.Errorf("the repository is not initialized")
		return
	}

	rows, err := productRepo.db.Query("SELECT * FROM products")
	if err != nil {
		return
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			if err != nil {
				err = fmt.Errorf("%s; %s", err, closeErr)
			} else {
				err = closeErr
			}
		}
	}()

	for rows.Next() {
		product := entity.Product{}

		scanErr := rows.Scan(
			&product.Id,
			&product.MarketplaceId,
			&product.Name,
			&product.IsAdultOnly,
			&product.Link,
			&product.ImageId,
		)

		if scanErr != nil {
			if err != nil {
				err = fmt.Errorf("%s; %s", err, scanErr)
			} else {
				err = scanErr
			}
			continue
		}

		products = append(products, product)
	}

	return
}

func NewProductRepo() (productRepo *ProductRepo, err error) {
	connectionString, err := config.GetConnectionString()
	if err != nil {
		return
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return
	}

	productRepo = &ProductRepo{
		isInitialized: true,
		db:            db,
	}

	return
}

func DisposeProductRepo(productRepo *ProductRepo) (err error) {
	if productRepo == nil || !productRepo.isInitialized {
		err = fmt.Errorf("the repository is not initialized")
		return
	}

	productRepo.isInitialized = false

	if productRepo.db != nil {
		err = productRepo.db.Close()
	}

	return
}