package services

import (
	"oasis-be/database"
	"oasis-be/models"
	dto "oasis-be/models/dto"
)

type FollowUpOrderService struct {
}

func (f FollowUpOrderService) GetFollowOrder(param dto.FilterOrder, page int, limit int, internalStatus int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetFollowUpOrder(param, offset, limit, internalStatus)

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
