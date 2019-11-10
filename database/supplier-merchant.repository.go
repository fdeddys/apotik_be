package database

import (
	"fmt"
	"log"
	"oasis-be/constants"
	"oasis-be/models"
	"oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	_ "strconv"

	"strings"
	"sync"

	"github.com/jinzhu/gorm"
)

// GetSupplierMerchantPaging Data Supplier Merchant
func GetSupplierMerchantPaging(param dto.FilterSupplierMerchantDto, offset int, limit int, supplier_id int64) ([]dbmodels.SupplierMerchant, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplier []dbmodels.SupplierMerchant
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Preload("Merchant").Find(&supplier).Error
		fmt.Println("test")
		if err != nil {
			return supplier, 0, err
		}
		return supplier, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysSupplierMerchant(db, offset, limit, &supplier, param, supplier_id, errQuery)
	go AsyncQueryCountsSupplierMerchant(db, &total, param, &supplier_id, errCount)

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
	return supplier, total, nil
}

// AsyncQuerysCountSupplierMerchant ...
func AsyncQueryCountsSupplierMerchant(db *gorm.DB, total *int, param dto.FilterSupplierMerchantDto, supplier_id *int64, resChan chan error) {
	var searchCode = "%"
	if strings.TrimSpace(param.MerchantCode) != "" {
		searchCode = param.MerchantCode + searchCode
	}

	err := db.Preload("Merchant").Model(&dbmodels.SupplierMerchant{}).Where("merchant_code ilike ? and supplier_id = ? ", searchCode, *supplier_id).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQuerysSupplierMerchant ...
func AsyncQuerysSupplierMerchant(db *gorm.DB, offset int, limit int, supplier *[]dbmodels.SupplierMerchant, param dto.FilterSupplierMerchantDto, supplier_id int64, resChan chan error) {

	var searchCode = "%"
	if strings.TrimSpace(param.MerchantCode) != "" {
		searchCode = param.MerchantCode + searchCode
	}

	err := db.Preload("Merchant").Order("merchant_code ASC").Offset(offset).Limit(limit).Find(&supplier, "merchant_code ilike ? and supplier_id = ?", searchCode, supplier_id).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// Repository Save
func SaveSupplierMerchant(supplier *dbmodels.SupplierMerchant) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if r := db.Save(&supplier); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}

// Repository Update
func UpdateSupplierMerchant(supplier *dbmodels.SupplierMerchant) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if r := db.Save(&supplier); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}

func GetListMerchantBySupplier(supplier_id int64, merchant_id int64) []dbmodels.SupplierMerchant {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierMerchant []dbmodels.SupplierMerchant

	err := db.Table("supplier_merchant_list s").Select("s.*").Joins("left join merchant m on m.code = s.merchant_code").Where("s.supplier_id = ? and m.id = ?",
		supplier_id, merchant_id).Find(&supplierMerchant).Error
	if err != nil {
		return supplierMerchant
	}
	return supplierMerchant
}

// UpdateFirstOrder ...
func UpdateFirstOrder(supplierID int64, merchantCode string) (errCode, errDesc string) {

	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierMerchant dbmodels.SupplierMerchant

	r := db.Model(&supplierMerchant).Where("supplier_id = ? and merchant_code = ?", supplierID, merchantCode).Update(dbmodels.SupplierMerchant{FirstOrder: 1})
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG

}

// func FirstOrderMerchantBySupplier(supplier_id int, merchant_id int) (first_order int){
// 	db := GetDbCon()
// 	db.Debug().LogMode(true)

// 	var supplierMerchant dbmodels.SupplierMerchant

// 	err := db.Table("supplier_merchant_list s").Select("s.first_order").Joins("left join merchant m on m.code = s.merchant_code").Where("s.supplier_id = ? and m.id = ?",
// 	strconv.Itoa(supplier_id), strconv.Itoa(merchant_id)).Find(&supplierMerchant).Error
// 	if err != nil {
// 		return first_order
// 	}

// 	first_order = supplierMerchant.FirstOrder
// 	return first_order
// }
