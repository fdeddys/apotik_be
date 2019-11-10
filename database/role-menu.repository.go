package database

import (
	"fmt"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	"oasis-be/models/dto"
	"strconv"
	"time"
)

// GetAllActiveMenu ...
func GetAllActiveMenu() ([]dbmodels.Menu, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var menus []dbmodels.Menu
	var err error

	err = db.Find(&menus, "status = ?", 1).Error

	fmt.Println("Menus => ", menus)

	if err != nil {
		return menus, err
	}
	return menus, nil
}

// GetMenuByRole ...
func GetMenuByRole(roleID int) ([]dto.RoleMenuDto, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var menus []dto.RoleMenuDto
	var err error

	// err = db.Find(&menus, "status = ?", 1).Error

	err = db.Raw(`
		select b.description as menuDescription, a.status, a.menu_id as menuId
		from m_role_menu a
		left join m_menus b on a.menu_id = b.id
		where a.role_id = ?
		and b.status = 1
	`, roleID).Scan(&menus).Error

	fmt.Println("Menus => ", menus)

	if err != nil {
		return menus, err
	}
	return menus, nil
}

// SaveMenuByRole ...
func SaveMenuByRole(roleID int, menuIds []int) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if len(menuIds) == 0 {
		res.ErrCode = "05"
		res.ErrDesc = "Error, Menu Id empty"
		return res
	}

	var roleIDStr string
	roleIDStr = strconv.Itoa(roleID)
	query := "DELETE FROM m_role_menu " +
		" WHERE role_id = " + roleIDStr + ";"
	dbres := db.Exec(query)
	if dbres.Error != nil {
		res.ErrCode = "05"
		res.ErrDesc = "Error, Role Id not found"
		return res
	}

	errCount := false
	for _, id := range menuIds {
		saveData := dbmodels.RoleMenu{}
		saveData.RoleID = roleID
		saveData.MenuID = id
		saveData.Status = 1
		saveData.LastUpdateBy = 1
		saveData.LastUpdate = time.Now()
		dbres := db.Save(&saveData)
		if dbres.Error != nil {
			errCount = true
			break
		}
	}

	if errCount {
		var roleIDStr string
		roleIDStr = strconv.Itoa(roleID)
		query := "DELETE FROM m_role_menu " +
			" WHERE role_id = " + roleIDStr + ";"
		dbres := db.Exec(query)
		if dbres.Error != nil {
			res.ErrCode = "05"
			res.ErrDesc = "Error, Role Id not found"
			return res
		}
		res.ErrCode = "05"
		res.ErrDesc = "Error save menu"
		return res
	}

	res.ErrCode = "00"
	res.ErrDesc = "Save Success"
	return res
}
