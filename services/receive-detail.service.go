package services

import (
	"distribution-system-be/constants"
	"distribution-system-be/database"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
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

// SaveReceiveDetail ...
func (r ReceiveDetailService) SaveReceiveDetail(receiveDetail *dbmodels.ReceiveDetail) (errCode string, errDesc string) {

	if _, err := database.GetReceiveByReceiveID(receiveDetail.ReceiveID); err != nil {
		return "99", err.Error()
	}

	if err, errDesc := database.SaveReceiveDetail(receiveDetail); err != constants.ERR_CODE_00 {
		return err, errDesc
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// DeleteReceiveDetailByID ...
func (r ReceiveDetailService) DeleteReceiveDetailByID(receiveDetailID int64) (errCode string, errDesc string) {

	if err, errDesc := database.DeleteReceiveDetailById(receiveDetailID); err != constants.ERR_CODE_00 {
		return err, errDesc
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}
