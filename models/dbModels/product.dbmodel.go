package dbmodels

import (
	"time"
)

//Product model ...
type Product struct {
	ID   int64  `json:"id" gorm:"column:id"`
	Code string `json:"code" gorm:"column:code"`
	Name string `json:"name" gorm:"column:name"`

	ProductGroupID int64        `json:"product_group_id" gorm:"column:product_group_id"`
	ProductGroup   ProductGroup `gorm:"foreignkey:id; association_foreignkey:ProductGroupID"`

	BrandID int64 `json:"brandId" gorm:"column:brand_id"`
	Brand   Brand `gorm:"foreignkey:id; association_foreignkey:BrandID"`

	SmallUomID int64  `json:"smallUomId" gorm:"column:small_uom_id"`
	SmallUom   Lookup `gorm:"foreignkey:id; association_foreignkey:SmallUomID"`

	BigUomID int64  `json:"bigUomId" gorm:"column:big_uom_id"`
	BigUom   Lookup `gorm:"foreignkey:id; association_foreignkey:BigUomID"`

	Status       int       `json:"status" gorm:"column:status"`
	LastUpdateBy string    `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"lastUpdate"  gorm:"column:last_update"`
	QtyUom       int16     `json:"qtyUom" gorm:"column:qty_uom"`
	QtyStock     int16     `json:"qtyStock" gorm:"column:qty_stock"`
	Hpp          int64     `json:"hpp" gorm:"column:hpp"`
}

// TableName ...
func (t *Product) TableName() string {
	return "public.product"
}

// KafkaReq ...
type KafkaReq struct {
	Topic string        `json:"topic"`
	Data  ProductVendor `json:"data"`
}

// ProductVendor ...
type ProductVendor struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Uom  string `json:"unit_of_measurement_code"`
}
