package services

import (
	repository "oasis-be/database"
	dbmodels "oasis-be/models/dbModels"
)

// MenuService ...
type MenuService struct {
}

// GetMenuByUser ...
func (h MenuService) GetMenuByUser(user string) []dbmodels.Menu {
	var res []dbmodels.Menu
	// var err error
	res, _ = repository.GetUserMenus(user)

	return res
}
