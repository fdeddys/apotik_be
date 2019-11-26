package dbmodels

import "time"

// SalesOrderDetail ...
type SalesOrderDetail struct {
	ID           int64 `json:"id" gorm:"column:id"`
	SalesOrderID int64 `json:"salesOrderId" gorm:"column:sales_order_id"`

	ProductID int64   `json:"productId" gorm:"column:product_id"`
	Product   Product `json:"product" gorm:"foreignkey:id;association_foreignkey:ProductID;association_autoupdate:false;association_autocreate:false"`
	Qty       int32   `json:"qty" gorm:"column:qty"`
	Price     int32   `json:"price" gorm:"column:price"`
	Disc      int16   `json:"disc" gorm:"column:disc"`

	UomID int64  `json:"uomId" gorm:"column:uom"`
	UOM   Lookup `json:"uom" gorm:"foreignkey:id;association_foreignkey:UomID;association_autoupdate:false;association_autocreate:false"`

	LastUpdateBy string    `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"lastUpdate" gorm:"column:last_update"`
}

// TableName ...
func (o *SalesOrderDetail) TableName() string {
	return "public.sales_order_detail"
}
