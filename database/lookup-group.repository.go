package database

import (
	dbmodels "distribution-system-be/models/dbModels"
)

//GetLookupGroup ...
func GetLookupGroup() ([]dbmodels.LookupGroup, string, string, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var lookupGroup []dbmodels.LookupGroup
	err := db.Model(&dbmodels.Lookup{}).Find(&lookupGroup).Error

	if err != nil {
		return nil, "02", "Error query data to DB", err
	}
	return lookupGroup, "00", "success", nil
}
