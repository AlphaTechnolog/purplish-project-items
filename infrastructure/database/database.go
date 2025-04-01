package database

import (
	"database/sql"
	"log"
)

func MustOpenDB(driverName, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
		panic(err)
	}

	return db
}
