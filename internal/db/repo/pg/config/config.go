package config

import (
	"fmt"
	"os"
)

func GetConnectionString() (connectionString string, err error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if user == "" || password == "" || dbname == "" || sslmode == "" {
		err = fmt.Errorf(
			"environment variables not set (%s: %s, %s: %s, %s: %s, %s: %s)",
			"DB_USER", user,
			"DB_PASSWORD", password,
			"DB_DBNAME", dbname,
			"DB_SSLMODE", sslmode,
		)
		return
	}

	connectionString = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, sslmode)

	return
}
