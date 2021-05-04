package database

import (
	"distribution-system-be/constants"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	"distribution-system-be/models/dto"
	"fmt"
	"log"
	"strconv"
	"sync"
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

// UpdateWarehouse ...
func UpdateWarehouse(updatedWarehouse dbmodels.Warehouse) models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	var warehouse dbmodels.Warehouse
	err := db.Model(&dbmodels.Warehouse{}).Where("id=?", &updatedWarehouse.ID).First(&warehouse).Error
	if err != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error select data to DB"
	}

	warehouse.Name = updatedWarehouse.Name
	warehouse.Status = updatedWarehouse.Status
	warehouse.LastUpdateBy = updatedWarehouse.LastUpdateBy
	warehouse.LastUpdate = updatedWarehouse.LastUpdate
	warehouse.Code = updatedWarehouse.Code
	warehouse.WarehouseIn = updatedWarehouse.WarehouseIn
	warehouse.WarehouseOut = updatedWarehouse.WarehouseOut

	err2 := db.Save(&warehouse)
	if err2 != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error update data to DB"
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"

	return res
}

//GetWarehouseLike ...
func GetWarehouseLike(warehouseTerms string) ([]dbmodels.Warehouse, string, string, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var warehouse []dbmodels.Warehouse
	err := db.Model(&dbmodels.Warehouse{}).Where("name ~* ?", &warehouseTerms).Find(&warehouse).Error

	if err != nil {
		return nil, constants.ERR_CODE_51, constants.ERR_CODE_51_MSG, err
	}
	return warehouse, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, nil
}

//SaveWarehouse ...
func SaveWarehouse(warehouse dbmodels.Warehouse) models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	if warehouse.ID < 1 {
		warehouse.Code = GenerateWarehouseCode()
	}

	if r := db.Save(&warehouse); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error save data to DB"
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"
	return res
}

// GetWarehouseFilter ...
func GetWarehouseFilter(id int) ([]dbmodels.Warehouse, string, string, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var warehouse []dbmodels.Warehouse
	err := db.Model(&dbmodels.Warehouse{}).Where("id = ?", &id).First(&warehouse).Error

	if err != nil {
		return nil, "02", "Error query data to DB", err
	}
	return warehouse, "00", "success", nil
}

// GetWarehouse ...
func GetWarehouse(param dto.FilterPaging, offset int, limit int) ([]dbmodels.Warehouse, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var warehouse []dbmodels.Warehouse
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&warehouse).Error
		if err != nil {
			return warehouse, 0, err
		}
		return warehouse, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQueryCount(db, &total, param, &dbmodels.Warehouse{}, "name", errCount)
	go AsyncQuery(db, offset, limit, &warehouse, param, "name", errQuery)

	resErrCount := <-errCount
	resErrQuery := <-errQuery

	wg.Done()

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return warehouse, 0, resErrCount
	}

	if resErrQuery != nil {
		return warehouse, 0, resErrQuery
	}

	return warehouse, total, nil
}

func GenerateWarehouseCode() string {
	db := GetDbCon()
	db.Debug().LogMode(true)
	header := "W"
	defaultCode := "W001"

	var warehouses []dbmodels.Warehouse
	err := db.Model(&dbmodels.Warehouse{}).Order("id desc").Find(&warehouses).Error

	if err != nil {
		return defaultCode
	}
	if len(warehouses) > 0 {
		// fmt.Printf("ini latest code nya : %s \n", warehouse[0].Code)
		code := warehouses[0].Code
		runes := []rune(code)
		latestNumb := string(runes[1:len(code)])
		fmt.Println("latest numb-", latestNumb)
		// woprefix := strings.TrimPrefix(warehouse[0].Code, "")
		latestCode, err := strconv.Atoi(latestNumb)
		if err != nil {
			fmt.Println("error =>", err.Error())
			return defaultCode
		}
		// fmt.Printf("ini latest code nya : %d \n", latestCode)
		wpadding := fmt.Sprintf("%v%03s", header, strconv.Itoa(latestCode+1))
		// fmt.Printf("ini pake padding : %s \n", "B"+wpadding)
		return wpadding
	}
	return defaultCode

}
