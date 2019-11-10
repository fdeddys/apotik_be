package database

import (
	"oasis-be/models/dbModels"
	_ "strconv"
	dto "oasis-be/models/dto"
	"sync"
	"github.com/jinzhu/gorm"
	"strings"
	"oasis-be/models"
	"oasis-be/constants"
)

func GetNooDocPaging(param dto.FilterSupplierNooDocDto, offset int, limit int) ([]dbmodels.SupplierNooDoc, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplier []dbmodels.SupplierNooDoc
	var total int
	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&supplier).Error
		if err != nil {
			return supplier, 0, err
		}
		return supplier, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysNooDoc(db, offset, limit, &supplier, param, errQuery)
	go AsyncQueryCountsNooDoc(db, &total, param, errCount)
	
	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return supplier, 0, resErrQuery
	}

	if resErrCount != nil {
		return supplier, 0, resErrCount
	}

	lookup, _, _, _ := GetLookupByGroup("NOO_DOCUMENT")
	

	for i:=0;i<len(supplier);i++ {
		for j:=0;j<len(supplier[i].MerchantPict);j++ {
			for k:=0; k<len(lookup); k++{
				if lookup[k].Code == supplier[i].MerchantPict[j].LookupCode {
					splitName := strings.TrimPrefix(lookup[k].Name, "img ")
					isExist := CheckImage(splitName + "/" + supplier[i].MerchantCode, "supplier")
					if isExist {
						supplier[i].MerchantPict[j].PictPath = GetImage(splitName + "/" + supplier[i].MerchantCode, "supplier")
					} else {
						supplier[i].MerchantPict[j].PictPath = GetImage("no_image", "supplier")
					}
				}
			}

		}
	}

	return supplier, total, nil
}


func AsyncQueryCountsNooDoc(db *gorm.DB, total *int, param dto.FilterSupplierNooDocDto, resChan chan error) {
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Model(&dbmodels.SupplierNooDoc{}).Preload("Supplier").Preload("MerchantPict").Preload("Merchant").Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


func AsyncQuerysNooDoc(db *gorm.DB, offset int, limit int, supplier *[]dbmodels.SupplierNooDoc, param dto.FilterSupplierNooDocDto, resChan chan error) {
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Order("id ASC").Offset(offset).Preload("Supplier").Preload("MerchantPict").Limit(limit).Preload("Merchant").Find(&supplier).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
	
}


func UpdateNooDoc(nooDoc *dbmodels.SupplierNooDoc) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if r := db.Save(&nooDoc); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}