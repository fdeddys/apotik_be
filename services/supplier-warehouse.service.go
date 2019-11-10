package services

import (
	"oasis-be/database"
	"oasis-be/models"
	"oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"fmt"
	"time"
	// "strings"
	// "strconv"
)

type SupplierWarehouseService struct {

}

// Save Data Supplier Warehouse
func (s SupplierWarehouseService) SaveDataSupplierWarehouse(supplierWarehouses *dbmodels.SupplierWarehouse) models.Response {
	var supplierWarehouse dbmodels.SupplierWarehouse
	// var code int64
	// var codeWarehouse string

	// supplierWarehouse, err := database.GetLastSupplierWarehouse()
	// if err != nil {

	// }else{
	// 	if supplierWarehouse != (dbmodels.SupplierWarehouse{}){
	// 		if supplierWarehouse.Code == "" {
	// 			code = 1
	// 		}else{
	// 			codeWarehouse = strings.TrimPrefix(supplierWarehouse.Code, "SW")
	// 			code, err = strconv.ParseInt(codeWarehouse, 10, 64) 
	// 			code = code + 1
	// 		}
	// 	} else {
	// 		code = 1
	// 	}
	// 	codeWarehouse = "SW" + fmt.Sprintf("%06d", code)
	// }

	// supplierWarehouses.Code = codeWarehouse
	supplierWarehouses.LastUpdate = time.Now()
	supplierWarehouse.LastUpdateBy = dto.CurrUser

	res := database.SaveSupplierWarehouse(supplierWarehouses)
	fmt.Println("save : ", res)

	return res
}

// Update Data Supplier Warehouse
func (s SupplierWarehouseService) UpdateDataSupplierWarehouse(supplierWarehouse *dbmodels.SupplierWarehouse) models.Response {
	var data dbmodels.SupplierWarehouse

	data.ID = supplierWarehouse.ID
	data.Code = supplierWarehouse.Code
	data.SupplierId = supplierWarehouse.SupplierId
	data.Description = supplierWarehouse.Description
	data.Url = supplierWarehouse.Url
	data.LastUpdateBy = dto.CurrUser
	data.LastUpdate = time.Now()

	res := database.UpdateSupplierWarehouse(&data)
	fmt.Println("update : ", res)

	return res
}


// Get Data Supplier Warehouse Paging
func (s SupplierWarehouseService) GetDataSupplierWarehousePaging(param dto.FilterSupplierWarehouseDto, page int, limit int, supplier_id int64) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetSupplierWarehousePaging(param, offset, limit, supplier_id)

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


func (s SupplierWarehouseService) GetDataWarehouseBySupplier(supplier_code string) models.ResponseSupplierWarehouse {
	var res models.ResponseSupplierWarehouse

	data, err := database.GetWarehouseBySupplierId(supplier_code)

	if err != nil {
		fmt.Println("Error")
	}

	res.Data = data
	return res
}

func (s SupplierWarehouseService) GetDataLastSupplierWarehouse() dbmodels.SupplierWarehouse {
	var res dbmodels.SupplierWarehouse
	var err error

	res, err = database.GetLastSupplierWarehouse()
	if err != nil {
		return res
	}
	return res
}
