package database

import (
	"oasis-be/models/dbModels"
	_ "strconv"
	dto "oasis-be/models/dto"
	"sync"
	"github.com/jinzhu/gorm"
	"strings"
	"oasis-be/models"
	// "strconv"
	"oasis-be/constants"
	// "fmt"
)

func GetNooDocBySupplierId(supplier_id int64)([]dbmodels.Lookup){
	db := GetDbCon()
	db.Debug().LogMode(true)

	var noodoc []dbmodels.Lookup

	err := db.Table("supplier_noo_doc d").Select("l.name").Joins("left join lookup l on l.code = d.lookup_code").Where("l.lookup_group = ? and d.supplier_id = ?", "NOO_DOCUMENT", supplier_id).Find(&noodoc)
	if err != nil {
		return noodoc
	}
	return noodoc
}


// Repository Save
func SaveSupplierNooDoc(supplierNooDoc *dbmodels.SupplierNooDoc) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if  r := db.Save(&supplierNooDoc); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}


// Repository Update
func UpdateSupplierNooDoc(supplierNooDoc *dbmodels.SupplierNooDoc) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if  r := db.Save(&supplierNooDoc); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}


func GetNooDocBySuppAndMerch(supplier_id int64, merchant_code string)([]dbmodels.SupplierNooDoc, error){
	db := GetDbCon()
	db.Debug().LogMode(true)

	var noodoc []dbmodels.SupplierNooDoc
	var err error

	err = db.Where("merchant_code ilike ? and supplier_id = ?", "%"+merchant_code+"%", supplier_id).Find(&noodoc).Error
	if err != nil {
		return noodoc, err
	}
	return noodoc, nil
}


func CheckNooMerchantBySupplier(supplier_id int64, merchant_id int64)([]dbmodels.SupplierNooDoc) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var noodoc []dbmodels.SupplierNooDoc
	
	err := db.Table("supplier_noo_doc s").Select("s.*").Joins("left join merchant m on m.code = s.merchant_code").Where("s.supplier_id = ? and m.id = ? ", supplier_id, merchant_id).Find(&noodoc).Error

	if err != nil {
		return noodoc
	}
	return noodoc
}


func GetNOODocById(id int64)([]dbmodels.SupplierNooDoc){
	db := GetDbCon()
	db.Debug().LogMode(true)

	var noodoc []dbmodels.SupplierNooDoc

	err := db.Where("id = ? ", id).Find(&noodoc)
	if err != nil {
		return noodoc
	}
	return noodoc
}


/* ---------------------------- List Supplier By Approve --------------------------------------------- */

func GetSupplierNooDocApprovePaging(param dto.FilterSupplierNooDocDto, offset int, limit int, supplier_id int64) ([]dbmodels.SupplierNooDoc, int, error) {
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

	go AsyncQuerysSupplierNooDocApprove(db, offset, limit, &supplier, param, supplier_id, errQuery)
	go AsyncQueryCountsSupplierNooDocApprove(db, &total, param, &supplier_id, errCount)
	
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

func AsyncQueryCountsSupplierNooDocApprove(db *gorm.DB, total *int, param dto.FilterSupplierNooDocDto, supplier_id *int64, resChan chan error) {
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Model(&dbmodels.SupplierNooDoc{}).Preload("Supplier").Preload("MerchantPict").Preload("Merchant").Where("supplier_id = ? and approval_status = '1' ", *supplier_id).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


func AsyncQuerysSupplierNooDocApprove(db *gorm.DB, offset int, limit int, supplier *[]dbmodels.SupplierNooDoc, param dto.FilterSupplierNooDocDto, supplier_id int64, resChan chan error) {
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Order("id ASC").Offset(offset).Preload("Supplier").Preload("MerchantPict").Limit(limit).Preload("Merchant").Where("supplier_id = ? and approval_status = '1'", supplier_id).Find(&supplier).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
	
}

func GetStatusApproveByMerchantAndSupplier(suplier_id int64, merchant_code string)([]dbmodels.SupplierNooDoc, string, string, error){
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplierNooDoc []dbmodels.SupplierNooDoc

	err := db.Preload("Supplier").Preload("MerchantPict").Preload("Merchant").Where("merchant_code = ? and supplier_id = ?", merchant_code, suplier_id).Find(&supplierNooDoc).Error
	if err != nil {
		return nil, "02", "Error query data to DB", err
	}
	return supplierNooDoc, "00", "success", nil
}
/* ---------------------------- List Supplier By Approve --------------------------------------------- */






