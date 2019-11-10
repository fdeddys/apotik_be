package services

import (
	"fmt"
	"oasis-be/constants"
	"oasis-be/database"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"strconv"
	"strings"
	"time"
)

type MerchantService struct {
}

// Get Data Merchant Paging
func (m MerchantService) GetDataMerchantPaging(param dto.FilterName, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetMerchantPaging(param, offset, limit)
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

// Save Data Merchant
func (m MerchantService) SaveDataMerchant(merchant *dbmodels.Merchant) models.Response {
	var listMerchant dbmodels.Merchant
	var code int64
	var codeMerchant string

	listMerchant, error := database.GetLastMerchant()
	code = 0

	if error != nil {
		code = 1
	} else {
		if listMerchant != (dbmodels.Merchant{}) {
			if listMerchant.Code == "" {
				code = code + 1
			} else {
				codeMerchant = strings.TrimPrefix(listMerchant.Code, string(listMerchant.Code[0]))
				code, error = strconv.ParseInt(codeMerchant, 10, 64)
				code = code + 1
			}
		} else {
			code = 1
		}
	}
	codeMerchant = "M" + fmt.Sprintf("%06d", code)
	// if len(listMerchant) > 0 {
	// 	codeMerchant = strings.TrimPrefix(listMerchant[0].Code, string(listMerchant[0].Code[0]))
	// 	code, err := strconv.ParseInt(codeMerchant, 10, 64)
	// 	if err != nil {
	// 		var res models.Response
	// 		res.ErrCode = "05"
	// 		res.ErrCode = "Failed parse code merchant to integer"
	// 	}
	// 	code = code + 1
	// 	codeMerchant = fmt.Sprintf("%06d", code)
	// }else{
	// 	code = 1;
	// 	codeMerchant = fmt.Sprintf("%06d", code)
	// }

	merchant.Code = codeMerchant
	merchant.LastUpdate = time.Now()
	merchant.LastUpdateBy = dto.CurrUser

	res := database.SaveMerchant(merchant)
	fmt.Println("save : ", merchant)

	return res
}

// Update Data Merchant
func (m MerchantService) UpdateDataMerchant(merchant *dbmodels.Merchant) models.Response {
	var data dbmodels.Merchant
	data.ID = merchant.ID
	data.Name = merchant.Name
	data.Code = merchant.Code
	data.IssuerCode = merchant.IssuerCode
	data.Top = merchant.Top
	data.Status = merchant.Status
	data.LastUpdateBy = dto.CurrUser
	data.LastUpdate = time.Now()

	// var res models.Response
	res := database.UpdateMerchant(&data)
	fmt.Println("update : ", res)

	return res
}

// Get Data List Merchant
func (m MerchantService) GetDataMerchantList() models.ResponseMerchant {
	var res models.ResponseMerchant

	data, err := database.GetListMerchant()

	if err != nil {
		return res
	}

	res.Data = data
	return res
}

func (m MerchantService) GetDataMerchantListByName(name string) models.ContentResponse {
	var res models.ContentResponse

	data, err := database.GetListMerchantBySearch(name)

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

func (m MerchantService) GetMerchantById(merchant_id int64) dbmodels.Merchant {
	var res dbmodels.Merchant
	res = database.GetMerchantById(merchant_id)
	return res
}

// func (m MerchantService) GetDataCheckOrder(supplier string, merchant string) []dbmodels.Order {
// 	res := database.GetOrderBySupplierAndMerchant(supplier, merchant)
// 	return res
// }
