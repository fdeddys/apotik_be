package database

import (
	dbmodel "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"reflect"
	"sync"

	"github.com/jinzhu/gorm"
)

func GetHistoryVendorPaging(param dto.FilterName, offset int, limit int) ([]dbmodel.HistoryVendor, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var history []dbmodel.HistoryVendor
	var total int
	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&history).Error
		if err != nil {
			return history, 0, err
		}
		return history, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQueryHistory(db, offset, limit, &history, param, "module_name", errQuery)
	go AsyncQueryCount(db, &total, param, &dbmodel.HistoryVendor{}, "module_name", errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return history, 0, resErrQuery
	}

	if resErrCount != nil {
		return history, 0, resErrCount
	}

	return history, total, nil
}

func AsyncQueryHistory(db *gorm.DB, offset int, limit int, modelsWVal interface{}, param interface{}, fieldLookup string, resChan chan error) {
	modelsDump := reflect.ValueOf(modelsWVal).Interface()
	paramDump := reflect.ValueOf(param)
	strQuery := paramDump.Field(0).Interface().(string)
	// var criteriaName = ""

	// if strings.TrimSpace(strQuery) != "" {
	// 	criteriaName = strQuery
	// }

	criteriaName := strQuery
	if criteriaName == "" {
		criteriaName = "%"
	} else {
		criteriaName = "%" + strQuery + "%"
	}

	var err error
	// err = db.Set("gorm:auto_preload", true).Order("last_update desc").Offset(offset).Limit(limit).Find(modelsDump, fieldLookup+" ~* ?", criteriaName).Error

	err = db.Set("gorm:auto_preload", true).Order("last_update desc").Offset(offset).Limit(limit).Find(modelsDump, "COALESCE("+fieldLookup+",'') ILIKE ?", criteriaName).Error

	if err != nil {
		resChan <- err
	}

	resChan <- nil
}
