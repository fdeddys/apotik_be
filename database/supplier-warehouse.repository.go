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
	"fmt"
)

// Get Data Supplier Warehouse
func GetSupplierWarehousePaging(param dto.FilterSupplierWarehouseDto, offset int, limit int, supplier_id int64) ([]dbmodels.SupplierWarehouse, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierWarehouse []dbmodels.SupplierWarehouse
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&supplierWarehouse).Error
		if err != nil {
			return supplierWarehouse, 0, err
		}
		return supplierWarehouse, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysSupplierWarehouse(db, offset, limit, &supplierWarehouse, param, supplier_id, errQuery)
	go AsyncQueryCountsSupplierWarehouse(db, &total, param, &supplier_id, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return supplierWarehouse, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return supplierWarehouse, 0, resErrCount
	}
	return supplierWarehouse, total, nil
}


// AsyncQuerysCountSupplierWarehouse ...
func AsyncQueryCountsSupplierWarehouse(db *gorm.DB, total *int, param dto.FilterSupplierWarehouseDto, supplier_id *int64, resChan chan error) {
	var searchCode = "%"
	if strings.TrimSpace(param.Code) != "" {
		searchCode = param.Code + searchCode
	}

	err := db.Model(&dbmodels.SupplierWarehouse{}).Where("code ilike ? and supplier_id = ? ", searchCode, *supplier_id).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


// AsyncQuerysSupplierWarehouse ...
func AsyncQuerysSupplierWarehouse(db *gorm.DB, offset int, limit int, supplierWarehouse *[]dbmodels.SupplierWarehouse, param dto.FilterSupplierWarehouseDto, supplier_id int64, resChan chan error) {

	var searchCode = "%"
	if strings.TrimSpace(param.Code) != "" {
		searchCode = param.Code + searchCode
	}

	err := db.Order("code ASC").Offset(offset).Limit(limit).Find(&supplierWarehouse, "code ilike ? and supplier_id = ?", searchCode, supplier_id).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


// Repository Save
func SaveSupplierWarehouse(supplierWarehouse *dbmodels.SupplierWarehouse) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if  r := db.Save(&supplierWarehouse); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"

	return res
}


// Repository Update
func UpdateSupplierWarehouse(supplierWarehouse *dbmodels.SupplierWarehouse) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if  r := db.Save(&supplierWarehouse); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"

	return res
}

func GetLastSupplierWarehouse()(dbmodels.SupplierWarehouse, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierWarehouse dbmodels.SupplierWarehouse
	var err error

	err = db.Order("code desc limit 1").Find(&supplierWarehouse).Error
	if err != nil {
		return supplierWarehouse, err
	}
	return supplierWarehouse, nil
}

func GetWarehouseBySupplierId(supplier_code string)([]dbmodels.SupplierWarehouse, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierWarehouse []dbmodels.SupplierWarehouse
	var err error

	err = db.Joins("left join supplier on supplier.id = supplier_warehouset_list.supplier_id").Where("supplier.code = ? ", supplier_code).Find(&supplierWarehouse).Error
	fmt.Println(err)
	if err != nil {
		return supplierWarehouse, err
	}
	return supplierWarehouse, nil
}
