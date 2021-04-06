package services

import (
	"distribution-system-be/database"
	"distribution-system-be/models"
)

type WarehouseService struct {
}

// Get Data Customer Paging
func (m WarehouseService) GetAllWarehouse() models.ContentResponse {
	var res models.ContentResponse

	data, code, msg := database.GetAllWarehouse()
	res.ErrCode = code
	res.ErrDesc = msg
	res.Contents = data
	return res
}
