package dbmodels

import (
	"time"
)

// Merchant ...
type Merchant struct {
	ID           				int64 					`json:"id" gorm:"column:id"`
	Code     					string 					`json:"code" gorm:"column:code"`
	Name        				string 					`json:"name" gorm:"column:name"`
	IssuerCode     				int8 					`json:"issuerCode" gorm:"column:issuer_code"`
	Top       					int    					`json:"top" gorm:"column:top"`
	Status 						int 					`json:"status" gorm:"column:status"`
	LastUpdateBy				string					`json:"lastUpdateBy"`
	LastUpdate					time.Time				`json:"lastUpdate"`
	Issuer 						Issuer					`json:"Issuer" gorm:"foreignkey:ID;association_foreignkey:IssuerCode"`
	FirstOrder					int						`json:"firstOrder" gorm:"column:first_order"`
}

// TableName ...
func (m *Merchant) TableName() string {
	return "public.merchant"
}
