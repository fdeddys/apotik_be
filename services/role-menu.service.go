package services

import (
	repository "oasis-be/database"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
)

// RoleMenuService ...
type RoleMenuService struct {
}

// GetActiveMenu ...
func (h RoleMenuService) GetActiveMenu() []dbmodels.Menu {
	var res []dbmodels.Menu
	res, _ = repository.GetAllActiveMenu()

	return res
}

// GetMenuByRole ...
func (h RoleMenuService) GetMenuByRole(roleid int) []dto.RoleMenuDto {
	var res []dto.RoleMenuDto
	// var err error
	res, _ = repository.GetMenuByRole(roleid)

	return res
}

// SaveMenuByRole ...
func (h RoleMenuService) SaveMenuByRole(roleid int, menuids []int) models.Response {
	res := repository.SaveMenuByRole(roleid, menuids)
	return res
}
