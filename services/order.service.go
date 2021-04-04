package services

import (
	"distribution-system-be/constants"
	"distribution-system-be/database"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"fmt"
	"time"
)

// OrderService ...
type OrderService struct {
}

// GetDataOrderById ...
func (o OrderService) GetDataOrderById(orderID int64) dbmodels.SalesOrder {

	var res dbmodels.SalesOrder
	// var err error
	res, _ = database.GetSalesOrderByOrderId(orderID)

	return res
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

// Save ...
func (o OrderService) Save(order *dbmodels.SalesOrder) (errCode, errDesc, orderNo string, orderID int64, status int8) {

	if order.ID == 0 {
		newOrderNo, errCode, errMsg := generateNewOrderNo()
		if errCode != constants.ERR_CODE_00 {
			return errCode, errMsg, "", 0, 0
		}
		order.SalesOrderNo = newOrderNo
		order.Status = 10
		order.SalesmanID = dto.CurrUserId
	}
	order.LastUpdateBy = dto.CurrUser
	order.LastUpdate = time.Now()

	// fmt.Println("isi order ", order)
	err, errDesc, newID, status := database.SaveSalesOrderNo(order)
	if err != constants.ERR_CODE_00 {
		return err, errDesc, "", 0, 0
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, order.SalesOrderNo, newID, status
}

// Approve ...
func (o OrderService) Approve(order *dbmodels.SalesOrder) (errCode, errDesc string) {

	// cek qty
	valid, errCode, errDesc := validateQty(order.ID)
	if !valid {
		return errCode, errDesc
	}
	// fmt.Println("isi order ", order)
	err, errDesc := database.SaveSalesOrderApprove(order)
	if err != constants.ERR_CODE_00 {
		return err, errDesc
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

func validateQty(orderID int64) (isValid bool, errCode, errDesc string) {

	salesOrderDetails := database.GetAllDataDetail(orderID)
	for idx, orderDetail := range salesOrderDetails {
		fmt.Println("idx -> ", idx)

		product, errCodeProd, _ := database.FindProductByID(orderDetail.ProductID)
		if errCodeProd != constants.ERR_CODE_00 {
			return false, errCodeProd, fmt.Sprintf("[%v] Product not found or inactive !", orderDetail.ProductID)
		}
		curQty := product.QtyStock
		orderQty := orderDetail.QtyOrder

		if orderQty > curQty {
			return false, "99", fmt.Sprintf("[%v] qty order = %v more than qty stock = %v!", product.Name, orderQty, curQty)
		}
	}
	return true, "", ""
}

// Reject ...
func (o OrderService) Reject(order *dbmodels.SalesOrder) (errCode, errDesc string) {

	// cek qty
	// validateQty()
	// fmt.Println("isi order ", order)
	err, errDesc := database.RejectSalesOrder(order)
	if err != constants.ERR_CODE_00 {
		return err, errDesc
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

func generateNewOrderNo() (newOrderNo string, errCode string, errMsg string) {

	t := time.Now()
	bln := t.Format("01")
	thn := t.Format("06")
	header := "SO"

	err, number, errdesc := database.AddSequence(bln, thn, header)
	if err != constants.ERR_CODE_00 {
		return "", err, errdesc
	}
	newNumb := fmt.Sprintf("00000%v", number)
	newNumb = newNumb[len(newNumb)-5 : len(newNumb)]
	newOrderNo = fmt.Sprintf("%v%v%v%v", header, thn, bln, newNumb)

	return newOrderNo, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG

}

// // PrintPdf ...
// func (o OrderService) PrintPdf(order *dbmodels.Order) (errCode string, errDesc string) {

// 	// if err, errDesc := database.SaveSalesOrderNo(order); err != constants.ERR_CODE_00 {
// 	// 	return err, errDesc
// 	// }

// 	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
// }
