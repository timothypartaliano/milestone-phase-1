package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase(dburl string) *sql.DB{
	db, err := sql.Open("mysql", dburl)
	if err != nil {
		log.Fatal("Failed to initialize connection to database")
	}

	return db
}