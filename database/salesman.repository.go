package database

import (
	"distribution-system-be/constants"
	dbmodels "distribution-system-be/models/dbModels"
)

// GetSales ...
func GetAllSalesman() ([]dbmodels.Salesman, string, string) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var salesman []dbmodels.Salesman
	err := db.Where("status = ?", 1).Find(&salesman).Error
	if err != nil {
		return nil, constants.ERR_CODE_51, constants.ERR_CODE_51_MSG + "  " + err.Error()
	}
	return salesman, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}
