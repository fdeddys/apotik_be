package services

import (
	"oasis-be/database"
	"oasis-be/models"
	"oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"fmt"
	"time"
	"oasis-be/constants"
	"strings"
	"strconv"
)

type IssuerService struct {

}

// Save Data Issuer
func (i IssuerService) SaveDataIssuer(issuer *dbmodels.Issuer) models.Response {
	var lastIssuer dbmodels.Issuer
	var err error
	var code int64
	var codeIssuer string

	lastIssuer, err = database.GetLastIssuer()

	if err != nil {
		var res models.Response
		res.ErrCode = "05"
		res.ErrDesc = "Failed load data"
	}else{
		if lastIssuer == (dbmodels.Issuer{}) {
			code = 1
		}else {
			codeIssuer = strings.TrimPrefix(lastIssuer.Code, string(lastIssuer.Code[0]))
			code, err = strconv.ParseInt(codeIssuer, 10, 64)
			code = code + 1
		}
		codeIssuer = string(lastIssuer.Code[0]) + fmt.Sprintf("%06d", code)
	}

	issuer.Code = codeIssuer
	issuer.LastUpdate = time.Now()
	issuer.LastUpdateBy = dto.CurrUser

	res := database.SaveIssuer(issuer)
	fmt.Println("save : ", res)

	return res
}



// Update Data Issuer
func (i IssuerService) UpdateDataIssuer(issuer *dbmodels.Issuer) models.Response {
	var data dbmodels.Issuer

	data.ID = issuer.ID
	data.Name = issuer.Name
	data.Code = issuer.Code
	data.LastUpdate = time.Now()
	data.LastUpdateBy = "system"

	res := database.UpdateIssuer(&data)
	fmt.Println("update : ", res)

	return res
}


// Get Data Issuer Paging
func (i IssuerService) GetDataIssuerPaging(param dto.FilterName, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetIssuerPaging(param, offset, limit)

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


// Get Data List Issuer
func (i IssuerService) GetDataIssuerList() models.ResponseIssuer {
	var res models.ResponseIssuer

	data, err := database.GetListIssuer()

	if err != nil {
		return res
	}

	res.Data = data
	return res
}


// Get Data List Issuer
func (i IssuerService) GetDataIssuerListByName(name string) models.ContentResponse {
	var res models.ContentResponse

	data, err := database.GetListIssuerBySearch(name)

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

// Get Data Last Issuer 
func (i IssuerService) GetDataLastIssuer() models.ContentResponse {
	var res models.ContentResponse

	data, err := database.GetLastIssuer()

	if err != nil {
		res.ErrCode = "05"
		res.ErrDesc = "Failed load data"
		res.Contents = ""
	}

	res.ErrCode 	= constants.ERR_CODE_00
	res.ErrDesc 	= constants.ERR_CODE_00_MSG
	res.Contents 	= data

	return res
}


