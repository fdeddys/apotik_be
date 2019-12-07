package services

import (
	"distribution-system-be/database"
	"distribution-system-be/models"
	"distribution-system-be/models/dto"
)

// ReceiveDetailService ...
type ReceiveDetailService struct {
}

// GetDataReceiveDetailPage ...
func (r ReceiveDetailService) GetDataReceiveDetailPage(param dto.FilterReceiveDetail, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetReceiveDetailPage(param, offset, limit)

	if err != nil {
		res.Error = err.Error()
		return res
	}

	res.Contents = data
	res.TotalRow = totalData
	res.Page = page
	res.Count = limit

	return res
}
