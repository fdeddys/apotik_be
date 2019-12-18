package dbmodels

import "time"

// Receive ...
type Receive struct {
	ID          int64     `json:"id" gorm:"column:id"`
	ReceiveNo   string    `json:"receiveNo" gorm:"column:receive_no"`
	ReceiveDate time.Time `json:"receiveDate" gorm:"column:receive_date"`

	SupplierID int64    `json:"supplierId" gorm:"column:supplier_id"`
	Supplier   Supplier `json:"supplier" gorm:"foreignkey:id;association_foreignkey:SupplierID;association_autoupdate:false;association_autocreate:false"`

	Note       string  `json:"note" gorm:"column:note"`
	Tax        float32 `json:"tax" gorm:"column:tax"`
	Total      float32 `json:"total" gorm:"column:total"`
	GrandTotal float32 `json:"grandTotal" gorm:"column:grand_total"`
	// status
	// 10 = new order
	// 20 = approve
	// 30 = reject
	// 40 = paid
	Status       int8      `json:"status" gorm:"column:status"`
	LastUpdateBy string    `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"lastUpdate" gorm:"column:last_update"`
}

// TableName ...
func (o *Receive) TableName() string {
	return "public.receive"
}