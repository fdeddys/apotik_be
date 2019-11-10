package database

import (
	"log"
	"oasis-be/models"
	"oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	_ "strconv"
	"strings"
	"sync"

	"oasis-be/constants"

	"github.com/jinzhu/gorm"
)

// Get Data Merchant
func GetMerchantPaging(param dto.FilterName, offset int, limit int) ([]dbmodels.Merchant, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var merchant []dbmodels.Merchant
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&merchant).Error
		if err != nil {
			return merchant, 0, err
		}
		return merchant, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysMerchant(db, offset, limit, &merchant, param, errQuery)
	go AsyncQueryCountsMerchant(db, &total, param, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return merchant, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return merchant, 0, resErrCount
	}
	return merchant, total, nil
}

func AsyncQueryCountsMerchant(db *gorm.DB, total *int, param dto.FilterName, resChan chan error) {
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}
	// err := db.Table("issuer").Select("issuer.*, merchant.*").Joins("left join issuer on merchant.issuer_code = issuer.id").Model(&dbmodels.Merchant{}).Where("merchant.name ilike ?", searchName).Count(&*total).Error
	err := db.Preload("Issuer").Model(&dbmodels.Merchant{}).Where("name ilike ?", searchName).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil

}

// AsyncQuerys ...
func AsyncQuerysMerchant(db *gorm.DB, offset int, limit int, merchant *[]dbmodels.Merchant, param dto.FilterName, resChan chan error) {

	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Preload("Issuer").Order("name ASC").Offset(offset).Limit(limit).Find(&merchant, "name like ?", searchName).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// Repository Save
func SaveMerchant(merchant *dbmodels.Merchant) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if r := db.Save(&merchant); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}

// Repository Update
func UpdateMerchant(merchant *dbmodels.Merchant) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if r := db.Save(&merchant); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}

func GetListMerchant() ([]dbmodels.Merchant, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var merchant []dbmodels.Merchant
	var err error

	err = db.Find(&merchant).Error
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

// Get Last Merchant
func GetLastMerchant() (dbmodels.Merchant, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var merchant dbmodels.Merchant
	var err error

	err = db.Order("code desc limit 1").Find(&merchant).Error
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

func GetListMerchantBySearch(name string) ([]dbmodels.Merchant, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var merchant []dbmodels.Merchant
	var err error

	err = db.Where("name ilike ? ", "%"+name+"%").Find(&merchant).Error
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

// merchant check supplier
func GetOrderBySupplierAndMerchant(supplier string, merchant string) []dbmodels.Order {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var order []dbmodels.Order
	var err error

	err = db.Where("supplier_code ilike ? and merchant_code ilike ?", supplier, merchant).Find(&order).Error

	if err != nil {
		return order
	}

	return order
}

// merchant by id
func GetMerchantById(merchant_id int64) dbmodels.Merchant {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var merchant dbmodels.Merchant
	var err error

	err = db.Where("id = ?", merchant_id).Find(&merchant).Error

	if err != nil {
		return merchant
	}
	return merchant
}

func FindMerchantByPhone(phoneNumb string) dbmodels.Merchant {

	db := GetDbCon()
	db.Debug().LogMode(true)

	var merchant dbmodels.Merchant
	var err error

	err = db.Order("id desc limit 1").Where("phone_numb = ?", phoneNumb).Find(&merchant).Error

	if err == nil {
		return merchant
	}
	return dbmodels.Merchant{}
}
