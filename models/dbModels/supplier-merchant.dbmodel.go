package dbmodels

import (
	"time"
)

// Supplier Merchant
type SupplierMerchant struct {
	ID 					int64 		`json:"id" gorm:"column:id"`
	MerchantCode		string 		`json:"merchant_code" gorm:"column:merchant_code"`
	LastUpdateBy		string 		`json:"last_update_by"`
	LastUpdate			time.Time	`json:"last_update"`
	SupplierId			int64 		`json:"supplier_id" gorm:"column:supplier_id"`
	Merchant     		Merchant  	`json:"Merchant" gorm:"foreignkey:Code;association_foreignkey:MerchantCode"`
	FirstOrder   int8      `json:"first_order"`
}

type SupplierMerchantReq struct {
	SupplierID   	int64  		`json:"supplier_id"`
	MerchantID		int64			`json:"merchant_id"`
}

// TableName ...
func (s *SupplierMerchant) TableName() string {
	return "public.supplier_merchant_list"
}
