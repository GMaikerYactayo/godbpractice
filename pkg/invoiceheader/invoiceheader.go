package invoiceheader

import (
	"database/sql"
	"time"
)

type Model struct {
	ID       uint
	Client   string
	CreateAt time.Time
	UpdateAt time.Time
}

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, *Model) error
}

// Service of invoiceHeader
type Service struct {
	storage Storage
}

// NewService return a pinter of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate invoiceHeader
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

//Create is used for create invoiceHeader
//func (s *Service) Create(m *Model) error {
//	m.CreateAt = time.Now()
//	return s.storage.Create(m)
//}
