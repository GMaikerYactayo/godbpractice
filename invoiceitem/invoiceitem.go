package invoiceitem

import "time"

type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreateAt        time.Time
	UpdateAt        time.Time
}

type Models []*Model

type Storage interface {
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}
