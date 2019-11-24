package database

import (
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
)

// GetOrderDetailPage ...
func GetOrderDetailPage(param dto.FilterOrderDetail, offset, limit int) ([]dbmodels.SalesOrderDetail, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var orderDetails []dbmodels.SalesOrderDetail
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&orderDetails).Error
		if err != nil {
			return orderDetails, 0, err
		}
		return orderDetails, 0, nil
	}

	// order, err := GetOrderByOrderNo(param.OrderNo)
	// if err != nil {
	// 	return orderDetails, 0, err
	// }

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysOrderDetails(db, offset, limit, &orderDetails, param.OrderID, errQuery)
	go AsyncQueryCountsOrderDetails(db, &total, param.OrderID, offset, limit, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return orderDetails, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return orderDetails, 0, resErrCount
	}
	return orderDetails, total, nil
}

// AsyncQueryCountsOrderDetails ...
func AsyncQueryCountsOrderDetails(db *gorm.DB, total *int, orderID int64, offset int, limit int, resChan chan error) {

	var err error
	// if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
	// err = db.Model(&dbmodels.OrderDetail{}).Joins().Offset(offset).Where("order_id = ? AND  order_date between ? and ? ", orderId, param.StartDate, param.EndDate).Count(&*total).Error
	// err = db.Model(&dbmodels.OrderDetail{}).Joins("left join 'order' on 'order'.id = order_detail.id").Offset(offset).Where("order_no = ? AND  order_date between ? and ? ", param.OrderNo, param.StartDate, param.EndDate).Count(&*total).Error
	// } else {
	err = db.Model(&dbmodels.SalesOrderDetail{}).Offset(offset).Where("sales_order_id = ?", orderID).Count(total).Error
	// err = db.Model(&dbmodels.OrderDetail{}).Joins("left join 'order' on 'order'.id = order_detail.id").Offset(offset).Where("order_no = ?", orderID).Count(total).Error
	// }

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// GetAllDataDetail ...
func GetAllDataDetail(orderID int64) []dbmodels.SalesOrderDetail {

	db := GetDbCon()
	db.Debug().LogMode(true)

	var orderDetails []dbmodels.SalesOrderDetail

	db.Preload("Product").Preload("Lookup", "lookup_group = ?", "UOM").Find(&orderDetails, " order_id = ? and qty_receive > 0 ", orderID)

	return orderDetails
}

// AsyncQuerysOrderDetails ...
func AsyncQuerysOrderDetails(db *gorm.DB, offset int, limit int, orderDetails *[]dbmodels.SalesOrderDetail, orderID int64, resChan chan error) {

	var err error

	// var merchant dbmodels.Merchant
	// fmt.Println("isi dari filter [", param, "] ")
	// if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
	// fmt.Println("isi dari filter [", param.StartDate, '-', param.EndDate, "] ")
	// err = db.Order("name ASC").Offset(offset).Limit(limit).Find(&supplier, "name ilike ?", searchName).Error
	// err = db.Preload("Product").Order("product_code DESC").Offset(offset).Limit(limit).Find(&orderDetails, " order_id = ? AND  order_date between ? and ? ", param.OrderNo, param.StartDate, param.EndDate).Error
	// } else {
	// err = db.Order("name ASC").Offset(offset).Limit(limit).Find(&supplier, "name ilike ?", searchName).Error
	// fmt.Println("isi dari kosong ")

	err = db.Offset(offset).Limit(limit).Preload("Product").Preload("UOM").Find(&orderDetails, " sales_order_id = ? ", orderID).Error
	if err != nil {
		fmt.Println("error --> ", err)
	}

	fmt.Println("order--> ", orderDetails)

	// err = db.Order("order_date DESC").Offset(offset).Limit(limit).Find(&orders).Error
	// }

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}
