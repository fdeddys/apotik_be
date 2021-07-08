package services

import (
	"distribution-system-be/constants"
	"distribution-system-be/database"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	"distribution-system-be/models/dto"
)

// PurchaseOrderDetailService ...
type PurchaseOrderDetailService struct {
}

// GetDataPurchaseOrderDetailPage ...
func (r PurchaseOrderDetailService) GetDataPurchaseOrderDetailPage(param dto.FilterPurchaseOrderDetail, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetPurchaseOrderDetailPage(param, offset, limit)

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

// SavePurchaseOrderDetail ...
func (r PurchaseOrderDetailService) SavePurchaseOrderDetail(purchaseOrderDetail *dbmodels.PurchaseOrderDetail) (errCode string, errDesc string) {

	if _, err := database.GetPurchaseOrderByPurchaseOrderID(purchaseOrderDetail.PurchaseOrderID); err != nil {
		return "99", err.Error()
	}

	if err, errDesc := database.SavePurchaseOrderDetail(purchaseOrderDetail); err != constants.ERR_CODE_00 {
		return err, errDesc
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// DeletePurchaseOrderDetailByID ...
func (r PurchaseOrderDetailService) DeletePurchaseOrderDetailByID(purchaseOrderDetailID int64) (errCode string, errDesc string) {

	if err, errDesc := database.DeletePurchaseOrderDetailById(purchaseOrderDetailID); err != constants.ERR_CODE_00 {
		return err, errDesc
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}
