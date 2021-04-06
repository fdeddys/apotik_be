package database

import (
	"distribution-system-be/constants"
	dbmodels "distribution-system-be/models/dbModels"
)

// GetAllWarehouse ...
func GetAllWarehouse() ([]dbmodels.Warehouse, string, string) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var warehouse []dbmodels.Warehouse
	// err := db.Model(&dbmodels.Warehouse{}).Find(&warehouse).Error

	err := db.Where("status = ?", 1).Find(&warehouse).Error

	if err != nil {
		return nil, constants.ERR_CODE_51, constants.ERR_CODE_51_MSG + "  " + err.Error()
	}
	// else {
	return warehouse, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
	// }
}
