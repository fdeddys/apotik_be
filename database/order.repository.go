package database

import (
	"fmt"
	"log"
	"oasis-be/constants"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
)

// FindOrderReadyToPay ...
// Fungsi validasi Order ready to pay
// isAutodebet -> pengecekan apakah pernah di lakukan audodebet
//    jika belum pernah boleh autodebet cumen 1x
//    jika sudah pernah boleh manual payment
func FindOrderReadyToPay(orderNo string) (errCode string, errDesc string, orderID int64) {

	var order dbmodels.Order
	db := GetDbCon()
	db.Debug().LogMode(true)

	r := db.Where("order_no = ? and internal_status = ? ", orderNo, 2).First(&order)
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = "Order ID not found or internal status not PAYMENT or Payment already autodebet ! "
		orderID = 0
		fmt.Println("Error update ", errDesc)
		return
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, order.ID
}

// UpdateInternalStatus ...
// update untuk payment Autodebet
func UpdateInternalStatus(orderID int64, status string) (errCode string, errDesc string) {
	fmt.Println("update internal ", orderID, status)

	var order dbmodels.Order
	db := GetDbCon()
	db.Debug().LogMode(true)

	internalStatus := 2
	if status == "1" {
		internalStatus = 3
	}
	r := db.Model(&order).Where("id = ?", orderID).Update(dbmodels.Order{Autodebet: 1, InternalStatus: internalStatus})
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// InsertStatusOrder  ...
func InsertStatusOrder(orderID int8, status string) {
	fmt.Println("insert status => ", orderID, status)
}

//SaveSalesOrderNo ...
func SaveSalesOrderNo(order *dbmodels.Order) (errCode string, errDesc string) {

	fmt.Println(" Update Sales Order numb ------------------------------------------ ")
	var newOrder dbmodels.Order
	db := GetDbCon()
	db.Debug().LogMode(true)

	// r := db.Model(&newOrder).Where("id = ?", order.ID).Update(dbmodels.Order{OrderNo: order.OrderNo, StatusCode: "001", WarehouseCode: order.WarehouseCode, InternalStatus: 1, OrderDate: order.OrderDate})
	r := db.Model(&newOrder).Where("id = ?", order.ID).Update(dbmodels.Order{StatusCode: "001", WarehouseCode: order.WarehouseCode, InternalStatus: 1, Note: order.Note})
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
func GetOrderByOrderNo(orderNo string) (dbmodels.Order, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)
	order := dbmodels.Order{}

	err := db.Preload("Supplier").Preload("Merchant").Where(" order_no = ?  ", orderNo).First(&order).Error

	return order, err

}

// GetOrderPage ...
func GetOrderPage(param dto.FilterOrder, offset, limit, internalStatus int) ([]dbmodels.Order, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var orders []dbmodels.Order
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

func getParam(param dto.FilterOrder, internalStatus int) (merchantCode, salesNo, orderNumber string, byStatus bool) {
	phoneNumb := param.MerchantPhone
	// var merchantCode string

	merchantCode = "%"
	if phoneNumb != "" {
		merchant := FindMerchantByPhone(phoneNumb)

		fmt.Println("hasil search merchant ", merchant)
		if merchant.ID != 0 {
			merchantCode = merchant.Code
		}
	}

	salesNo = param.SalesNo
	if salesNo == "" {
		salesNo = "%"
	} else {
		salesNo = "%" + param.SalesNo + "%"
	}

	orderNumber = param.OrderNumber
	if orderNumber == "" {
		orderNumber = "%"
	} else {
		orderNumber = "%" + param.OrderNumber + "%"
	}

	byStatus = true
	if internalStatus == -1 {
		byStatus = false
	}

	return
}

// AsyncQueryCountsOrders ...
func AsyncQueryCountsOrders(db *gorm.DB, total *int, internalStatus int, orders *[]dbmodels.Order, param dto.FilterOrder, resChan chan error) {

	merchantCode, salesNo, orderNumber, byStatus := getParam(param, internalStatus)

	fmt.Println("ISI MERCHANT ", merchantCode, "sales no ", salesNo, "order no ", orderNumber, "internal status ", internalStatus, " fill status ", byStatus)

	var err error
	if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {

		// err = db.Model(&orders).Where("internal_status = ? AND  order_no ilike ? AND order_date between ? and ?", internalStatus, orderNumber, param.StartDate, param.EndDate).Count(&*total).Error
		err = db.Model(&orders).Where(" ( (internal_status = ?) or ( not ?) ) AND  COALESCE(order_no, '') ilike ? AND order_date between ? and ? and COALESCE(salesman_no, '') ILIKE ? ", internalStatus, byStatus, orderNumber, param.StartDate, param.EndDate, salesNo).Count(&*total).Error
	} else {
		// err = db.Model(&orders).Where("internal_status = ? AND order_no ilike ?", internalStatus, orderNumber).Count(&*total).Error
		err = db.Model(&orders).Where(" ( (internal_status = ?) or ( not ?) ) AND COALESCE(order_no,'') ilike ? and COALESCE(salesman_no, '') ILIKE ? ", internalStatus, byStatus, orderNumber, salesNo).Count(&*total).Error
	}

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQuerysOrders ...
func AsyncQuerysOrders(db *gorm.DB, offset int, limit int, internalStatus int, orders *[]dbmodels.Order, param dto.FilterOrder, resChan chan error) {

	var err error

	merchantCode, salesNo, orderNumber, byStatus := getParam(param, internalStatus)

	fmt.Println("ISI MERCHANT ", merchantCode, "sales no ", salesNo, "order no ", orderNumber, "internal status ", internalStatus, " fill status ", byStatus)

	// salesNo := param.SalesNo
	// if salesNo == "" {
	// 	salesNo = "%"
	// } else {
	// 	salesNo = "%" + param.SalesNo + "%"
	// }

	// orderNumber := param.OrderNumber
	// if orderNumber == "" {
	// 	orderNumber = "%"
	// } else {
	// 	orderNumber = "%" + param.OrderNumber + "%"
	// }
	// fmt.Println("Isi sales no [", salesNo, "] ")
	// var merchant dbmodels.Merchant
	fmt.Println("isi dari filter [", param, "] ")
	if strings.TrimSpace(param.StartDate) != "" && strings.TrimSpace(param.EndDate) != "" {
		fmt.Println("isi dari filter [", param.StartDate, '-', param.EndDate, "] ")
		// err = db.Order("name ASC").Offset(offset).Limit(limit).Find(&supplier, "name ilike ?", searchName).Error
		// err = db.Preload("Merchant").Preload("Supplier").Order("order_date DESC").Offset(offset).Limit(limit).Find(&orders, " internal_status = ? AND order_no ilike ? AND order_date between ? and ? ", internalStatus, orderNumber, param.StartDate, param.EndDate).Error
		// err = db.Preload("Merchant").Preload("Supplier").Order("order_date DESC").Offset(offset).Limit(limit).Find(&orders, " internal_status = ? AND COALESCE(order_no, '') ilike ? AND order_date between ? and ? and COALESCE(salesman_no, '') ILIKE ?  ", internalStatus, orderNumber, param.StartDate, param.EndDate, salesNo).Error
		err = db.Preload("Merchant").Preload("Supplier").Order("order_date DESC").Offset(offset).Limit(limit).Find(&orders, " ( (internal_status = ?) or ( not ?) ) AND COALESCE(order_no, '') ilike ? AND order_date between ? and ? and COALESCE(salesman_no, '') ILIKE ?  ", internalStatus, byStatus, orderNumber, param.StartDate, param.EndDate, salesNo).Error
	} else {
		// err = db.Order("name ASC").Offset(offset).Limit(limit).Find(&supplier, "name ilike ?", searchName).Error
		fmt.Println("isi dari kosong ")

		// err = db.Offset(offset).Limit(limit).Preload("Merchant").Preload("Supplier").Find(&orders, " internal_status = ? AND order_no ilike ?", internalStatus, orderNumber).Error
		// err = db.Offset(offset).Limit(limit).Preload("Merchant").Preload("Supplier").Find(&orders, " internal_status = ? AND COALESCE(order_no,'') ilike ? and COALESCE(salesman_no, '') ILIKE ? ", internalStatus, orderNumber, salesNo).Error
		err = db.Offset(offset).Limit(limit).Preload("Merchant").Preload("Supplier").Find(&orders, " ( (internal_status = ?) or ( not ?) ) AND COALESCE(order_no,'') ilike ? and COALESCE(salesman_no, '') ILIKE ? ", internalStatus, byStatus, orderNumber, salesNo).Error
		if err != nil {
			fmt.Println("error --> ", err)
		}

		fmt.Println("order--> ", orders)

		// err = db.Order("order_date DESC").Offset(offset).Limit(limit).Find(&orders).Error
	}

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// RejectSO ...
func RejectSO(orderID int64) (errCode string, errDesc string) {
	fmt.Println("Order Rejected ..... ", orderID)
	var order dbmodels.Order
	db := GetDbCon()
	db.Debug().LogMode(true)

	r := db.Model(&order).Where("id = ?", orderID).Update(dbmodels.Order{InternalStatus: 4})
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG

}

// RjctSOByMerchantSupp ...
func RjctSOByMerchantSupp(merchantCode, supplierCode string) (errCode string, errDesc string) {
	fmt.Println("Order Rejected ..... MerchantCode : ", merchantCode, ", SupplierCode : ", supplierCode)
	var order dbmodels.Order
	db := GetDbCon()
	db.Debug().LogMode(true)

	r := db.Model(&order).Where("merchant_code = ? and supplier_code", supplierCode).Update(dbmodels.Order{InternalStatus: 4})
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

//UpdatePaymentManual ...
func UpdatePaymentManual(order *dbmodels.Order) (errCode string, errDesc string) {

	fmt.Println(" Update payment ------------------------------------------ ")
	var newOrder dbmodels.Order
	db := GetDbCon()
	db.Debug().LogMode(true)

	// r := db.Model(&order).Where("id = ?", orderID).Update(dbmodels.Order{Autodebet: 1, InternalStatus: internalStatus})
	// r := db.Model(&newOrder).Where("id = ?", order.ID).Update(dbmodels.Order{OrderNo: order.OrderNo, StatusCode: "001", WarehouseCode: order.WarehouseCode, InternalStatus: 1, OrderDate: order.OrderDate})
	r := db.Model(&newOrder).Where("id = ?", order.ID).Update(dbmodels.Order{InternalStatus: 3, PaymentNote: order.PaymentNote, ManualPaymentCode: order.ManualPaymentCode, ManualPaymentNumber: order.ManualPaymentNumber, ManualPaymentDate: order.ManualPaymentDate})
	if r.Error != nil {
		errCode = constants.ERR_CODE_80
		errDesc = r.Error.Error()
		fmt.Println("Error update ", errDesc)
		return
	}

	// fmt.Println("Order [database]=> order id", order.OrderNo)

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}
