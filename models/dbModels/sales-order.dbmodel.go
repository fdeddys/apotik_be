package dbmodels

import "time"

// SalesOrder ...
type SalesOrder struct {
	ID           int64     `json:"id" gorm:"column:id"`
	SalesOrderNo string    `json:"salesOrderNo" gorm:"column:sales_order_no"`
	OrderDate    time.Time `json:"orderDate" gorm:"column:order_date"`

	CustomerID int64    `json:"customerId" gorm:"column:customer_id"`
	Customer   Customer `json:"customer" gorm:"foreignkey:id;association_foreignkey:CustomerID;association_autoupdate:false;association_autocreate:false"`

	Note         string    `json:"note" gorm:"column:note"`
	LastUpdateBy string    `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"lastUpdate" gorm:"column:last_update"`
	Tax          int64     `json:"tax" gorm:"column:tax"`
	Total        int64     `json:"total" gorm:"column:total"`
	GrandTotal   int32     `json:"grandTotal" gorm:"column:grand_total"`

	SalesmanID int64 `json:"salesmanId" gorm:"column:salesman_id"`
	Salesman   User  `json:"salesman" gorm:"foreignkey:id;association_foreignkey:SalesmanID;association_autoupdate:false;association_autocreate:false"`

	// status
	// 10 = new order
	// 20 = approve
	// 30 = reject
	// 40 = paid
	Status int8 `json:"status" gorm:"column:status"`
	Top    int8 `json:"top" gorm:"column:top"`
	IsCash int8 `json:"isCash" gorm:"column:is_cash"`
}

// TableName ...
func (o *SalesOrder) TableName() string {
	return "public.sales_order"
}
