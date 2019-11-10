package database

import (
	"fmt"
	"log"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"strconv"
	"strings"
	"sync"
)

// GetSupplierGroupDetails ...
func GetSupplierGroupDetails(id int) ([]dbmodels.SupplierGroup, string, string, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierGroup []dbmodels.SupplierGroup
	err := db.Model(&dbmodels.SupplierGroup{}).Where("id = ?", &id).First(&supplierGroup).Error

	if err != nil {
		return nil, "02", "Error query data to DB", err
	}
	// else {
	return supplierGroup, "00", "success", nil
	// }
}

// GetSupplierGroupPaging ...
func GetSupplierGroupPaging(param dto.FilterSupplierGroup, offset int, limit int) ([]dbmodels.SupplierGroup, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierGroup []dbmodels.SupplierGroup
	var total int
	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&supplierGroup).Error
		if err != nil {
			return supplierGroup, 0, err
		}
		return supplierGroup, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuery(db, offset, limit, &supplierGroup, param, "name", errQuery)
	go AsyncQueryCount(db, &total, param, &dbmodels.SupplierGroup{}, "name", errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return supplierGroup, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("err-->", resErrCount)
		return supplierGroup, 0, resErrCount
	}

	return supplierGroup, total, nil
}

// UpdateSupplierGroup ...
func UpdateSupplierGroup(updateSupplierGroup dbmodels.SupplierGroup) models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierGroup dbmodels.SupplierGroup
	err := db.Model(&dbmodels.SupplierGroup{}).Where("id=?", &updateSupplierGroup.ID).First(&supplierGroup).Error
	if err != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error select data to DB"
	}

	supplierGroup.Name = updateSupplierGroup.Name
	supplierGroup.LastUpdateBy = updateSupplierGroup.LastUpdateBy
	supplierGroup.LastUpdate = updateSupplierGroup.LastUpdate
	supplierGroup.Code = updateSupplierGroup.Code

	err2 := db.Save(&supplierGroup)
	if err2 != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error update data to DB"
	}

	res.ErrCode = "00"
	res.ErrDesc = "success"

	return res
}

//SaveSupplierGroup ...
func SaveSupplierGroup(supplierGroup dbmodels.SupplierGroup) models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	if r := db.Save(&supplierGroup); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error save data to DB"
	}

	res.ErrCode = "00"
	res.ErrDesc = "success"
	return res
}

// GetListSupplierGroup
func GetListSuppliergroup() ([]dbmodels.SupplierGroup, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierGroup []dbmodels.SupplierGroup
	var err error

	err = db.Find(&supplierGroup).Error
	if err != nil {
		return supplierGroup, err
	}
	return supplierGroup, nil
}

func GenerateCode() string {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierGroup []dbmodels.SupplierGroup
	err := db.Order("code desc").First(&supplierGroup).Error

	if err != nil {
		return "SG000001"
	}

	if len(supplierGroup) > 0 {
		prefix := strings.TrimPrefix(supplierGroup[0].Code, "SG")
		lastCode, err := strconv.Atoi(prefix)

		if err != nil {
			fmt.Println("error")
			return "SG000001"
		}
		newCode := fmt.Sprintf("%06s", strconv.Itoa(lastCode+1))

		return "SG" + newCode

	}
	return "SG000001"
}
