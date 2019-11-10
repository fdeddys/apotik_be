package dbmodels

import (
	"time"
)


// Supplier
type Supplier struct {
	ID 				int64 	`json:"id" gorm:"column:id"`
	Code			string	`json:"code" gorm:"column:code"`
	Name			string 	`json:"name" gorm:"column:name"`
	Alamat			string	`json:"alamat" gorm:"column:alamat"`
	Kota			string 	`json:"kota" gorm:"column:kota"`
	Status			int		`json:"status" gorm:"column:status"`
	PicName			string	`json:"picName" gorm:"column:pic_name"`
	PicPhone		string  `json:"picPhone" gorm:"column:pic_phone"`
	LogoPath		string 	`json:"logo_path" gorm:"column:logo_path"`
	Email			string	`json:"email" gorm:"column:email"`
	Position		string	`json:"position" gorm:"column:position"`
	BankName		string	`json:"bankName" gorm:"column:bank_name"`
	BankAccountName	string	`json:"bankAccountName" gorm:"column:bank_no"`
	HostUrl			string	`json:"hostUrl" gorm:"column:host_url "`
	LastUpdateBy	string 	`json:"last_update_by"`
	LastUpdate		time.Time
	Margin			int		`json:"margin" gorm:"column:margin"`
}


type RequestSupplierRegister struct {
	SupplierID				int64 				`json:"supplier_id"`
	MerchantID				int64					`json:"merchant_id"`
	MerchantName			string				`json:"merchant_name"`
	DocLabels				[]string			`json:"doc_labels"`
	DocImages				[]string 			`json:"doc_images"`
}


// TableName ...
func (s *Supplier) TableName() string {
	return "public.supplier"
}