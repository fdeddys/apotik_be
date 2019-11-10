package dbmodels

import (
	"time"
)

//SupplierGroup model ...
type SupplierGroup struct {
	ID           int64     `json:"id" gorm:"column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	LastUpdateBy string    `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"last_update"`
	Code         string    `json:"code" gorm:"column:code"`
}

// TableName ...
func (t *SupplierGroup) TableName() string {
	return "public.supplier_group"
}
