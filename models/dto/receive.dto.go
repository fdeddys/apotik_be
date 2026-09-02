package dto

import "time"

// FilterReceive ...
type FilterReceive struct {
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	Status        int    `json:"status"`
	ReceiveNumber string `json:"receiveNumber"`
	SupplierName  string `json:"supplierName"`
	PurchaseOrderNo string `json:"purchaseOrderNo"`
}

// FilterReceiveDetail ...
type FilterReceiveDetail struct {
	ReceiveNo string `json:"receiveNo"`
	ReceiveID int64  `json:"receiveId"`
}

// ReceiveSaveResult ...
type ReceiveSaveResult struct {
	ErrDesc   string `json:"errDesc"`
	ErrCode   string `json:"errCode"`
	ReceiveNo string `json:"receiveNo"`
	Status    int8   `json:"status"`
	ID        int64  `json:"id"`
}

// ReceiveDetailSaveResult ...
type ReceiveDetailSaveResult struct {
	ErrDesc string `json:"errDesc"`
	ErrCode string `json:"errCode"`
	ID      int64  `json:"id"`
}

// ReceiveDetailResult ...
type ReceiveDetailResult struct {
	ErrDesc string `json:"errDesc"`
	ErrCode string `json:"errCode"`
}

// FilterPurchasePrice ...
type FilterPurchasePrice struct {
	ProductID   int64  `json:"productId"`
	ProductName string `json:"productName"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

// PurchasePriceRow ...
type PurchasePriceRow struct {
	SupplierName string    `json:"supplierName"`
	ReceiveNo    string    `json:"receiveNo"`
	ReceiveDate  time.Time `json:"receiveDate"`
	ProductName  string    `json:"productName"`
	UomName      string    `json:"uomName"`
	Qty          int64     `json:"qty"`
	Price        float32   `json:"price"`
	Disc1        float32   `json:"disc1"`
	Disc2        float32   `json:"disc2"`
	Tax          float32   `json:"tax"`
}

// SupplierHeaderDto ...
type SupplierHeaderDto struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// PriceDetailDto ...
type PriceDetailDto struct {
	NetPrice    float32   `json:"netPrice"`
	Price       float32   `json:"price"`
	Disc1       float32   `json:"disc1"`
	Disc2       float32   `json:"disc2"`
	Tax         float32   `json:"tax"`
	ReceiveDate time.Time `json:"receiveDate"`
	ReceiveNo   string    `json:"receiveNo"`
}

// ProductMatrixRowDto ...
type ProductMatrixRowDto struct {
	ProductID   int64                    `json:"productId"`
	ProductCode string                   `json:"productCode"`
	ProductName string                   `json:"productName"`
	UomName     string                   `json:"uomName"`
	Prices      map[int64]PriceDetailDto `json:"prices"`
}

// PurchasePriceMatrixResponse ...
type PurchasePriceMatrixResponse struct {
	Suppliers []SupplierHeaderDto   `json:"suppliers"`
	Products  []ProductMatrixRowDto `json:"products"`
	TotalRow  int                   `json:"totalRow"`
	Page      int                   `json:"page"`
	Count     int                   `json:"count"`
	Error     string                `json:"error,omitempty"`
}

// FilterPurchasePriceMatrix ...
type FilterPurchasePriceMatrix struct {
	ProductName string `json:"productName"`
}

