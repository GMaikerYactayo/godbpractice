package invoiceitem

import "time"

type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreateAt        time.Time
	UpdateAt        time.Time
}

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
}

// Service of invoiceItem
type Service struct {
	storage Storage
}

// NewService return a pinter of service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate invoiceItem
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
