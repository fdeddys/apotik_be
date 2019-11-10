package database

import (
	"oasis-be/models/dbModels"
	_ "strconv"
	dto "oasis-be/models/dto"
	"sync"
	"log"
	"github.com/jinzhu/gorm"
	"strings"
	"oasis-be/models"
	"oasis-be/constants"
)

// Get Data Issuer
func GetIssuerPaging(param dto.FilterName, offset int, limit int) ([]dbmodels.Issuer, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var issuer []dbmodels.Issuer
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&issuer).Error
		if err != nil {
			return issuer, 0, err
		}
		return issuer, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysIssuer(db, offset, limit, &issuer, param, errQuery)
	go AsyncQueryCountsIssuer(db, &total, param, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return issuer, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return issuer, 0, resErrCount
	}
	return issuer, total, nil
}


// AsyncQuerysCountIssuer ...
func AsyncQueryCountsIssuer(db *gorm.DB, total *int, param dto.FilterName, resChan chan error) {
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Model(&dbmodels.Issuer{}).Where("name ilike ?", searchName).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


// AsyncQuerysIssuer ...
func AsyncQuerysIssuer(db *gorm.DB, offset int, limit int, issuer *[]dbmodels.Issuer, param dto.FilterName, resChan chan error) {

	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Order("name ASC").Offset(offset).Limit(limit).Find(&issuer, "name ilike ?", searchName).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


// Repository Save
func SaveIssuer(issuer *dbmodels.Issuer) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if  r := db.Save(&issuer); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}


// Repository Update
func UpdateIssuer(issuer *dbmodels.Issuer) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if  r := db.Save(&issuer); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}


// Repository List
func GetListIssuer()([]dbmodels.Issuer, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var issuer []dbmodels.Issuer
	var err error

	err = db.Find(&issuer).Error
	if err != nil {
		return issuer, err
	}
	return issuer, nil
}


// Repository Search Issuer
func GetListIssuerBySearch(name string)([]dbmodels.Issuer, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var issuer []dbmodels.Issuer
	var err error

	err = db.Where("name ilike ? ", "%"+name+"%").Find(&issuer).Error
	if err != nil {
		return issuer, err
	}
	return issuer, nil
}


// Get Last Issuer
func GetLastIssuer()(dbmodels.Issuer, error){
	db := GetDbCon()
	db.Debug().LogMode(true)

	var issuer dbmodels.Issuer
	var err  error
	err = db.Order("code desc limit 1").Find(&issuer).Error
	if err != nil {
		return issuer, err
	}
	return issuer, nil
}