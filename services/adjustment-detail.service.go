package services

import (
	"distribution-system-be/constants"
	"distribution-system-be/database"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	"distribution-system-be/models/dto"
)

// AdjustmentDetailService ...
type AdjustmentDetailService struct {
}

// GetDataAdjustmentDetailPage ...
func (r AdjustmentDetailService) GetDataAdjustmentDetailPage(param dto.FilterAdjustmentDetail, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetAdjustmentDetailPage(param, offset, limit)

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

// SaveAdjustmentDetail ...
func (r AdjustmentDetailService) SaveAdjustmentDetail(adjustmentDetail *dbmodels.AdjustmentDetail) (errCode string, errDesc string) {

	if _, err := database.GetReceiveByReceiveID(adjustmentDetail.AdjustmentID); err != nil {
		return "99", err.Error()
	}

	if err, errDesc := database.SaveAdjustmentDetail(adjustmentDetail); err != constants.ERR_CODE_00 {
		return err, errDesc
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// DeleteAdjustmentDetailByID ...
func (r AdjustmentDetailService) DeleteAdjustmentDetailByID(adjustmentDetailID int64) (errCode string, errDesc string) {

	if err, errDesc := database.DeleteAdjustmentDetailById(adjustmentDetailID); err != constants.ERR_CODE_00 {
		return err, errDesc
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}
