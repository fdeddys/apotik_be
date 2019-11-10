package services

import (
	repository "oasis-be/database"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"time"
)

// SupplierGroupService ...
type SupplierGroupService struct {
}

// GetSupplierGroupPaging ...
func (h SupplierGroupService) GetSupplierGroupPaging(param dto.FilterSupplierGroup, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := repository.GetSupplierGroupPaging(param, offset, limit)

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

// GetSupplierGroupDetails ...
func (h SupplierGroupService) GetSupplierGroupDetails(id int) models.ContentResponse {

	var res models.ContentResponse

	data, errCode, errDesc, err := repository.GetSupplierGroupDetails(id)

	if err != nil {
		res.Contents = nil
		res.ErrCode = "02"
		res.ErrDesc = "Error query data to DB"
		return res
	}

	res.Contents = data
	res.ErrCode = errCode
	res.ErrDesc = errDesc

	return res
}

// SaveSupplierGroup ...
func (h SupplierGroupService) SaveSupplierGroup(supplierGroup *dbmodels.SupplierGroup) models.NoContentResponse {
	supplierGroup.Code = repository.GenerateCode()
	supplierGroup.LastUpdate = time.Now()
	supplierGroup.LastUpdateBy = dto.CurrUser
	// var res models.ResponseSave
	res := repository.SaveSupplierGroup(*supplierGroup)

	return res
}

// UpdateSupplierGroup ...
func (h SupplierGroupService) UpdateSupplierGroup(supplierGroup *dbmodels.SupplierGroup) models.NoContentResponse {
	supplierGroup.LastUpdate = time.Now()
	supplierGroup.LastUpdateBy = dto.CurrUser

	res := repository.UpdateSupplierGroup(*supplierGroup)

	return res
}

// GetListSupplierGroup ...
func (h SupplierGroupService) GetListSupplierGroup() models.ResponseSupplierGroup {
	var res models.ResponseSupplierGroup

	data, err := repository.GetListSuppliergroup()

	if err != nil {
		return res
	}

	res.Data = data
	return res
}
