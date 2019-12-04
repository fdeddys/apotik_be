package dbmodels

import "time"

// SalesOrder ...
type Receive struct {
	ID          int64     `json:"id" gorm:"column:id"`
	ReceiveNo   string    `json:"salesOrderNo" gorm:"column:sales_order_no"`
	ReceiveDate time.Time `json:"orderDate" gorm:"column:order_date"`
}

// TableName ...
func (o *Receive) TableName() string {
	return "public.receive"
}
