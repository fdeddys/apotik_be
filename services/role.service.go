package services

import (
	repository "oasis-be/database"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
)

// RoleService ...
type RoleService struct {
}

// GetRoleFilterPaging ...
func (h RoleService) GetRoleFilterPaging(param dto.FilterName, page int, limit int) models.ResponsePagination {

	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := repository.GetRole(param, offset, limit)

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

// SaveRole ...
func (h RoleService) SaveRole(role *dbmodels.Role) models.NoContentResponse {
	res := repository.SaveRole(*role)

	return res
}

// UpdateRole ...
func (h RoleService) UpdateRole(role *dbmodels.Role) models.NoContentResponse {

	res := repository.UpdateRole(*role)

	return res
}
