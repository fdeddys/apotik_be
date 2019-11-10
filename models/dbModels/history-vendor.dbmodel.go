package dbmodels

import (
	"time"
)

type HistoryVendor struct {
	Id           int64     `json:"id"  gorm:"PRIMARY KEY;column:id"`
	Code         string    `json:"code" gorm:"column:code"`
	ModuleName   string    `json:"module_name" gorm:"column:module_name"`
	StatusName   string    `json:"status_name" gorm:"column:status_name"`
	LastUpdateBy string    `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"last_update"`
}

// Table
func (h *HistoryVendor) TableName() string {
	return "public.history_vendor"
}
