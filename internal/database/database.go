package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"os"
)

var DB *sqlx.DB

func EstablishConnection() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}
}
