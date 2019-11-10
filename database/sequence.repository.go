package database

import (
	"oasis-be/constants"
	dbmodels "oasis-be/models/dbModels"
)

// AddSequence ...
func AddSequence(month, year string) (errcd string, newNumb int8, errdesc string) {
	db := GetDbCon()

	var seq dbmodels.Sequence
	var urut int8

	db.Where("year = ? and month = ?", year, month).First(&seq)

	if seq.ID == 0 {
		errCode, errDesc := NewSequence(year, month)
		return errCode, 1, errDesc
	}
	urut = seq.Seq + 1
	db.Model(&seq).Where("id = ?", seq.ID).Update("seq", urut)

	// var code = ""
	// code = constants.ERR_CODE_00
	return constants.ERR_CODE_00, urut, constants.ERR_CODE_00_MSG
}

// NewSequence ...
func NewSequence(year, month string) (errcode string, errdesc string) {
	db := GetDbCon()

	var seq dbmodels.Sequence
	seq.Month = month
	seq.Subject = "SO"
	seq.Year = year
	seq.Seq = 1
	err := db.Save(&seq)
	if err.Error != nil {
		return constants.ERR_CODE_80, constants.ERR_CODE_80_MSG
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}
