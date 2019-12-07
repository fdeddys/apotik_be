package services

import (
	"distribution-system-be/database"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
)

// ReceiveService ...
type ReceiveService struct {
}

// GetDataPage ...
func (r ReceiveService) GetDataPage(param dto.FilterReceive, page, limit, status int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetReceivePage(param, offset, limit, status)

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

// GetDataReceiveByID ...
func (r ReceiveService) GetDataReceiveByID(reveiveID int64) dbmodels.Receive {

	var res dbmodels.Receive
	// var err error
	res, _ = database.GetReceiveByOrderID(reveiveID)

	return res
}
