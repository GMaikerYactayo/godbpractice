package storage

import (
	"database/sql"
	"fmt"
	"github.com/GMaikerYactayo/godbpractice/pkg/product"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

// Driver of Storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

var (
	db   *sql.DB
	once sync.Once
)

func New(d Driver) {
	switch d {
	case MySQL:
		newMySqlDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgresql://"+
			"postgres:password@localhost:5432/db_practice?sslmode=disable")
		if err != nil {
			log.Fatalf("Can´t open db: %v", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatalf("Can´t do ping: %v", err)
		}
		fmt.Println("Connected to postgres")
	})
}

// NewMySqlDB connection to MySQL
func newMySqlDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/db_practice?charset=utf8&parseTime=true")
		if err != nil {
			log.Fatalf("Can´t open db: %v", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatalf("Can´t do ping: %v", err)
		}
		fmt.Println("Connected to mysql")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

// DAOProduct factory of product.Storage
func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQL:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("driver not implemented")
	}
}
