package services


import (
	"oasis-be/database"
	"oasis-be/models"
	"oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"fmt"
	"time"
)

type SupplierNooChecklistService struct {

}

func (s SupplierNooChecklistService) GetDataSupplierNooChecklistPaging(param dto.FilterSupplierNooChecklistDto, page int, limit int, supplier_id int64) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetSupplierNooChecklistPaging(param, offset, limit, supplier_id)
	fmt.Println("data supplier: ", data)
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

// Save Data Supplier Noo Checklist
func (s SupplierNooChecklistService) SaveDataSupplierNooChecklist(supplierNooChecklist *dbmodels.SupplierNooChecklist) models.Response {
	supplierNooChecklist.LastUpdate = time.Now()
	supplierNooChecklist.LastUpdateBy = dto.CurrUser
	supplierNooChecklist.IsMandatory = 0

	res := database.SaveSupplierNooChecklist(supplierNooChecklist)
	fmt.Println("save : ", res)

	return res
}

//Update Data Supplier Noo Checklist
func (s SupplierNooChecklistService) UpdateDataSupplierNooChecklist(supplierNooChecklist *dbmodels.SupplierNooChecklist) models.Response {
	var data dbmodels.SupplierNooChecklist

	data.ID			= supplierNooChecklist.ID
	data.SupplierID = supplierNooChecklist.SupplierID
	data.LookupCode = supplierNooChecklist.LookupCode
	data.LastUpdate = time.Now()
	data.IsMandatory = supplierNooChecklist.IsMandatory
	data.LastUpdateBy = dto.CurrUser

	res := database.UpdateSupplierNooChecklist(&data)
	fmt.Println("update : ", res)

	return res
}


func (s SupplierNooChecklistService) GetDataNooCheckListBySupplier(supplier_id int64) []dbmodels.ResponseSupplierNooChecklist{
	var res []dbmodels.ResponseSupplierNooChecklist
	res = database.GetNooChecklistBySupplierId(supplier_id)
	return res
}