package dbmodels

import "time"

// PurchaseOrderDetail ...
type PurchaseOrderDetail struct {
	ID              int64   `json:"id" gorm:"column:id"`
	PurchaseOrderID int64   `json:"purchaseOrderId" gorm:"column:po_id"`
	ProductID       int64   `json:"productId" gorm:"column:product_id"`
	Product         Product `json:"product" gorm:"foreignkey:id;association_foreignkey:ProductID;association_autoupdate:false;association_autocreate:false"`
	Qty             int64   `json:"qty" gorm:"column:qty"`
	Price           float32 `json:"price" gorm:"column:price"`
	Disc1           float32 `json:"disc1" gorm:"column:disc1"`
	Disc2           float32 `json:"disc2" gorm:"column:disc2"`
	UomID           int64   `json:"uomId" gorm:"column:uom"`
	UOM             Lookup  `json:"uom" gorm:"foreignkey:id;association_foreignkey:UomID;association_autoupdate:false;association_autocreate:false"`

	PoUomID  int64   `json:"poUomId" gorm:"column:po_uom_id"`
	PoUOM    Lookup  `json:"poUom" gorm:"foreignkey:id;association_foreignkey:PoUomID"`
	PoUOMQty int64   `json:"poUomQty" gorm:"column:po_uom_qty"`
	PoQty    int64   `json:"poQty" gorm:"column:po_qty"`
	PoPrice  float32 `json:"poPrice" gorm:"column:po_price"`

	LastUpdateBy string    `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"lastUpdate" gorm:"column:last_update"`
}

// TableName ...
func (o *PurchaseOrderDetail) TableName() string {
	return "public.po_detail"
}
