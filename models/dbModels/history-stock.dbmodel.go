package dbmodels

import (
	"time"
)

// HistoryStock ...
type HistoryStock struct {
	ID           int64     `json:"id" gorm:"column:id"`
	Code         string    `json:"code" gorm:"column:code"`
	Name         string    `json:"name" gorm:"column:name"`
	Debet        float32   `json:"debet" gorm:"column:debet"`
	Kredit       float32   `json:"kredit" gorm:"column:kredit"`
	Saldo        float32   `json:"saldo" gorm:"column:saldo"`
	TransDate    time.Time `json:"transDate" gorm:"column:trans_date"`
	Description  string    `json:"description" gorm:"column:description"`
	LastUpdateBy string    `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"lastUpdate" gorm:"column:last_update"`
	ReffNo       string    `json:"reffNo" gorm:"column:reff_no"`
	Price        float32   `json:"price" gorm:"column:price"`
	Hpp          float32   `json:"hpp" gorm:"column:hpp"`
}

// TableName ...
func (m *HistoryStock) TableName() string {
	return "public.history_stock"
}
