package services

import (
	"fmt"
	"oasis-be/constants"
	"oasis-be/database"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"time"
)

// OrderService ...
type OrderStatusService struct {
}

// SaveOrderStatus ...
func SaveOrderStatus(orderNo, statusName string) (errCode string, errDesc string) {
	// func (o OrderStatusService) SaveOrderStatus(orderNo, statusName string) (errCode string, errDesc string) {

	// func (o OrderStatusService) Save(orderStatus *dbmodels.OrderStatus) (errCode string, errDesc string) {

	fmt.Println("Order STATUS... -------->", orderNo)
	orderStatus := dbmodels.OrderStatus{}
	orderStatus.LastUpdateBy = dto.CurrUser
	orderStatus.LastUpdate = time.Now()
	orderStatus.OrderNo = orderNo
	orderStatus.StatusName = statusName

	err, errDesc := database.SaveOrderStatus(&orderStatus)
	if err != constants.ERR_CODE_00 {
		return err, errDesc
	}

	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// GetDataOrderStatusPage ...
func (o OrderStatusService) GetDataOrderStatusPage(param dto.FilterOrderDetail, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetOrderStatusPage(param, offset, limit)

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
