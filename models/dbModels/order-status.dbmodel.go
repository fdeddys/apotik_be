package dbmodels

import "time"

// OrderStatus ...
type OrderStatus struct {
	ID           int64     `json:"id" gorm:"column:id"`
	OrderNo      string    `json:"orderNo" gorm:"column:order_no"`
	StatusName   string    `json:"statusName" gorm:"column:status_name"`
	LastUpdateBy string    `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"lastUpdate" gorm:"column:last_update"`
}

// TableName ...
func (o *OrderStatus) TableName() string {
	return "public.order_status"
}
