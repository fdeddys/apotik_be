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
	"fmt"
)

// Get Data Supplier
func GetSupplierPaging(param dto.FilterName, offset int, limit int) ([]dbmodels.Supplier, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplier []dbmodels.Supplier
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

	go AsyncQuerysSupplier(db, offset, limit, &supplier, param, errQuery)
	go AsyncQueryCountsSupplier(db, &total, param, errCount)
	fmt.Println(errQuery)
	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return supplier, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return supplier, 0, resErrCount
	}

	for i := 0; i < len(supplier); i++ {
		isExist := CheckImage(supplier[i].Code, "supplier")
		if isExist {
			supplier[i].LogoPath = GetImage(supplier[i].Code, "supplier")
		} else {
			supplier[i].LogoPath = GetImage("no_image", "supplier")
		}
	}
	return supplier, total, nil
}


// AsyncQuerysCountSupplier ...
func AsyncQueryCountsSupplier(db *gorm.DB, total *int, param dto.FilterName, resChan chan error) {
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Model(&dbmodels.Supplier{}).Where("name ilike ?", searchName).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


// AsyncQuerysSupplier ...
func AsyncQuerysSupplier(db *gorm.DB, offset int, limit int, supplier *[]dbmodels.Supplier, param dto.FilterName, resChan chan error) {

	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Order("id asc").Offset(offset).Limit(limit).Find(&supplier, "name ilike ?", searchName).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


// Repository Save
func SaveSupplier(supplier *dbmodels.Supplier) models.ResponseSupplier {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.ResponseSupplier

	if  r := db.Save(&supplier); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
		res.Code    = ""
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG
	res.Code = supplier.Code

	return res
}


// Repository Update
func UpdateSupplier(supplier *dbmodels.Supplier) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if  r := db.Save(&supplier); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}


// Get Last Supplier
func GetLastSupplier()(dbmodels.Supplier, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplier dbmodels.Supplier
	var err  error
	err = db.Order("code desc limit 1").Find(&supplier).Error
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}

// Get Supplier By ID
func GetSupplierById(id int64)(dbmodels.Supplier) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplier dbmodels.Supplier
	var err error

	err = db.First(&supplier, id).Error
	if err != nil {
		return supplier
	}else{
		isExist := CheckImage(supplier.Code, "supplier")
		if isExist {
			supplier.LogoPath = GetImage(supplier.Code, "supplier")
		} else {
			// supplier.LogoPath = GetImage("no_image", "supplier")
			supplier.LogoPath = "../../../assets/img/no_image.png"
		}
	}

	return supplier
}

func GetListSupplier()([]dbmodels.Supplier){
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplier []dbmodels.Supplier

	err := db.Find(&supplier).Error
	if err != nil {
		return supplier
	}
	
	for i := 0; i < len(supplier); i++ {
		isExist := CheckImage(supplier[i].Code, "supplier")
		if isExist {
			supplier[i].LogoPath = GetImage(supplier[i].Code, "supplier")
		} else {
			supplier[i].LogoPath = GetImage("no_image", "supplier")
		}
	}
	
	return supplier

}