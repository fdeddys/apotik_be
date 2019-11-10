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

func GetSupplierNooChecklistPaging(param dto.FilterSupplierNooChecklistDto, offset int, limit int, supplier_id int64) ([]dbmodels.SupplierNooChecklist, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplier []dbmodels.SupplierNooChecklist
	var respSupplierNooChecklist []dbmodels.ResponseSupplierNooChecklist
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

	go AsyncQuerysSupplierNooChecklist(db, offset, limit, &respSupplierNooChecklist, param, supplier_id, errQuery)
	go AsyncQueryCountsSupplierNooChecklist(db, &total, param, &supplier_id, errCount)
	

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

	for i:=0; i<len(respSupplierNooChecklist);i++{
		supplier = append(supplier, dbmodels.SupplierNooChecklist{ID: respSupplierNooChecklist[i].ID, 
			Lookup:dbmodels.Lookup{Name:respSupplierNooChecklist[i].Name, ID: respSupplierNooChecklist[i].LookupID, Status: 
			respSupplierNooChecklist[i].Status, LookupGroup: respSupplierNooChecklist[i].LookupGroup, Code: respSupplierNooChecklist[i].Code}, 
			ChecklistType: respSupplierNooChecklist[i].ChecklistType,
			SupplierID: respSupplierNooChecklist[i].SupplierID, LookupCode: respSupplierNooChecklist[i].Code, LastUpdate:
			respSupplierNooChecklist[i].LastUpdate, LastUpdateBy: respSupplierNooChecklist[i].LastUpdateBy})
	}
	return supplier, total, nil
}


// AsyncQuerysCountSupplierPrice ...
func AsyncQueryCountsSupplierNooChecklist(db *gorm.DB, total *int, param dto.FilterSupplierNooChecklistDto, supplier_id *int64, resChan chan error) {
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Table("supplier_noo_check_list c").Select("c.id as id, l.name as name, c.supplier_id as supplier_id, " + 
	"l.code as lookup_code, c.supplier_id as supplier_id, l.code as code, c.is_mandatory as is_mandatory, " +
	"c.checklist_type as checklist_type, l.status as status, c.last_update_by as last_update_by, " + 
	"c.last_update as last_update, l.lookup_group as lookup_group, l.id as lookup_id").Joins("left join lookup l " + 
	"on l.code = c.lookup_code").Where("l.lookup_group in (CASE WHEN c.checklist_type = 'IMG' THEN 'NOO_DOCUMENT' ELSE 'NOO_CHECK_LIST' END) and supplier_id = ? ", *supplier_id).Count(&*total).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


func AsyncQuerysSupplierNooChecklist(db *gorm.DB, offset int, limit int, supplier *[]dbmodels.ResponseSupplierNooChecklist, param dto.FilterSupplierNooChecklistDto, supplier_id int64, resChan chan error){
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = param.Name + searchName
	}

	err := db.Table("supplier_noo_check_list c").Select("c.id as id, l.name as name, c.supplier_id as supplier_id, " + 
	"l.code as lookup_code, c.supplier_id as supplier_id, l.code as code, c.is_mandatory as is_mandatory, " +
	"c.checklist_type as checklist_type, l.status as status, c.last_update_by as last_update_by, " + 
	"c.last_update as last_update, l.lookup_group as lookup_group, l.id as lookup_id").Joins("left join lookup l " + 
	"on l.code = c.lookup_code").Where("l.lookup_group in (CASE WHEN c.checklist_type = 'IMG' THEN 'NOO_DOCUMENT' ELSE 'NOO_CHECK_LIST' END) and c.supplier_id = ?", 
	supplier_id).Find(&supplier).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}


// Repository Save
func SaveSupplierNooChecklist(supplierNooChecklist *dbmodels.SupplierNooChecklist) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if  r := db.Save(&supplierNooChecklist); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}


// Repository Update
func UpdateSupplierNooChecklist(supplierNooChecklist *dbmodels.SupplierNooChecklist) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if  r := db.Save(&supplierNooChecklist); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}


func GetNooChecklistBySupplierId(supplier_id int64)([]dbmodels.ResponseSupplierNooChecklist){
	db := GetDbCon()
	db.Debug().LogMode(true)

	var noochecklist []dbmodels.ResponseSupplierNooChecklist

	err := db.Table("supplier_noo_check_list c").Select("c.id as id, l.name as name, c.supplier_id as supplier_id, " + 
	"l.code as lookup_code, c.supplier_id as supplier_id, l.code as code, c.is_mandatory as is_mandatory, " +
	"c.checklist_type as checklist_type, l.status as status, c.last_update_by as last_update_by, " + 
	"c.last_update as last_update, l.lookup_group as lookup_group, l.id as lookup_id").Joins("left join lookup l " + 
	"on l.code = c.lookup_code").Where("l.lookup_group in (CASE WHEN c.checklist_type = 'IMG' THEN 'NOO_DOCUMENT' ELSE 'NOO_CHECK_LIST' END) and c.supplier_id = ?", 
	supplier_id).Find(&noochecklist)

	if err != nil {
		return noochecklist
	}
	return noochecklist
}