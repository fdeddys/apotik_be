package services

import (
	"distribution-system-be/database"
	"distribution-system-be/models"
)

type SalesmanService struct {
}

// Get Data Customer Paging
func (s SalesmanService) GetAllSalesman() models.ContentResponse {
	var res models.ContentResponse

	data, code, msg := database.GetAllSalesman()
	res.ErrCode = code
	res.ErrDesc = msg
	res.Contents = data
	return res
}
