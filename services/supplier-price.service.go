package services

import (
	"oasis-be/database"
	"oasis-be/models"
	"oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"fmt"
	"time"
	"oasis-be/constants"
)

type SupplierPriceService struct {

}

// Save Data Supplier Price
func (s SupplierPriceService) SaveDataSupplierPrice(supplierPrice *dbmodels.SupplierPrice) models.Response {
	supplierPrice.LastUpdate = time.Now()
	supplierPrice.LastUpdateBy = dto.CurrUser

	res := database.SaveSupplierPrice(supplierPrice)
	fmt.Println("save : ", res)

	return res
}

//Update Data Supplier Price
func (s SupplierPriceService) UpdateDataSupplierPrice(supplierPrice *dbmodels.SupplierPrice) models.Response {
	var data dbmodels.SupplierPrice

	data.ID = supplierPrice.ID
	data.SupplierId = supplierPrice.SupplierId
	data.ProductCode = supplierPrice.ProductCode
	data.SellPrice = supplierPrice.SellPrice
	data.BuyPrice = supplierPrice.BuyPrice
	data.SellMarginCode = supplierPrice.SellMarginCode
	data.LastUpdate = time.Now()
	data.LastUpdateBy = dto.CurrUser
	data.PriceMargin = supplierPrice.PriceMargin

	res := database.UpdateSupplierPrice(&data)
	fmt.Println("update : ", res)

	return res
}


// Get Data Supplier Price Paging
func (s SupplierPriceService) GetDataSupplierPricePaging(param dto.FilterSupplierPriceDto, page int, limit int, supplier_id int64) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetSupplierPricePaging(param, offset, limit, supplier_id)

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

// Get List Product By Supplier Id
func (s SupplierPriceService) GetDataListProductBySuppId(supplier_id int64) models.ContentResponse {
	var res models.ContentResponse

	data, err := database.GetListProductById(supplier_id)
	if err != nil {
		res.ErrCode = "05"
		res.ErrDesc = "Failed load data"
		return res
	}

	res.ErrCode = "00"
	res.ErrDesc = constants.ERR_CODE_00_MSG
	res.Contents = data

	return res
}