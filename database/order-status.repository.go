package database

import (
	"fmt"
	"log"
	"oasis-be/constants"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"sync"

	"github.com/jinzhu/gorm"
)

//SaveOrderStatus ...
func SaveOrderStatus(orderStatus *dbmodels.OrderStatus) (errCode string, errDesc string) {

	fmt.Println(" Save Sales status ------------------------------------------ ")
	db := GetDbCon()
	db.Debug().LogMode(true)

	r := db.Create(&orderStatus)
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG

}

// GetOrderStatusPage ...
func GetOrderStatusPage(param dto.FilterOrderDetail, offset, limit int) ([]dbmodels.OrderStatus, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var orderStatuses []dbmodels.OrderStatus
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&orderStatuses).Error
		if err != nil {
			return orderStatuses, 0, err
		}
		return orderStatuses, 0, nil
	}

	order, err := GetOrderByOrderNo(param.OrderNo)
	if err != nil {
		return orderStatuses, 0, err
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysOrderStatuses(db, offset, limit, &orderStatuses, order.OrderNo, errQuery)
	go AsyncQueryCountsOrderStatuses(db, &total, order.OrderNo, offset, limit, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return orderStatuses, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return orderStatuses, 0, resErrCount
	}
	return orderStatuses, total, nil
}

// AsyncQueryCountsOrderStatuses ...
func AsyncQueryCountsOrderStatuses(db *gorm.DB, total *int, orderNo string, offset int, limit int, resChan chan error) {

	// var err error

	err := db.Model(&dbmodels.OrderStatus{}).Offset(offset).Where("order_no = ?", orderNo).Count(total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQuerysOrderStatuses ...
func AsyncQuerysOrderStatuses(db *gorm.DB, offset int, limit int, orderStatuses *[]dbmodels.OrderStatus, orderNo string, resChan chan error) {

	// var err error

	err := db.Order("id asc").Offset(offset).Limit(limit).Find(&orderStatuses, "order_no = ?", orderNo).Error
	if err != nil {
		fmt.Println("error --> ", err)
	}

	fmt.Println("order--> ", orderStatuses)

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}
