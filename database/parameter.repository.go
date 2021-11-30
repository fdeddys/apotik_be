package database

import (
	"distribution-system-be/constants"
	dbmodels "distribution-system-be/models/dbModels"
)

// GetParameterByNama ...
func GetParameterByNama(nama string) (dbmodels.Parameter, string, string, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var parameter dbmodels.Parameter
	err := db.Model(&dbmodels.Parameter{}).Where("name = ?", &nama).First(&parameter).Error

	if err != nil {
		return parameter, constants.ERR_CODE_51, constants.ERR_CODE_51_MSG, err
	}
	// else {
	return parameter, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, nil
	// }
}

func GetParameter() ([]dbmodels.Parameter, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var parameter []dbmodels.Parameter
	err := db.Model(&dbmodels.Parameter{}).Find(&parameter).Error

	if err != nil {
		return parameter, err
	}
	// else {
	return parameter, nil
	// }
}
