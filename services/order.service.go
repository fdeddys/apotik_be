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
func (o OrderService) Save(order *dbmodels.SalesOrder) (errCode, errDesc, orderNo string, orderID int64) {

	if order.ID == 0 {
		newOrderNo, errCode, errMsg := generateNewOrderNo()
		if errCode != constants.ERR_CODE_00 {
			return errCode, errMsg, "", 0
		}
		order.SalesOrderNo = newOrderNo
		order.Status = 10
	}
	order.LastUpdateBy = dto.CurrUser
	order.LastUpdate = time.Now()

	// fmt.Println("isi order ", order)
	err, errDesc, newID := database.SaveSalesOrderNo(order)
	if err != constants.ERR_CODE_00 {
		return err, errDesc, "", 0
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, order.SalesOrderNo, newID
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
