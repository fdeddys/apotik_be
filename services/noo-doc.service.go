package services


import (
	"oasis-be/database"
	"oasis-be/models"
	"oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"time"
	"fmt"
)

type NooDocService struct {

}

func (n NooDocService) GetDataNooDocPaging(param dto.FilterSupplierNooDocDto, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetNooDocPaging(param, offset, limit)

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

func (s SupplierNooDocService) UpdateDataNooDoc(supplierNooDoc *dbmodels.SupplierNooDoc) models.Response {
	var data dbmodels.SupplierNooDoc

	data.ID			= supplierNooDoc.ID
	data.SupplierID = supplierNooDoc.SupplierID
	data.LookupCode = supplierNooDoc.LookupCode
	data.LastUpdate = time.Now()
	data.LastUpdateBy = dto.CurrUser
	data.MerchantCode = supplierNooDoc.MerchantCode
	data.ApprovalStatus = supplierNooDoc.ApprovalStatus

	res := database.UpdateSupplierNooDoc(&data)
	fmt.Println("update : ", res)

	return res
}