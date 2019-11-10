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
	// "strconv"
	"oasis-be/constants"
)


func GetSupplierPricePaging(param dto.FilterSupplierPriceDto, offset int, limit int, supplier_id int64) ([]dbmodels.SupplierPrice, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplier []dbmodels.SupplierPrice
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

	go AsyncQuerysSupplierPrice(db, offset, limit, &supplier, param, supplier_id, errQuery)
	go AsyncQueryCountsSupplierPrice(db, &total, param, &supplier_id, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return supplier, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("err-->", resErrCount)
		return supplier, 0, resErrCount
	}
	return supplier, total, nil
}


// AsyncQuerysCountSupplierPrice ...
func AsyncQueryCountsSupplierPrice(db *gorm.DB, total *int, param dto.FilterSupplierPriceDto, supplier_id *int64, resChan chan error) {
	var searchCode = "%"
	if strings.TrimSpace(param.ProductCode) != "" {
		searchCode = param.ProductCode + searchCode
	}

	err := db.Preload("Product").Preload("Lookup", "lookup_group = ?", "SELL_MARGIN").Model(&dbmodels.SupplierPrice{}).Where("product_code ilike ? and supplier_id = ? ", searchCode, *supplier_id).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


// AsyncQuerysSupplierPrice ...
func AsyncQuerysSupplierPrice(db *gorm.DB, offset int, limit int, supplier *[]dbmodels.SupplierPrice, param dto.FilterSupplierPriceDto, supplier_id int64, resChan chan error) {

	var searchCode = "%"
	if strings.TrimSpace(param.ProductCode) != "" {
		searchCode = param.ProductCode + searchCode
	}

	err := db.Preload("Product").Preload("Lookup", "lookup_group = ?", "SELL_MARGIN").Order("id ASC").Offset(offset).Limit(limit).Find(&supplier, "product_code ilike ? and supplier_id = ?", searchCode, supplier_id).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


// Repository Save
func SaveSupplierPrice(supplier *dbmodels.SupplierPrice) models.Response {
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


// Repository Update
func UpdateSupplierPrice(supplier *dbmodels.SupplierPrice) models.Response {
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

// Get Supplier Product By 
func GetListProductById(supplier_id int64)([]dbmodels.SupplierPrice, error){
	db := GetDbCon()
	db.Debug().LogMode(true)

	var product []dbmodels.SupplierPrice
	var err error

	err = db.Preload("Product").Preload("Lookup", "lookup_group = ?", "SELL_MARGIN").Where("supplier_id = ?", supplier_id).Find(&product).Error
	if err != nil {
		return product, err
	}

	for i := 0; i < len(product); i++ {
		isExist := CheckImage(product[i].ProductCode, "product")
		if isExist {
			product[i].Product.IMG1 = GetImage(product[i].ProductCode, "product")
		} else {
			product[i].Product.IMG1 = GetImage("no_image", "product")
		}
	}

	return product, nil
}