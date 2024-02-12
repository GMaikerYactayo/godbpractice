package invoiceheader

import "time"

type Model struct {
	ID       uint
	Client   string
	CreateAt time.Time
	UpdateAt time.Time
}

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
}

// Service of invoiceHeader
type Service struct {
	storage Storage
}

// NewService return a pinter of service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate invoiceHeader
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
