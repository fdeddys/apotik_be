package dbmodels

import "time"

//Adjustment model ...
type Adjustment struct {
	ID             int64     `json:"id" gorm:"column:id"`
	AdjustmentNo   string    `json:"adjustmentNo" gorm:"column:adjustment_no"`
	AdjustmentDate time.Time `json:"adjustmentDate" gorm:"column:adjustment_date"`

	Note  string  `json:"note" gorm:"column:note"`
	Total float32 `json:"total" gorm:"column:total"`

	// status
	// 0 = new Rec
	// 10 = approve
	// 20 = reject
	Status       int8      `json:"status" gorm:"column:status"`
	LastUpdateBy string    `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"lastUpdate" gorm:"column:last_update"`
}

// TableName ...
func (t *Adjustment) TableName() string {
	return "public.adjustment"
}
