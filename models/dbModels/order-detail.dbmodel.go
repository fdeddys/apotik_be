package dbmodels

import "time"

// OrderDetail ...
type OrderDetail struct {
	ID           int64     `json:"id" gorm:"column:id"`
	OrderID      int64     `json:"orderId" gorm:"column:order_id"`
	ProductCode  string    `json:"productCode" gorm:"column:product_code"`
	Product      Product   `json:"product" gorm:"foreignkey:code;association_foreignkey:ProductCode"`
	Qty          int8      `json:"qty" gorm:"column:qty"`
	QtyReceive   int8      `json:"qtyReceive" gorm:"column:qty_receive"`
	Price        int32     `json:"price" gorm:"column:price"`
	Disc         int16     `json:"disc" gorm:"column:disc"`
	Uom          string    `json:"uomCode" gorm:"column:uom"`
	Lookup       Lookup    `json:"uom" gorm:"foreignkey:code;association_foreignkey:uom"`
	LastUpdateBy string    `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"lastUpdate" gorm:"column:last_update"`
}

// TableName ...
func (o *OrderDetail) TableName() string {
	return "public.order_detail"
}
