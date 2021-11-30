package services

import (
	"distribution-system-be/constants"
	"distribution-system-be/database"
	dbmodels "distribution-system-be/models/dbModels"
)

// ParameterService ...
type ParameterService struct {
}

// GetDataOrderById ...
func (p ParameterService) GetByName(paramName string) dbmodels.Parameter {

	// var res dbmodels.Parameter
	// var err error
	res, errCode, _, _ := database.GetParameterByNama(paramName)
	if errCode == constants.ERR_CODE_00 {
		return res
	}

	return dbmodels.Parameter{ID: 0, Name: "", Value: ""}
}
