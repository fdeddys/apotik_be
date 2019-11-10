package services

import (
	"mime/multipart"
	"oasis-be/constants"
	repository "oasis-be/database"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"time"
)

// ProductService ...
type ProductService struct {
}

// GetProductFilterPaging ...
func (h ProductService) GetProductFilterPaging(param dto.FilterProduct, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := repository.GetProductListPaging(param, offset, limit)

	if err != nil {
		res.Error = err.Error()
		return res
	}

	res.Contents = data
	res.TotalRow = totalData
	res.Page = page
	res.Count = len(data)

	return res
}

// GetProductDetails ...
func (h ProductService) GetProductDetails(id int) models.ContentResponse {

	var res models.ContentResponse

	data, errCode, errDesc, err := repository.GetProductDetails(id)

	if err != nil {
		res.Contents = nil
		res.ErrCode = "02"
		res.ErrDesc = "Error query data to DB"
		return res
	}

	res.Contents = data
	res.ErrCode = errCode
	res.ErrDesc = errDesc

	return res
}

// SaveProduct ...
func (h ProductService) SaveProduct(product *dbmodels.Product) models.ContentResponse {
	product.LastUpdate = time.Now()
	product.LastUpdateBy = dto.CurrUser

	// var res models.ResponseSave
	res := repository.SaveProduct(*product)

	return res
}

// UpdateProduct ...
func (h ProductService) UpdateProduct(product *dbmodels.Product) models.NoContentResponse {
	product.LastUpdate = time.Now()
	product.LastUpdateBy = dto.CurrUser

	res := repository.UpdateProduct(*product)

	return res
}

func (h ProductService) ProductList() []dbmodels.Product {
	res := repository.ProductList()
	return res
}

func (h ProductService) UploadImage(file multipart.File, fileName string) models.NoContentResponse {
	res := repository.UploadImage(file, fileName, "product")
	return res
}

// GetProductLike ...
func (h ProductService) GetProductLike(productterms string) models.ContentResponse {
	var res models.ContentResponse

	data, errCode, errDesc, err := repository.GetProductLike(productterms)

	if err != nil {
		res.Contents = nil
		res.ErrCode = constants.ERR_CODE_51
		res.ErrDesc = constants.ERR_CODE_51_MSG
		return res
	}

	res.Contents = data
	res.ErrCode = errCode
	res.ErrDesc = errDesc

	return res
}
