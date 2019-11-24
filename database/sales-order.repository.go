package database

import (
	"distribution-system-be/constants"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
)

//SaveSalesOrderNo ...
func SaveSalesOrderNo(order *dbmodels.SalesOrder) (errCode string, errDesc string) {

	fmt.Println(" Update Sales Order numb ------------------------------------------ ")
	var newOrder dbmodels.SalesOrder
	db := GetDbCon()
	db.Debug().LogMode(true)

	// r := db.Model(&newOrder).Where("id = ?", order.ID).Update(dbmodels.SalesOrder{OrderNo: order.OrderNo, StatusCode: "001", WarehouseCode: order.WarehouseCode, InternalStatus: 1, OrderDate: order.OrderDate})
	r := db.Model(&newOrder).Where("id = ?", order.ID).Update(dbmodels.SalesOrder{Status: 20, Note: order.Note})
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}

	// fmt.Println("Order [database]=> order id", order.OrderNo)

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// GetOrderByOrderNo ...
func GetOrderByOrderNo(orderNo string) (dbmodels.SalesOrder, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)
	order := dbmodels.SalesOrder{}

	err := db.Preload("Supplier").Preload("Merchant").Where(" sales_order_no = ?  ", orderNo).First(&order).Error

	return order, err

}

// GetSalesOrderByOrderId ...
func GetSalesOrderByOrderId(orderID int) (dbmodels.SalesOrder, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)
	order := dbmodels.SalesOrder{}

	err := db.Preload("Customer").Where(" id = ?  ", orderID).First(&order).Error

	return order, err

}

// GetOrderPage ...
func GetOrderPage(param dto.FilterOrder, offset, limit, internalStatus int) ([]dbmodels.SalesOrder, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var orders []dbmodels.SalesOrder
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&orders).Error
		if err != nil {
			return orders, 0, err
		}
		return orders, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysOrders(db, offset, limit, internalStatus, &orders, param, errQuery)
	go AsyncQueryCountsOrders(db, &total, internalStatus, &orders, param, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return orders, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return orders, 0, resErrCount
	}
	return orders, total, nil
}

func getParam(param dto.FilterOrder, status int) (merchantCode, orderNumber string, byStatus bool) {

	merchantCode = "%"

	orderNumber = param.OrderNumber
	if orderNumber == "" {
		orderNumber = "%"
	} else {
		orderNumber = "%" + param.OrderNumber + "%"
	}

	byStatus = true
	if status == -1 {
		byStatus = false
	}

	return
}

// AsyncQueryCountsOrders ...
func AsyncQueryCountsOrders(db *gorm.DB, total *int, status int, orders *[]dbmodels.SalesOrder, param dto.FilterOrder, resChan chan error) {

	merchantCode, orderNumber, byStatus := getParam(param, status)

	fmt.Println("ISI MERCHANT ", merchantCode, " orderNumber ", orderNumber, "  status ", status, " fill status ", byStatus)

	var err error
	if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
		err = db.Model(&orders).Where(" ( (status = ?) or ( not ?) ) AND  COALESCE(sales_order_no, '') ilike ? AND order_date between ? and ?  ", status, byStatus, orderNumber, param.StartDate, param.EndDate).Count(&*total).Error
	} else {
		err = db.Model(&orders).Where(" ( (status = ?) or ( not ?) ) AND COALESCE(sales_order_no,'') ilike ? ", status, byStatus, orderNumber).Count(&*total).Error
	}

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQuerysOrders ...
func AsyncQuerysOrders(db *gorm.DB, offset int, limit int, status int, orders *[]dbmodels.SalesOrder, param dto.FilterOrder, resChan chan error) {

	var err error

	merchantCode, orderNumber, byStatus := getParam(param, status)

	fmt.Println("ISI MERCHANT ", merchantCode, " order no ", orderNumber, "  status ", status, " fill status ", byStatus)

	fmt.Println("isi dari filter [", param, "] ")
	if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
		fmt.Println("isi dari filter [", param.StartDate, '-', param.EndDate, "] ")
		err = db.Preload("Customer").Preload("Salesman").Order("order_date DESC").Offset(offset).Limit(limit).Find(&orders, " ( ( status = ?) or ( not ?) ) AND COALESCE(sales_order_no, '') ilike ? AND order_date between ? and ?   ", status, byStatus, orderNumber, param.StartDate, param.EndDate).Error
	} else {
		fmt.Println("isi dari kosong ")
		err = db.Offset(offset).Limit(limit).Preload("Customer").Preload("Salesman").Find(&orders, " ( ( status = ?) or ( not ?) ) AND COALESCE(sales_order_no,'') ilike ?  ", status, byStatus, orderNumber).Error
		if err != nil {
			fmt.Println("error --> ", err)
		}
		fmt.Println("order--> ", orders)

	}

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}
