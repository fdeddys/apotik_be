package dto

// FilterReceive ...
type FilterReceive struct {
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	Status        string `json:"status"`
	ReceiveNumber string `json:"receiveNumber"`
}

// FilterReceiveDetail ...
type FilterReceiveDetail struct {
	ReceiveNo string `json:"receiveNo"`
	ReceiveID int64  `json:"receiveId"`
}
