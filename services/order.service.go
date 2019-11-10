package services

import (
	"fmt"
	"oasis-be/constants"
	"oasis-be/database"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"oasis-be/redis"
	"time"
)

// OrderService ...
type OrderService struct {
}

// ManualPay ...
func (o OrderService) ManualPay(order *dbmodels.Order) (errCode string, errDesc string) {

	fmt.Println("isi order ", order)

	errCode, errMsg, _ := database.FindOrderReadyToPay(order.OrderNo)
	if errCode != constants.ERR_CODE_00 {
		return errCode, errMsg
	}

	if err, errDesc := database.UpdatePaymentManual(order); err != constants.ERR_CODE_00 {

		// go database.InsertStatusOrder(orderID, "Payment manual success")
		return err, errDesc
	}

	SaveOrderStatus(order.OrderNo, constants.SO_MANUAL_PAY)
	SaveOrderStatus(order.OrderNo, constants.SO_COMPLETE)
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// PaymentAutodebet ...
// 1. search so ready with status ready for payment
// 2. set status
// 3. insert to list status
func (o OrderService) PaymentAutodebet(autoDebetRequest *dto.OrderAutodebetRequest) (errCode string, errDesc string) {

	fmt.Println("isi order ", autoDebetRequest)

	errCode, errMsg, orderID := database.FindOrderReadyToPay(autoDebetRequest.OrderNo)
	if errCode != constants.ERR_CODE_00 {
		return errCode, errMsg
	}

	if err, errDesc := database.UpdateInternalStatus(orderID, autoDebetRequest.PaymentSuccess); err != constants.ERR_CODE_00 {

		// go SaveOrderStatus(autoDebetRequest.OrderNo, "AUTODEBET [ "+autoDebetRequest.PaymentSuccess+" ] ")

		// go database.InsertStatusOrder(orderID, autoDebetRequest.PaymentSuccess)
		return err, errDesc
	}

	if autoDebetRequest.PaymentSuccess == "1" {
		SaveOrderStatus(autoDebetRequest.OrderNo, constants.SO_AUTODEBET_SUCCESS)
		SaveOrderStatus(autoDebetRequest.OrderNo, constants.SO_COMPLETE)
	} else {
		// success = "Failed"
		go SaveOrderStatus(autoDebetRequest.OrderNo, constants.SO_AUTODEBET_FAILED)
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// GetDataPage ...
func (o OrderService) GetDataPage(param dto.FilterOrder, page int, limit int, internalStatus int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetOrderPage(param, offset, limit, internalStatus)

	if err != nil {
		res.Error = err.Error()
		return res
	}

	res.Contents = data
	res.TotalRow = totalData
	res.Page = page
	res.Count = limit

	return res
}

// RejectSO REJECT  SO ....
func (o OrderService) RejectSO(orderNo string) (errCode string, errDesc string) {

	order, err := database.GetOrderByOrderNo(orderNo)
	if err != nil {
		return constants.ERR_CODE_50, "Cannot find record"
	}

	fmt.Println("order =>", order)
	errCode, errDesc = database.RejectSO(order.ID)
	go SaveOrderStatus(orderNo, constants.SO_REJECTED)

	return
}

// ReleaseSO ...
func (o OrderService) ReleaseSO(orderNo string) (errCode string, errDesc string) {

	order, err := database.GetOrderByOrderNo(orderNo)
	if err != nil {
		return constants.ERR_CODE_50, "Cannot find record"
	}

	var salesOrder dto.SalesOrder
	salesOrder.Code = order.OrderNo
	salesOrder.WarehouseCode = order.WarehouseCode
	salesOrder.CustomerCode = order.MerchantCode
	salesOrder.TransactionAt = fmt.Sprintf("%v", order.OrderDate)
	salesOrder.StateAction = "complete"

	orderDetails := database.GetAllDataDetail(order.ID)

	items := []dto.SalerOrderItem{}
	for _, orderDetail := range orderDetails {
		var item dto.SalerOrderItem
		item.ItemCode = orderDetail.ProductCode
		item.Quantity = fmt.Sprintf("%v", orderDetail.Qty)
		item.Price = fmt.Sprintf("%v", orderDetail.Price)
		item.Description = orderDetail.Product.Name
		item.Uom = orderDetail.Lookup.Code

		items = append(items, item)
	}

	salesOrder.SalesOrderItems = items
	salesOrder.SupplierCode = order.SupplierCode

	// orderNo := fmt.Sprintf("%v", req.OrderNo)
	go SaveOrderStatus(orderNo, constants.SO_RELEASE_SO)
	// go setFirstOrder(order.Supplier.ID, order.Supplier.Code, order.MerchantCode)

	ProduceToKafkaServer(salesOrder)
	// go SoService.SendtoKafka(&salesOrder)

	fmt.Println(" isi -> ", salesOrder)
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

func setFirstOrder(supplierID int64, supplierCode, merchantCode string) {

	if isFirstOrder(supplierCode, merchantCode) == true {
		database.UpdateFirstOrder(supplierID, merchantCode)
	}

}

func isFirstOrder(supplierCode, merchantCode string) bool {

	dataExis := true
	key := fmt.Sprintf("OASIS:FIRST-ORDER:%v:%v", supplierCode, merchantCode)

	if _, err := redis.GetRedisKey(key); err == nil {
		return dataExis
	}
	// cek di database ada ga?
	data := database.GetOrderBySupplierAndMerchant(supplierCode, merchantCode)

	if len(data) > 0 {
		redis.SaveRedis(key, "1")
		return dataExis
	}

	redis.SaveRedis(key, "1")
	return !dataExis
}

// Save ...
func (o OrderService) Save(order *dbmodels.Order) (errCode string, errDesc string) {

	// newOrderNo, errCode, errMsg := generateNewSO()
	// if errCode != constants.ERR_CODE_00 {
	// 	return newOrderNo, errCode, errMsg
	// }
	// order.OrderNo = newOrderNo
	// order.OrderDate = time.Now()

	// fmt.Println("isi order ", order)
	if err, errDesc := database.SaveSalesOrderNo(order); err != constants.ERR_CODE_00 {
		return err, errDesc
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

func generateNewSO() (newSO string, errCode string, errMsg string) {

	t := time.Now()
	bln := t.Format("01")
	thn := t.Format("06")

	err, number, errdesc := database.AddSequence(bln, thn)
	if err != constants.ERR_CODE_00 {
		return "", err, errdesc
	}
	newNumb := fmt.Sprintf("00000%v", number)
	newNumb = newNumb[len(newNumb)-5 : len(newNumb)]
	newSO = fmt.Sprintf("SO%v%v%v", thn, bln, newNumb)

	return newSO, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG

}

// // PrintPdf ...
// func (o OrderService) PrintPdf(order *dbmodels.Order) (errCode string, errDesc string) {

// 	// if err, errDesc := database.SaveSalesOrderNo(order); err != constants.ERR_CODE_00 {
// 	// 	return err, errDesc
// 	// }

// 	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
// }
