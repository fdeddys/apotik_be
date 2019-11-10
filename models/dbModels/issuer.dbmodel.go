package dbmodels

import (
	"time"
)


// Issuer
type Issuer struct {
	ID           	int64  	`json:"id" gorm:"PRIMARY_KEY;column:id"`
	Code			string	`json:"code" gorm:"column:code"`
	Name			string 	`json:"name" gorm:"column:name"`
	LastUpdateBy	string 	`json:"last_update_by"`
	LastUpdate		time.Time
}


// TableName ...
func (i *Issuer) TableName() string {
	return "public.issuer"
}