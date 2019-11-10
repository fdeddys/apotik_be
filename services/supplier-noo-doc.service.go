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

type SupplierNooDocService struct {

}

// Save Data Supplier Noo Doc
func (s SupplierNooDocService) SaveDataSupplierNooDoc(supplierNooDoc *dbmodels.SupplierNooDoc) models.Response {
	var res models.Response

	supplierNooDoc.LastUpdate = time.Now()
	supplierNooDoc.LastUpdateBy = dto.CurrUser

	// var listNooDoc []dbmodels.SupplierNooDoc
	listNooDoc, err := database.GetNooDocBySuppAndMerch(supplierNooDoc.SupplierID, supplierNooDoc.MerchantCode)

	if err != nil{

	}else{
		if len(listNooDoc) > 0 {
			res.ErrCode = "05"
			res.ErrDesc = "Data existing"
			return res
		}else{
			res := database.SaveSupplierNooDoc(supplierNooDoc)
			fmt.Println("save : ", res)
			return res
		}
	}

	// fmt.Println("nilai:", len(listNooDoc))

	return res
}

//Update Data Supplier Noo Doc
func (s SupplierNooDocService) UpdateDataSupplierNooDoc(supplierNooDoc *dbmodels.SupplierNooDoc) models.Response {
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

func (s SupplierNooDocService) GetDataNooDocBySupplier(supplier_id int64) []dbmodels.Lookup{
	var res []dbmodels.Lookup
	res = database.GetNooDocBySupplierId(supplier_id)
	return res
}

func (s SupplierNooDocService) CheckDataNooMerchantBySupplier(supplier_id int64, merchant_id int64) []dbmodels.SupplierNooDoc{
	var res []dbmodels.SupplierNooDoc
	res = database.CheckNooMerchantBySupplier(supplier_id, merchant_id)
	return res
}

func (s SupplierNooDocService) GetDataNOODocById(id int64) []dbmodels.SupplierNooDoc {
	var res []dbmodels.SupplierNooDoc
	res = database.GetNOODocById(id)
	return res
}


// List Approve FE NOO Doc
func (s SupplierNooDocService) GetDataSupplierNooDocApprovePaging(param dto.FilterSupplierNooDocDto, page int, limit int, supplier_id int64) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetSupplierNooDocApprovePaging(param, offset, limit, supplier_id)

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


func (s SupplierNooDocService) GetStatusApproveByMerchantAndSupplier(supplier_id int64, merchant_code string) models.ContentResponse {
	var res models.ContentResponse
	
	data, _, _, err := database.GetStatusApproveByMerchantAndSupplier(supplier_id, merchant_code)

	if err != nil {
		res.Contents = nil
		res.ErrCode = constants.ERR_CODE_51
		res.ErrDesc = constants.ERR_CODE_51_MSG
		return res
	}

	if len(data) == 0 {
		res.Contents = data
		res.ErrCode = constants.ERR_CODE_40
		res.ErrDesc = constants.ERR_CODE_40_MSG
		return res
	}else{
		res.Contents = data
		res.ErrCode = constants.ERR_CODE_00
		res.ErrDesc = constants.ERR_CODE_00_MSG
	}

	// res.Contents = data
	// res.ErrCode = errCode
	// res.ErrDesc = errDesc

	return res
}