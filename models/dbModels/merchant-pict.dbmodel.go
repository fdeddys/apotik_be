package dbmodels

// Merchant Picture
type MerchantPict struct {
	ID 					int64 	`json:"id" gorm:"column:id"`
	LookupCode			string	`json:"lookupCode" gorm:"column:lookup_code"`
	PictPath			string 	`json:"pictPath" gorm:"column:pict_path"`
	Status				int 	`json:"status" gorm:"column:status"`
	MerchantCode		string	`json:"merchantCode" gorm:"column:merchant_code"`
}


// TableName ...
func (s *MerchantPict) TableName() string {
	return "public.merchant_pict"
}