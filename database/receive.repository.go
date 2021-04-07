package database

import (
	"distribution-system-be/constants"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

//SaveReceive ...
func SaveReceive(receive *dbmodels.Receive) (errCode string, errDesc string, id int64, status int8) {

	db := GetDbCon()
	db.Debug().LogMode(true)

	r := db.Save(&receive)
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		id = 0
		status = 0
		fmt.Println("Error update ", errDesc)
		return
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, receive.ID, receive.Status
}

// SaveReceiveApprove ...
func SaveReceiveApprove(receive *dbmodels.Receive) (errCode string, errDesc string) {

	fmt.Println(" Approve Receiving ------------------------------------------ ")
	db := GetDbCon()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// update stock
	// update history stock
	// hitung ulang
	var total float32
	var grandTotal float32
	total = 0
	grandTotal = 0
	receiveDetails := GetAllDataDetailReceive(receive.ID)
	for idx, receiveDetail := range receiveDetails {
		fmt.Println("idx -> ", idx)

		product, errCodeProd, errDescProd := FindProductByID(receiveDetail.ProductID)
		if errCodeProd != constants.ERR_CODE_00 {
			tx.Rollback()
			return errCodeProd, errDescProd
		}

		checkStock, _, _ := GetStockByProductAndWarehouse(product.ID, receive.WarehouseID)
		// curQty := checkStock.Qty
		curQty := checkStock.Qty
		updateQty := curQty + receiveDetail.Qty
		newHpp := reCalculateHpp(checkStock.Hpp, checkStock.Qty, receiveDetail.Price, receiveDetail.Qty)

		var historyStock dbmodels.HistoryStock
		historyStock.Code = product.Code
		historyStock.Description = "Receive"
		historyStock.Hpp = checkStock.Hpp
		historyStock.Name = product.Name
		historyStock.Price = receiveDetail.Price
		historyStock.ReffNo = receive.ReceiveNo
		historyStock.TransDate = receive.ReceiveDate
		historyStock.Debet = receiveDetail.Qty
		historyStock.Kredit = 0
		historyStock.Saldo = updateQty
		historyStock.LastUpdate = time.Now()
		historyStock.LastUpdateBy = dto.CurrUser

		UpdateStockAndHppProductByID(receiveDetail.ProductID, receive.WarehouseID, updateQty, newHpp)
		SaveHistory(historyStock)
		total = total + (receiveDetail.Price * float32(receiveDetail.Qty))
	}

	db.Debug().LogMode(true)
	// r := db.Model(&newOrder).Where("id = ?", order.ID).Update(dbmodels.SalesOrder{OrderNo: order.OrderNo, StatusCode: "001", WarehouseCode: order.WarehouseCode, InternalStatus: 1, OrderDate: order.OrderDate})

	if receive.Tax != 0 {
		grandTotal = total * 1.1
	}
	receive.GrandTotal = grandTotal
	receive.Total = total
	receive.LastUpdateBy = dto.CurrUser
	receive.LastUpdate = time.Now()
	receive.Status = 20
	r := db.Save(&receive)
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}

	// fmt.Println("Order [database]=> order id", order.OrderNo)

	tx.Commit()
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

func reCalculateHpp(hpp1 float32, qty1 int64, price2 float32, qty2 int64) float32 {

	totalRp := (hpp1 * float32(qty1)) + (price2 * float32(qty2))
	totalQty := qty1 + qty2
	return (totalRp / float32(totalQty))
}

// GetReceivePage ...
func GetReceivePage(param dto.FilterReceive, offset, limit, internalStatus int) ([]dbmodels.Receive, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var receives []dbmodels.Receive
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&receives).Error
		if err != nil {
			return receives, 0, err
		}
		return receives, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysReceives(db, offset, limit, internalStatus, &receives, param, errQuery)
	go AsyncQueryCountsReceives(db, &total, internalStatus, &receives, param, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return receives, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return receives, 0, resErrCount
	}
	return receives, total, nil
}

// AsyncQueryCountsReceives ...
func AsyncQueryCountsReceives(db *gorm.DB, total *int, status int, orders *[]dbmodels.Receive, param dto.FilterReceive, resChan chan error) {

	receiveNumber, byStatus := getParamReceive(param, status)

	fmt.Println(" Rec Number ", receiveNumber, "  status ", status, " fill status ", byStatus)

	var err error
	if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
		err = db.Model(&orders).Where(" ( (status = ?) or ( not ?) ) AND  COALESCE(receive_no, '') ilike ? AND order_date between ? and ?  ", status, byStatus, receiveNumber, param.StartDate, param.EndDate).Count(&*total).Error
	} else {
		err = db.Model(&orders).Where(" ( (status = ?) or ( not ?) ) AND COALESCE(receive_no,'') ilike ? ", status, byStatus, receiveNumber).Count(&*total).Error
	}

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQuerysReceives ...
func AsyncQuerysReceives(db *gorm.DB, offset int, limit int, status int, receives *[]dbmodels.Receive, param dto.FilterReceive, resChan chan error) {

	var err error

	receiveNumber, byStatus := getParamReceive(param, status)

	fmt.Println(" Receive no ", receiveNumber, "  status ", status, " fill status ", byStatus)

	fmt.Println("isi dari filter [", param, "] ")
	if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
		fmt.Println("isi dari filter [", param.StartDate, '-', param.EndDate, "] ")
		err = db.Preload("Supplier").Order("receive_date DESC").Offset(offset).Limit(limit).Find(&receives, " ( ( status = ?) or ( not ?) ) AND COALESCE(receive_no, '') ilike ? AND receive_date between ? and ?   ", status, byStatus, receiveNumber, param.StartDate, param.EndDate).Error
	} else {
		fmt.Println("isi dari kosong ")
		err = db.Offset(offset).Limit(limit).Preload("Supplier").Find(&receives, " ( ( status = ?) or ( not ?) ) AND COALESCE(receive_no,'') ilike ?  ", status, byStatus, receiveNumber).Error
		if err != nil {
			fmt.Println("receive --> ", err)
		}
		fmt.Println("receive--> ", receives)

	}

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

func getParamReceive(param dto.FilterReceive, status int) (receiveNumber string, byStatus bool) {

	receiveNumber = param.ReceiveNumber
	if receiveNumber == "" {
		receiveNumber = "%"
	} else {
		receiveNumber = "%" + param.ReceiveNumber + "%"
	}

	byStatus = true
	if status == -1 {
		byStatus = false
	}

	return
}

// GetReceiveByReceiveID ...
func GetReceiveByReceiveID(receiveID int64) (dbmodels.Receive, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)
	receive := dbmodels.Receive{}

	err := db.Preload("Supplier").Where(" id = ?  ", receiveID).First(&receive).Error

	return receive, err

}

//RejectReceive ...
func RejectReceive(receive *dbmodels.Receive) (errCode string, errDesc string) {

	fmt.Println(" Reject Receive numb ------------------------------------------ ")
	db := GetDbCon()
	db.Debug().LogMode(true)

	r := db.Model(&dbmodels.Receive{}).Where("id =?", receive.ID).Update(dbmodels.Receive{Status: 30})
	if r.Error != nil {
		fmt.Println("err reject ", r.Error)
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}
