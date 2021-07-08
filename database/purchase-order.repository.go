package database

import (
	"distribution-system-be/constants"
	dbmodels "distribution-system-be/models/dbModels"
	"distribution-system-be/models/dto"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
)

//SavePurchaseOrder ...
func SavePurchaseOrder(purchaseOrder *dbmodels.PurchaseOrder) (errCode string, errDesc string, id int64, status int8) {

	db := GetDbCon()
	db.Debug().LogMode(true)

	r := db.Save(&purchaseOrder)
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		id = 0
		status = 0
		fmt.Println("Error update ", errDesc)
		return
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, purchaseOrder.ID, purchaseOrder.Status
}

func ApprovePurchaseOrder(purchaseOrder *dbmodels.PurchaseOrder) (errCode string, errDesc string) {

	fmt.Println(" Reject Purchase Order numb ------------------------------------------ ")
	db := GetDbCon()
	db.Debug().LogMode(true)
	totalPO := countTotalPO(purchaseOrder.PurchaserNo)
	tax := float32(0)

	if purchaseOrder.IsTax {
		tax = getTaxValue()
		tax = totalPO * (tax / 100)
	}
	r := db.Model(&dbmodels.PurchaseOrder{}).Where("id =?", purchaseOrder.ID).Update(dbmodels.PurchaseOrder{
		Status:     constants.STATUS_APPROVE,
		Total:      totalPO,
		Tax:        tax,
		GrandTotal: totalPO + tax,
	})
	if r.Error != nil {
		fmt.Println("err reject ", r.Error)
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

func getTaxValue() float32 {

	tax := float32(10)
	parameteTax, errCode, _, _ := GetParameterByNama(constants.PARAMETER_TAX_VALUE)
	if errCode != constants.ERR_CODE_00 {
		value, err := strconv.ParseFloat(parameteTax.Value, 32)
		if err != nil {
			tax = float32(value)
		}
	}

	return tax
}

func countTotalPO(poNo string) (total float32) {

	poDetails := GetAllDataDetailPurchaseOrderByPoNo(poNo)

	for _, poDetail := range poDetails {
		total = poDetail.Price * float32(poDetail.Qty)
		total -= (total * poDetail.Disc1 / 100)
		total -= (total * poDetail.Disc2 / 100)
	}
	return
}

// GetPurchaseOrderPage ...
func GetPurchaseOrderPage(param dto.FilterPurchaseOrder, offset, limit, internalStatus int) ([]dbmodels.PurchaseOrder, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var purchaseOrders []dbmodels.PurchaseOrder
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&purchaseOrders).Error
		if err != nil {
			return purchaseOrders, 0, err
		}
		return purchaseOrders, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysPurchaseOrders(db, offset, limit, internalStatus, &purchaseOrders, param, errQuery)
	go AsyncQueryCountsPurchaseOrders(db, &total, internalStatus, &purchaseOrders, param, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return purchaseOrders, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return purchaseOrders, 0, resErrCount
	}
	return purchaseOrders, total, nil
}

// AsyncQueryCountsPurchaseOrders ...
func AsyncQueryCountsPurchaseOrders(db *gorm.DB, total *int, status int, purchaseOrders *[]dbmodels.PurchaseOrder, param dto.FilterPurchaseOrder, resChan chan error) {

	purchaseOrderNumber, byStatus, bySupplierID := getParamPurchaseOrder(param, status)

	fmt.Println(" Rec Number ", purchaseOrderNumber, "  status ", status, " fill status ", byStatus)

	var err error
	if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
		err = db.Model(&purchaseOrders).Where(" ( (status = ?) or ( not ?) ) AND COALESCE(po_no, '') ilike ? AND po_date between ? and ?  AND ( ( supplier_id = ? ) or ( not ?) )  ", status, byStatus, purchaseOrderNumber, param.StartDate, param.EndDate, param.SupplierId, bySupplierID).Count(&*total).Error
	} else {
		err = db.Model(&purchaseOrders).Where(" ( (status = ?) or ( not ?) ) AND COALESCE(po_no,'') ilike ?  AND ( ( supplier_id = ? ) or ( not ?) ) ", status, byStatus, purchaseOrderNumber, param.SupplierId, bySupplierID).Count(&*total).Error
	}

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQuerysPurchaseOrders ...
func AsyncQuerysPurchaseOrders(db *gorm.DB, offset int, limit int, status int, purchaseOrders *[]dbmodels.PurchaseOrder, param dto.FilterPurchaseOrder, resChan chan error) {

	var err error

	purchaseOrderNumber, byStatus, bySupplierID := getParamPurchaseOrder(param, status)

	fmt.Println(" PurchaseOrder no ", purchaseOrderNumber, "  status ", status, " fill status ", byStatus)

	fmt.Println("isi dari filter [", param, "] ")
	if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
		fmt.Println("isi dari filter [", param.StartDate, '-', param.EndDate, "] ")
		err = db.Preload("Supplier").Order("po_date DESC").Offset(offset).Limit(limit).Find(&purchaseOrders, " ( ( status = ?) or ( not ?) ) AND COALESCE(po_no, '') ilike ? AND po_date between ? and ?  AND ( ( supplier_id = ? ) or ( not ?) )  ", status, byStatus, purchaseOrderNumber, param.StartDate, param.EndDate, param.SupplierId, bySupplierID).Error
	} else {
		fmt.Println("isi dari kosong ")
		err = db.Offset(offset).Limit(limit).Preload("Supplier").Find(&purchaseOrders, " ( ( status = ?) or ( not ?) ) AND COALESCE(po_no,'') ilike ?  AND ( ( supplier_id = ? ) or ( not ?) )  ", status, byStatus, purchaseOrderNumber, param.SupplierId, bySupplierID).Error
		if err != nil {
			fmt.Println("purchaseOrder --> ", err)
		}
		fmt.Println("purchaseOrder--> ", purchaseOrders)

	}

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

func getParamPurchaseOrder(param dto.FilterPurchaseOrder, status int) (purchaseOrderNumber string, byStatus, bySupplierID bool) {

	purchaseOrderNumber = param.PurchaseOrderNumber
	if purchaseOrderNumber == "" {
		purchaseOrderNumber = "%"
	} else {
		purchaseOrderNumber = "%" + param.PurchaseOrderNumber + "%"
	}

	byStatus = true
	if status == -1 {
		byStatus = false
	}

	bySupplierID = true
	if param.SupplierId == 0 {
		bySupplierID = false
	}

	return
}

// GetPurchaseOrderByPurchaseOrderID ...
func GetPurchaseOrderByPurchaseOrderID(purchaseOrderID int64) (dbmodels.PurchaseOrder, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)
	purchaseOrder := dbmodels.PurchaseOrder{}

	err := db.Preload("Supplier").Where(" id = ?  ", purchaseOrderID).First(&purchaseOrder).Error

	return purchaseOrder, err

}

//RejectPurchaseOrder ...
func RejectPurchaseOrder(purchaseOrder *dbmodels.PurchaseOrder) (errCode string, errDesc string) {

	fmt.Println(" Reject PurchaseOrder numb ------------------------------------------ ")
	db := GetDbCon()
	db.Debug().LogMode(true)

	r := db.Model(&dbmodels.PurchaseOrder{}).Where("id =?", purchaseOrder.ID).Update(dbmodels.PurchaseOrder{Status: 30})
	if r.Error != nil {
		fmt.Println("err reject ", r.Error)
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

func UpdatePoPaid(poNo string) (errCode string, errDesc string) {

	fmt.Println(" Update PurchaseOrder numb ------------------------------------------ ")
	db := GetDbCon()
	db.Debug().LogMode(true)

	r := db.Model(&dbmodels.PurchaseOrder{}).Where("po_no =?", poNo).Update(dbmodels.PurchaseOrder{Status: 40})
	if r.Error != nil {
		fmt.Println("err PO Paid ", r.Error)
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}
