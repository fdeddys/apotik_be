package database

import (
	dbmodels "distribution-system-be/models/dbModels"
	"distribution-system-be/models/dto"
	"distribution-system-be/utils/util"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
)

// GetAllDataDetailReceive ...
func GetAllDataDetailReceive(receiveID int64) []dbmodels.ReceiveDetail {

	db := GetDbCon()
	db.Debug().LogMode(true)

	var receiveDetails []dbmodels.ReceiveDetail

	db.Preload("Product").Preload("UOM").Find(&receiveDetails, " receive_id = ? and qty > 0 ", receiveID)

	return receiveDetails
}

// GetReceiveDetailPage ...
func GetReceiveDetailPage(param dto.FilterReceiveDetail, offset, limit int) ([]dbmodels.ReceiveDetail, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var receiveDetails []dbmodels.ReceiveDetail
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&receiveDetails).Error
		if err != nil {
			return receiveDetails, 0, err
		}
		return receiveDetails, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysReceiveDetails(db, offset, limit, &receiveDetails, param.ReceiveID, errQuery)
	go AsyncQueryCountsReceiveDetails(db, &total, param.ReceiveID, offset, limit, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return receiveDetails, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return receiveDetails, 0, resErrCount
	}
	return receiveDetails, total, nil
}

// AsyncQueryCountsReceiveDetails ...
func AsyncQueryCountsReceiveDetails(db *gorm.DB, total *int, receiveID int64, offset int, limit int, resChan chan error) {

	var err error

	err = db.Model(&dbmodels.ReceiveDetail{}).Offset(offset).Where("receive_id = ?", receiveID).Count(total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQuerysReceiveDetails ...
func AsyncQuerysReceiveDetails(db *gorm.DB, offset int, limit int, receiveDetails *[]dbmodels.ReceiveDetail, receiveID int64, resChan chan error) {

	var err error

	err = db.Offset(offset).Limit(limit).Preload("Product.BigUom").Preload("Product.SmallUom").Preload("UOM").Find(&receiveDetails, "receive_id = ? ", receiveID).Error
	if err != nil {
		fmt.Println("error --> ", err)
	}

	fmt.Println("order--> ", receiveDetails)

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

//SaveReceiveDetail ...
func SaveReceiveDetail(receiveDetail *dbmodels.ReceiveDetail) (errCode string, errDesc string) {

	fmt.Println(" Update Receive Detail  ------------------------------------------ ")

	db := GetDbCon()
	db.Debug().LogMode(true)

	if r := db.Save(&receiveDetail); r.Error != nil {
		errCode = "99"
		errDesc = r.Error.Error()
		return
	}

	errCode = "00"
	errDesc = fmt.Sprintf("%v", receiveDetail.ID)
	return

}

//UpdateReceiveDetail ...
func UpdateReceiveDetail(idDetail, qty int64, price, disc1 float32, batchNo, ed string) (errCode string, errDesc string) {

	fmt.Println(" Update Receive Detail  ------------------------------------------ ")

	db := GetDbCon()
	db.Debug().LogMode(true)

	r := db.Model(dbmodels.ReceiveDetail{}).Where("id = ?", idDetail).Updates(
		dbmodels.ReceiveDetail{
			Qty:          qty,
			Price:        price,
			Disc1:        disc1,
			LastUpdate:   util.GetCurrDate(),
			LastUpdateBy: dto.CurrUser,
			BatchNo:      batchNo,
			Ed:           ed,
		})
	if r.Error != nil {
		errCode = "99"
		errDesc = r.Error.Error()
		return
	}

	errCode = "00"
	errDesc = fmt.Sprintf("%v", r.RowsAffected)
	return

}

// DeleteReceiveDetailById ...
func DeleteReceiveDetailById(id int64) (errCode string, errDesc string) {

	fmt.Println(" Delete Receive Detail  ---- ", id)

	db := GetDbCon()
	db.Debug().LogMode(true)

	if r := db.Where("id = ? ", id).Delete(dbmodels.ReceiveDetail{}); r.Error != nil {
		errCode = "99"
		errDesc = r.Error.Error()
		return
	}

	errCode = "00"
	errDesc = fmt.Sprintf("%v", id)
	return

}

// GetReceiveOrderDetailPage ...
func GetReceiveOrderDetailBatchExpiredPage(param dto.FilterBatchExpired, offset, limit int) ([]dbmodels.ReceiveDetail, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var ReceiveOrderDetails []dbmodels.ReceiveDetail
	var total int

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysReceiveOrderDetailsBatchExpired(db, offset, limit, &ReceiveOrderDetails, param, errQuery)
	go AsyncQueryCountsReceiveOrderDetailsBatchExpired(db, &total, param, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return ReceiveOrderDetails, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return ReceiveOrderDetails, 0, resErrCount
	}
	return ReceiveOrderDetails, total, nil
}

// AsyncQuerysReceiveOrders ...
func AsyncQuerysReceiveOrderDetailsBatchExpired(db *gorm.DB, offset int, limit int, ReceiveOrderDetails *[]dbmodels.ReceiveDetail, param dto.FilterBatchExpired, resChan chan error) {

	var err error
	byDate := true
	if param.ExpiredStart == "" || param.ExpiredEnd == "" {
		byDate = false
	}
	batch := "%" + param.Batch + "%"

	// err = db.Offset(offset).Limit(limit).Preload("Product.SmallUom").Preload("UOM").Find(&ReceiveOrderDetails, " ( batch_no ilike ? ) and ( (TO_DATE(ed,'YYYY-MM-DD') between TO_DATE(?, 'YYYY-MM-DD') and TO_DATE(?, 'YYYY-MM-DD') ) or ( not ?))  ", batch, param.ExpiredStart, param.ExpiredEnd, byDate).Error

	err = db.Offset(offset).Limit(limit).Preload("Product.SmallUom").Preload("UOM").Find(&ReceiveOrderDetails, " ( batch_no ilike ? ) and ( ( ed between ? and ? ) or ( not ?))  ", batch, param.ExpiredStart, param.ExpiredEnd, byDate).Error
	if err != nil {
		fmt.Println("error --> ", err)
	}

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQueryCountsReceiveOrders ...
func AsyncQueryCountsReceiveOrderDetailsBatchExpired(db *gorm.DB, total *int, param dto.FilterBatchExpired, resChan chan error) {

	var ReceiveOrderDetails []dbmodels.ReceiveDetail

	byDate := true
	if param.ExpiredStart == "" || param.ExpiredEnd == "" {
		byDate = false
	}
	batch := "%" + param.Batch + "%"

	var err error
	// err = db.Model(&ReceiveOrderDetails).Where(" ( batch_no ilike ? ) and ( (TO_DATE(ed,'YYYY-MM-DD') between TO_DATE(?,'YYYY-MM-DD' ) and TO_DATE(?, 'YYYY-MM-DD') ) or ( not ?))  ", batch, param.ExpiredStart, param.ExpiredEnd, byDate).Count(&*total).Error

	err = db.Model(&ReceiveOrderDetails).Where(" ( batch_no ilike ? ) and ( (ed  between ? and ? ) or ( not ?))  ", batch, param.ExpiredStart, param.ExpiredEnd, byDate).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}
