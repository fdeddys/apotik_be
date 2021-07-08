package database

import (
	"distribution-system-be/constants"
	dbmodels "distribution-system-be/models/dbModels"
	"distribution-system-be/models/dto"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
)

// GetAllDataDetailPurchaseOrder ...
func GetAllDataDetailPurchaseOrder(purchaseOrderID int64) []dbmodels.PurchaseOrderDetail {

	db := GetDbCon()
	db.Debug().LogMode(true)

	var purchaseOrderDetails []dbmodels.PurchaseOrderDetail

	db.Preload("Product").Preload("UOM").Find(&purchaseOrderDetails, " po_id = ? and qty > 0 ", purchaseOrderID)

	return purchaseOrderDetails
}

func GetAllDataDetailPurchaseOrderByPoNo(purchaseOrderNo string) []dbmodels.PurchaseOrderDetail {

	db := GetDbCon()
	db.Debug().LogMode(true)

	var purchaseOrder dbmodels.PurchaseOrder

	var purchaseOrderDetails []dbmodels.PurchaseOrderDetail

	errPO := db.Find(&purchaseOrder, " po_no = ?  ", purchaseOrderNo).Error
	if errPO == gorm.ErrRecordNotFound {
		return purchaseOrderDetails
	}

	err := db.Find(&purchaseOrderDetails, " po_id = ?  and qty > 0 ", purchaseOrder.ID).Error
	if err == gorm.ErrRecordNotFound {
		return purchaseOrderDetails
	}

	return purchaseOrderDetails
}

// GetPurchaseOrderDetailPage ...
func GetPurchaseOrderDetailPage(param dto.FilterPurchaseOrderDetail, offset, limit int) ([]dbmodels.PurchaseOrderDetail, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var purchaseOrderDetails []dbmodels.PurchaseOrderDetail
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&purchaseOrderDetails).Error
		if err != nil {
			return purchaseOrderDetails, 0, err
		}
		return purchaseOrderDetails, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysPurchaseOrderDetails(db, offset, limit, &purchaseOrderDetails, param.PurchaseOrderID, errQuery)
	go AsyncQueryCountsPurchaseOrderDetails(db, &total, param.PurchaseOrderID, offset, limit, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return purchaseOrderDetails, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return purchaseOrderDetails, 0, resErrCount
	}
	return purchaseOrderDetails, total, nil
}

// AsyncQueryCountsPurchaseOrderDetails ...
func AsyncQueryCountsPurchaseOrderDetails(db *gorm.DB, total *int, purchaseOrderID int64, offset int, limit int, resChan chan error) {

	var err error

	err = db.Model(&dbmodels.PurchaseOrderDetail{}).Offset(offset).Where("po_id = ?", purchaseOrderID).Count(total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQuerysPurchaseOrderDetails ...
func AsyncQuerysPurchaseOrderDetails(db *gorm.DB, offset int, limit int, purchaseOrderDetails *[]dbmodels.PurchaseOrderDetail, purchaseOrderID int64, resChan chan error) {

	var err error

	err = db.Offset(offset).Limit(limit).Preload("Product").Preload("UOM").Find(&purchaseOrderDetails, "po_id = ? ", purchaseOrderID).Error
	if err != nil {
		fmt.Println("error --> ", err)
	}

	fmt.Println("order--> ", purchaseOrderDetails)

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

//SavePurchaseOrderDetail ...
func SavePurchaseOrderDetail(purchaseOrderDetail *dbmodels.PurchaseOrderDetail) (errCode string, errDesc string) {

	fmt.Println(" Update PurchaseOrder Detail  ------------------------------------------ ")

	db := GetDbCon()
	db.Debug().LogMode(true)

	if r := db.Save(&purchaseOrderDetail); r.Error != nil {
		errCode = "99"
		errDesc = r.Error.Error()
		return
	}

	errCode = constants.ERR_CODE_00
	errDesc = fmt.Sprintf("%v", purchaseOrderDetail.ID)
	return

}

// DeletePurchaseOrderDetailById ...
func DeletePurchaseOrderDetailById(id int64) (errCode string, errDesc string) {

	fmt.Println(" Delete PurchaseOrder Detail  ---- ", id)

	db := GetDbCon()
	db.Debug().LogMode(true)

	if r := db.Where("id = ? ", id).Delete(dbmodels.PurchaseOrderDetail{}); r.Error != nil {
		errCode = "99"
		errDesc = r.Error.Error()
		return
	}

	errCode = "00"
	errDesc = fmt.Sprintf("%v", id)
	return

}
