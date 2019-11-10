package dbmodels

import (
	"time"
)

//Product model ...
type Product struct {
	ID             int64  `json:"id" gorm:"column:id"`
	Code           string `json:"code" gorm:"column:code"`
	Name           string `json:"name" gorm:"column:name"`
	ProductGroup   ProductGroup
	ProductGroupID int `json:"product_group_id" gorm:"column:product_group_id"`
	BrandID        int `json:"brand_id" gorm:"column:brand_id"`
	Brand          Brand
	UOM            string    `json:"uom" gorm:"column:uom"`
	UomLookup      Lookup    `gorm:"foreignkey:UOM; association_foreignkey:Code"`
	IMG1           string    `json:"img1" gorm:"column:img1"`
	Status         int       `json:"status" gorm:"column:status"`
	LastUpdateBy   string    `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate     time.Time `json:"last_update"`
	// StockStatus    string    `json:"stock_status" gorm:"column:stock_status"`
	// StockLookup     Lookup    `gorm:"foreignkey:StockStatus; association_foreignkey:Code"`
	// IMG string `json:"img" gorm:"column:img1"`
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
