package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgresql://"+
			"postgres:password@localhost:5432/db_practice?sslmode=disabled")
		if err != nil {
			log.Fatalf("Can´t open db: %v", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatalf("Can´t do ping: %v", err)
		}
		fmt.Println("Connected to postgres ")
	})
}

func Pool() *sql.DB {
	return db
}
