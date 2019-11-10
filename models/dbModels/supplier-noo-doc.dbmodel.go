package dbmodels

import (
	"time"
)

type SupplierNooDoc struct {
	ID   				int64				`json:"id" gorm:"column:id"`
	SupplierID			int64 				`json:"supplierID" gor:"column:supplier_id"`
	LookupCode			string				`json:"lookupCode" gorm:"column:lookup_code"`
	MerchantCode		string				`json:"merchantCode" gorm:"column:merchant_code"`
	LastUpdateBy		string				`json:"lastUpdateBy"`
	LastUpdate			time.Time
	Merchant 			Merchant			`gorm:"foreignkey:MerchantCode,association_foreignkey:Code"`
	ApprovalStatus		string 				`json:"approvalStatus" gorm:"column:approval_status"`
	MerchantPict		[]MerchantPict		`gorm:"many2many:SupplierNOODoc;foreignkey:merchant_code;association_foreignkey:merchant_code;association_jointable_foreignkey:merchant_code;jointable_foreignkey:merchant_code;"`
	Supplier 			Supplier			`json:"Supplier" gorm:"foreignkey:SupplierID,association_foreignkey:ID"`
}


type MerchantPictures struct {
	LookupCode 		string		`json:"lookupCode"`
	PicPath			string		`json:"pictPath"`
	LookupName		string		`json:"lookupName"`
}


type RequestSupplierNooDoc struct {
	ID   				int64				`json:"id" gorm:"column:id"`
	SupplierID			string 				`json:"supplier_id"`
	MerchantID			string				`json:"merchant_id"`
}

type ApproveSupplierNooDoc struct {
	ID   				int64				`json:"id" gorm:"column:id"`
	SupplierID			int 				`json:"supplier_id"`
	MerchantID			int					`json:"merchant_id"`
	ApprovalStatus		string				`json:"approval_status"`
}

type ReqNooDocByStatus struct {
	SupplierID		int64 		`json:"supplierID"`
	MerchantCode    string		`json:"merchantCode"`
}

// TableName ...
func (s *SupplierNooDoc) TableName() string {
	return "public.supplier_noo_doc"
}