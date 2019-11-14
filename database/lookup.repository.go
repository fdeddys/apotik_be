package database

import (
	"distribution-system-be/constants"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
)

//GetLookupByGroup ...
func GetLookupByGroup(lookupstr string) ([]dbmodels.Lookup, string, string, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var lookup []dbmodels.Lookup
	err := db.Model(&dbmodels.Lookup{}).Where("lookup_group ~* ?", &lookupstr).Find(&lookup).Error

	if err != nil {
		return nil, "02", "Error query data to DB", err
	}
	return lookup, "00", "success", nil
}

// GetPagingLookup ...
func GetPagingLookup(param dto.FilterLookup, offset int, limit int) ([]dbmodels.Lookup, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var lookup []dbmodels.Lookup
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&lookup).Error
		if err != nil {
			return lookup, 0, err
		}
		return lookup, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQueryCount(db, &total, param, &dbmodels.Lookup{}, "lookup_group", errCount)
	// if limit == 0 {
	// 	limit = total
	// }
	go AsyncQuery(db, offset, limit, &lookup, param, "lookup_group", errQuery)

	resErrCount := <-errCount
	resErrQuery := <-errQuery

	wg.Done()
	// wg.Done()

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return lookup, 0, resErrCount
	}

	if resErrQuery != nil {
		return lookup, 0, resErrQuery
	}

	return lookup, total, nil
}

// GetLookupFilter ...
func GetLookupFilter(id int) ([]dbmodels.Lookup, string, string, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var lookup []dbmodels.Lookup
	err := db.Model(&dbmodels.Lookup{}).Where("id = ?", &id).First(&lookup).Error

	if err != nil {
		return nil, constants.ERR_CODE_00, constants.ERR_CODE_51_MSG, err
	}
	// else {
	return lookup, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, nil
	// }
}

//SaveLookup ...
func SaveLookup(lookup dbmodels.Lookup) models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	if lookup.ID < 1 {
		lookup.Code = GenerateLookupCode(lookup.LookupGroup)
	}

	fmt.Println("Lookup ====> ", lookup)
	if r := db.Save(&lookup); r.Error != nil {
		res.ErrCode = constants.ERR_CODE_51
		res.ErrDesc = constants.ERR_CODE_51_MSG
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG
	return res
}

//GetDistinctLookup ...
func GetDistinctLookup() ([]dbmodels.Lookup, string, string, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var lookup []dbmodels.Lookup
	err := db.Select("DISTINCT lookup_group").Find(&lookup).Error

	// err := db.Model(&dbmodels.Lookup{}).Where("id = ?", &id).First(&lookup).Error

	if err != nil {
		return nil, "02", "Error query data to DB", err
	}
	// else {
	return lookup, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, nil
}

// UpdateLookup ...
func UpdateLookup(updatedlookup dbmodels.Lookup) models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	var lookup dbmodels.Lookup
	err := db.Model(&dbmodels.Lookup{}).Where("id=?", &updatedlookup.ID).First(&lookup).Error
	if err != nil {
		res.ErrCode = constants.ERR_CODE_51
		res.ErrDesc = constants.ERR_CODE_51_MSG
	}

	lookup.Name = updatedlookup.Name
	lookup.Status = updatedlookup.Status
	lookup.Code = updatedlookup.Code
	lookup.LookupGroup = updatedlookup.LookupGroup

	err2 := db.Save(&lookup)
	if err2 != nil {
		res.ErrCode = constants.ERR_CODE_51
		res.ErrDesc = constants.ERR_CODE_51_MSG
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}

// GenerateLookupCode ...
func GenerateLookupCode(loogkupGroup string) string {
	db := GetDbCon()
	db.Debug().LogMode(true)

	// err := db.Order(order).First(&models)
	var lookup []dbmodels.Lookup
	err := db.Model(&dbmodels.Lookup{}).Order("id desc").First(&lookup).Error
	// err := db.Model(&dbmodels.Brand{}).Where("id = 200000").Order("id desc").First(&brand).Error

	// prefix := loogkupGroup[:2]
	if err != nil {
		return "L00001"
	}
	if len(lookup) > 0 {
		// fmt.Printf("ini latest code nya : %s \n", brand[0].Code)
		prefix := strings.TrimPrefix(lookup[0].Code, "L")
		latestCode, err := strconv.Atoi(prefix)
		if err != nil {
			fmt.Printf("error")
			return "L00001"
		}
		// fmt.Printf("ini latest code nya : %d \n", latestCode)
		wpadding := fmt.Sprintf("%05s", strconv.Itoa(latestCode+1))
		// fmt.Printf("ini pake padding : %s \n", "B"+wpadding)
		return "L" + wpadding
	}
	return "L00001"

}
