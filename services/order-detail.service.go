package services

import "distribution-system-be/models"
import "distribution-system-be/database"
import dto "distribution-system-be/models/dto"

// OrderDetailService ...
type OrderDetailService struct {
}

// GetDataOrderDetailPage ...
func (o OrderDetailService) GetDataOrderDetailPage(param dto.FilterOrderDetail, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetOrderDetailPage(param, offset, limit)

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
