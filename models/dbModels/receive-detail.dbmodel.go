package dbmodels

import "time"

// ReceiveDetail ...
type ReceiveDetail struct {
	ID        int64 `json:"id" gorm:"column:id"`
	ReceiveID int64 `json:"receiveId" gorm:"column:receive_id"`

	ProductID int64   `json:"productId" gorm:"column:product_id"`
	Product   Product `json:"product" gorm:"foreignkey:id;association_foreignkey:ProductID;association_autoupdate:false;association_autocreate:false"`
	Qty       int64   `json:"qty" gorm:"column:qty"`
	Price     float32 `json:"price" gorm:"column:price"`
	Disc      float32 `json:"disc" gorm:"column:disc"`
	Hpp       float32 `json:"hpp" gorm:"column:hpp"`

	UomID int64  `json:"uomId" gorm:"column:uom"`
	UOM   Lookup `json:"uom" gorm:"foreignkey:id;association_foreignkey:UomID;association_autoupdate:false;association_autocreate:false"`

	LastUpdateBy string    `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"lastUpdate" gorm:"column:last_update"`
}

// TableName ...
func (o *ReceiveDetail) TableName() string {
	return "public.receive_detail"
}
