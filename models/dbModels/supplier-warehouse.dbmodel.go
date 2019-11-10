package dbmodels

import (
	"time"
)

// Supplier Warehouse
type SupplierWarehouse struct {
	ID          int64   `json:"id" gorm:"column:id"`
	SupplierId  int64  `json:"supplier_id" gorm:"column:supplier_id"`
	Description string `json:"description" gorm:"column:description"`
	Url         string `json:"url" gorm:"column:url"`
	Code        string `json:"code" gorm:"column:code"`
	// Supplier		Supplier 	`json:"preload"`
	LastUpdateBy string    `json:"last_update_by"`
	LastUpdate   time.Time `json:"last_update"`
	// Status		int 	`json:"status"`
}

// SupplierWarehouseReq ...
type SupplierWarehouseReq struct {
	SupplierCode string `json:"supplier_code" gorm:"supplier_code"`
}

// SupplierWarehouseRequest Endpoint
type SupplierWarehouseEndpoint struct {
	Topic 			string				`json:"topic"`
	Data			SWEndpointReq		`json:"data"`
}

// Field Data Supplier Warehouse Endpoint
type SWEndpointReq struct {
	AccessToken			string		`json:"access_token"`
	Name 				string 		`json:"name"`
	StateAction			string		`json:"state_action"`
	Code				string		`json:"code"`
}

// TableName ...
func (s *SupplierWarehouse) TableName() string {
	return "public.supplier_warehouset_list"
}
