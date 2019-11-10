package dbmodels

import (
	"time"
)


// Supplier Price
type SupplierPrice struct {
	ID 						int64 	`json:"id" gorm:"column:id"`
	SupplierId				int64	`json:"supplier_id" gorm:"column:supplier_id"`
	ProductCode				string 	`json:"product_code" gorm:"column:product_code"`
	SellPrice				int		`json:"sell_price" gorm:"column:sell_price"`
	BuyPrice				int		`json:"buy_price" gorm:"column:buy_price"`
	SellMarginCode			string	`json:"sell_margin_code" gorm:"column:sell_margin_code"`
	LastUpdateBy			string 	`json:"last_update_by"`
	LastUpdate				time.Time
	Lookup 					Lookup	`json:"Lookup" gorm:"foreignkey:Code;association_foreignkey:SellMarginCode"`
	Product 				Product	`json:"Product" gorm:"foreignkey:Code;association_foreignkey:ProductCode"`
	PriceMargin				int		`json:"priceMargin" gorm:"column:price_margin"`
}


type RequestSupplierProduct struct {
	SupplierID 			int64		`json:"supplier_id"`
	// Reg  				bool		`json:"reg"`
	// Page				int			`json:"page"`
	MerchantID			int64		`json:"merchant_id"`
}


// TableName ...
func (s *SupplierPrice) TableName() string {
	return "public.supplier_price_list"
}