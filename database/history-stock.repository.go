package database

import (
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
)

//SaveHistory ...
func SaveHistory(history dbmodels.HistoryStock) models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	if r := db.Save(&history); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error save data to DB"
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"
	return res
}
