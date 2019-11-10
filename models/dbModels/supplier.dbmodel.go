package dbmodels

import (
	"time"
)

// Supplier
type Supplier struct {
	ID              int64  `json:"id" gorm:"column:id"`
	Code            string `json:"code" gorm:"column:code"`
	Name            string `json:"name" gorm:"column:name"`
	Alamat          string `json:"alamat" gorm:"column:alamat"`
	Kota            string `json:"kota" gorm:"column:kota"`
	Status          int    `json:"status" gorm:"column:status"`
	PicName         string `json:"picName" gorm:"column:pic_name"`
	PicPhone        string `json:"picPhone" gorm:"column:pic_phone"`
	Tax             int    `json:"tax" gorm:"column:tax"`
	BankID          string `json:"bankId" gorm:"column:bank_id"`
	Bank            Lookup `json:"bank" gorm:"foreignkey:id;association_foreignkey:BankID"`
	BankAccountName string `json:"bankAccountName" gorm:"column:bank_acc_name"`
	BankAccountNo   string `json:"bankAccountNo" gorm:"column:bank_acc_no"`

	LastUpdateBy string `json:"last_update_by"`
	LastUpdate   time.Time
}

// TableName ...
func (s *Supplier) TableName() string {
	return "public.supplier"
}
