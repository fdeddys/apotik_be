package dbmodels

import "time"

//Stock model ...
type Stock struct {
	ID           int64     `json:"id" gorm:"column:id"`
	ProductID    int64     `json:"product_id" gorm:"column:product_id"`
	WarehouseID  int64     `json:"warehouse_id" gorm:"column:warehouse_id"`
	Qty          int64     `json:"qty" gorm:"column:qty"`
	Hpp          float32   `json:"hpp" gorm:"column:hpp"`
	LastUpdateBy string    `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"last_update" gorm:"column:last_update"`
}

// TableName ...
func (t *Stock) TableName() string {
	return "public.stock"
}
