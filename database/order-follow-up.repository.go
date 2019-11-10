package database

import (
	"fmt"
	"log"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
)

func GetFollowUpOrder(param dto.FilterOrder, offset, limit, internalStatus int) ([]dbmodels.FollowOrder, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var followUpOrder []dbmodels.FollowOrder
	var total int
	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&followUpOrder).Error
		if err != nil {
			return followUpOrder, 0, err
		}
		return followUpOrder, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysFollow(db, offset, limit, internalStatus, &followUpOrder, param, errQuery)
	go AsyncQueryCountsFollow(db, &total, internalStatus, &followUpOrder, param, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return followUpOrder, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("err-->", resErrCount)
		return followUpOrder, 0, resErrCount
	}

	return followUpOrder, total, nil
}

func AsyncQueryCountsFollow(db *gorm.DB, total *int, internalStatus int, orders *[]dbmodels.FollowOrder, param dto.FilterOrder, resChan chan error) {

	salesNo := param.SalesNo
	if salesNo == "" {
		salesNo = "%"
	} else {
		salesNo = "%" + param.SalesNo + "%"
	}

	var err error

	orderNumber := param.OrderNumber
	if orderNumber == "" {
		orderNumber = "%"
	} else {
		orderNumber = "%" + param.OrderNumber + "%"
	}

	if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
		err = db.Model(&orders).Where("internal_status = ? AND  COALESCE(order_no, '') ilike ? AND order_date between ? and ? and COALESCE(salesman_no, '') ILIKE ? ", internalStatus, orderNumber, param.StartDate, param.EndDate, salesNo).Count(&*total).Error
	} else {
		err = db.Model(&orders).Where("internal_status = ? AND COALESCE(order_no,'') ilike ? and COALESCE(salesman_no, '') ILIKE ? ", internalStatus, orderNumber, salesNo).Count(&*total).Error
	}

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

func AsyncQuerysFollow(db *gorm.DB, offset int, limit int, internalStatus int, orders *[]dbmodels.FollowOrder, param dto.FilterOrder, resChan chan error) {

	var err error

	salesNo := param.SalesNo
	if salesNo == "" {
		salesNo = "%"
	} else {
		salesNo = "%" + param.SalesNo + "%"
	}

	orderNumber := param.OrderNumber
	if orderNumber == "" {
		orderNumber = "%"
	} else {
		orderNumber = "%" + param.OrderNumber + "%"
	}
	fmt.Println("Isi sales no [", salesNo, "] ")
	fmt.Println("isi dari filter [", param, "] ")
	if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
		fmt.Println("isi dari filter [", param.StartDate, '-', param.EndDate, "] ")
		err = db.Preload("Merchant").Preload("Supplier").Preload("OrderDetail").Preload("OrderStatus").Order("order_date DESC").Offset(offset).Limit(limit).Find(&orders, " internal_status = ? AND COALESCE(order_no, '') ilike ? AND order_date between ? and ? and COALESCE(salesman_no, '') ILIKE ?  ", internalStatus, orderNumber, param.StartDate, param.EndDate, salesNo).Error
	} else {
		fmt.Println("isi dari kosong ")
		err = db.Offset(offset).Limit(limit).Preload("Merchant").Preload("Supplier").Preload("OrderDetail").Preload("OrderStatus").Find(&orders, " internal_status = ? AND COALESCE(order_no,'') ilike ? and COALESCE(salesman_no, '') ILIKE ? ", internalStatus, orderNumber, salesNo).Error
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
