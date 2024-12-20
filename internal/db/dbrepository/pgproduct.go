package dbrepository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"pseudo-discounts-distributor/internal/db/dbentity"
)

type PostgresProductRepository struct{}

func (p *PostgresProductRepository) GetAllProducts() []dbentity.ProductEntity {
	db := p.openConnection()
	defer p.closeConnection(db)

	query := "SELECT * FROM products"
	rows, err := db.Query(query)
	if err != nil {
		logrus.Fatalf("Failed to execute the query \"%s\": %s\n", query, err.Error())
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			logrus.Fatalln("Failed to close rows: ", err.Error())
		}
	}(rows)

	var products []dbentity.ProductEntity

	for rows != nil && rows.Next() {
		p := dbentity.ProductEntity{}

		err = rows.Scan(&p.Id, &p.MarketplaceId, &p.Name, &p.IsAdultOnly, &p.Link, &p.ImageId)
		if err != nil {
			logrus.Errorln("Failed to scan row: ", err.Error())
			continue
		}

		products = append(products, p)
	}

	return products
}

func (p *PostgresProductRepository) openConnection() *sql.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")
	sslmode := os.Getenv("DB_SSLMODE")
	if user == "" || password == "" || dbname == "" || sslmode == "" {
		logrus.Fatalf(
			"Invalid database connection parameters: user: \"%s\", password: \"%s\", dbname: \"%s\", sslmode: \"%s\"\n",
			user, password, dbname, sslmode,
		)
	}

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, sslmode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logrus.Fatalln("Failed to connect to the database: ", err.Error())
	}

	return db
}

func (p *PostgresProductRepository) closeConnection(db *sql.DB) {
	if err := db.Close(); err != nil {
		logrus.Fatalln("Failed to close the database connection: ", err.Error())
	}
}
