package services

import (
	repository "oasis-be/database"
	"oasis-be/models"
	"oasis-be/models/dto"
)

// DashboardService ...
type DashboardService struct {
}

// GetQtyOrder ...
func (d DashboardService) GetQtyOrder(param dto.FilterDto) models.ContentResponse {
	var res models.ContentResponse

	resContent, _ := repository.GetQtyOrd(param.StartDate, param.EndDate)

	res.Contents = resContent
	return res
}
