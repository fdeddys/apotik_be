package services

import (
	"oasis-be/database"
	"oasis-be/models"
	"oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"fmt"
	"time"
)

type SupplierMerchantService struct {

}


// Save Data Supplier Merchant
func (s SupplierMerchantService) SaveDataSupplierMerchant(supplierMerchant *dbmodels.SupplierMerchant) models.Response {
	supplierMerchant.LastUpdate = time.Now()
	supplierMerchant.LastUpdateBy = dto.CurrUser

	res := database.SaveSupplierMerchant(supplierMerchant)
	fmt.Println("save : ", res)

	return res
}

// Update Data Supplier Merchant
func (s SupplierMerchantService) UpdateDataSupplierMerchant(supplierMerchant *dbmodels.SupplierMerchant) models.Response {
	var data dbmodels.SupplierMerchant

	data.ID = supplierMerchant.ID
	data.MerchantCode = supplierMerchant.MerchantCode
	data.SupplierId = supplierMerchant.SupplierId
	data.LastUpdate = time.Now()
	data.LastUpdateBy = dto.CurrUser

	res := database.UpdateSupplierMerchant(&data)
	fmt.Println("update : ", res)

	return res
}

// Get Data Supplier Merchant Paging
func (s SupplierMerchantService) GetDataSupplierMerchantPaging(param dto.FilterSupplierMerchantDto, page int, limit int, supplier_id int64) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetSupplierMerchantPaging(param, offset, limit, supplier_id)

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

func (s SupplierMerchantService) GetDataListMerchantBySupplier(supplier_id int64, merchant_id int64) []dbmodels.SupplierMerchant {
	var res []dbmodels.SupplierMerchant
	res = database.GetListMerchantBySupplier(supplier_id, merchant_id)
	return res
}


// func (s SupplierMerchantService) CheckFirstOrderMerchantBySupplier(supplier_id int, merchant_id int) int{
// 	var firstOrder int
// 	firstOrder = database.FirstOrderMerchantBySupplier(supplier_id, merchant_id)
// 	return firstOrder
// }