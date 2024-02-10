package invoiceheader

import "time"

type Model struct {
	ID       uint
	Client   string
	CreateAt time.Time
	UpdateAt time.Time
}
