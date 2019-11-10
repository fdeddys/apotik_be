package dbmodels

import (
	"time"
)

type FollowOrder struct {
	ID                  int64       `json:"id" gorm:"column:id"`
	OrderNo             string      `json:"orderNo" gorm:"column:order_no"`
	OrderDate           time.Time   `json:"orderDate" gorm:"column:order_date"`
	MerchantCode        string      `json:"merchantCode" gorm:"column:merchant_code"`
	Merchant            Merchant    `json:"merchant" gorm:"foreignkey:code;association_foreignkey:MerchantCode"`
	SupplierCode        string      `json:"supplierCode" gorm:"column:supplier_code"`
	Supplier            Supplier    `json:"supplier" gorm:"foreignkey:code;association_foreignkey:SupplierCode"`
	Note                string      `json:"note" gorm:"column:note"`
	WarehouseCode       string      `json:"warehouseCode" gorm:"column:warehouse_code"`
	Total               int32       `json:"total" gorm:"column:total"`
	Tax                 int32       `json:"tax" gorm:"column:tax"`
	GrandTotal          int32       `json:"grandTotal" gorm:"column:grand_total"`
	DeliveryNo          string      `json:"deliveryNo" gorm:"column:delivery_no"`
	DeliveryDate        time.Time   `json:"deliveryDate" gorm:"column:delivery_date"`
	LastUpdateBy        string      `json:"lastUpdateBy" gorm:"column:last_update_by"`
	LastUpdate          time.Time   `json:"lastUpdate" gorm:"column:last_update"`
	StatusCode          string      `json:"status" gorm:"column:status_code"`
	InternalStatus      int         `json:"internalStatus" gorm:"column:internal_status"`
	Autodebet           int         `json:"autodebet" gorm:"column:autodebet_proses"`
	ManualPaymentCode   string      `json:"manualPaymentCode" gorm:"column:manual_payment_code"`
	ManualPaymentNumber string      `json:"manualPaymentNumber" gorm:"column:manual_payment_number"`
	PaymentNote         string      `json:"paymentNote" gorm:"column:payment_note"`
	ManualPaymentDate   time.Time   `json:"manualPaymentDate" gorm:"column:manual_payment_date"`
	SalesNo             string      `json:"salesNo" gorm:"column:salesman_no"`
	OrderDetail         OrderDetail `json:"orderDetail" gorm:"foreignkey:order_id; association_foreignkey:ID"`
	OrderStatus         OrderStatus `json:"orderStatus" gorm:"foreignkey:order_no;association_foreignkey:OrderNo"`
}

// TableName ...
func (f *FollowOrder) TableName() string {
	return "public.order"
}
