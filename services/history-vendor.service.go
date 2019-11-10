package services

import (
	"oasis-be/database"
	models "oasis-be/models"
	dto "oasis-be/models/dto"
)

type HistoryVendorService struct {
}

func (h HistoryVendorService) GetHistoryVendorPaging(param dto.FilterName, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetHistoryVendorPaging(param, offset, limit)

	if err != nil {
		res.Error = err.Error()
		return res
	}

	res.Contents = data
	res.TotalRow = totalData
	res.Page = page
	res.Count = len(data)

	return res
}
