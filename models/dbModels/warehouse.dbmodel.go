package dbmodels

import "time"

//Salesman model ...
type Warehouse struct {
	ID   int64  `json:"id" gorm:"column:id"`
	Code string `json:"code" gorm:"column:code"`
	Name string `json:"name" gorm:"column:name"`

	WarehouseIn  int8 `json:"wh_in" gorm:"column:wh_in"`
	WarehouseOut int8 `json:"wh_out" gorm:"column:wh_out"`

	Status       int       `json:"status" gorm:"column:status"`
	LastUpdateBy string    `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"last_update" gorm:"column:last_update"`
}

// TableName ...
func (t *Warehouse) TableName() string {
	return "public.warehouse"
}
