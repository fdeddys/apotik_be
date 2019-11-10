package dto

// FilterOrder ...
type FilterOrder struct {
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	InternalStatus string `json:"internalStatus"`
	OrderNumber    string `json:"orderNumber"`
	SalesNo        string `json:"salesNo"`
	MerchantPhone  string `json:"merchantPhone"`
}

// FilterOrderResult ...
type FilterOrderResult struct {
	ErrDesc      string      `json:"errDesc"`
	ErrCode      string      `json:"errCode"`
	Data         interface{} `json:"data"`
	Page         int         `json:"page"`
	PerPage      int         `json:"per_page"`
	TotalPages   int         `json:"total_pages"`
	TotalRecords int         `json:"total_records"`
}

// OrderSaveResult ...
type OrderSaveResult struct {
	ErrDesc string `json:"errDesc"`
	ErrCode string `json:"errCode"`
	OrderNo string `json:"orderNo"`
}

// FilterOrderDetail ...
type FilterOrderDetail struct {
	OrderNo   string `json:"orderNo"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// OrderAutodebetRequest ...
type OrderAutodebetRequest struct {
	OrderNo        string `json:"orderNo"`
	PaymentSuccess string `json:"paymentSuccess"`
}

// OrderAutodebetResult ...
type OrderAutodebetResult struct {
	ErrDesc string `json:"errDesc"`
	ErrCode string `json:"errCode"`
}

// KafkaBody ...
type KafkaBody struct {
	Topic string     `json:"topic"`
	Data  SalesOrder `json:"data"`
}

// SalesOrder ...
type SalesOrder struct {
	AccessToken     string           `json:"access_token"`
	WarehouseCode   string           `json:"warehouse_code"`
	CustomerCode    string           `json:"customer_code"`
	TransactionAt   string           `json:"transaction_at"`
	StateAction     string           `json:"state_action"`
	SalesOrderItems []SalerOrderItem `json:"sales_order_lines"`
	Code            string           `json:"code"`
	SupplierCode    string           `json:"supplier_code"`
}

// SalerOrderItem ...
type SalerOrderItem struct {
	ItemCode    string `json:"item_code"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Uom         string `json:"unit_of_measurement_code"`
}
