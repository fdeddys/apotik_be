package services

import (
	"distribution-system-be/database"
	"distribution-system-be/models"
	"distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type SupplierService struct {
}

// Save Data Supplier
func (s SupplierService) SaveDataSupplier(suppliers *dbmodels.Supplier) models.ResponseSupplier {
	var supplier dbmodels.Supplier
	var code int64
	var codeSupplier string

	supplier, err := database.GetLastSupplier()

	if err != nil {

	} else {
		if supplier != (dbmodels.Supplier{}) {
			if supplier.Code == "" {
				code = code + 1
			} else {
				codeSupplier = strings.TrimPrefix(supplier.Code, string(supplier.Code[0]))
				code, err = strconv.ParseInt(codeSupplier, 10, 64)
				code = code + 1
			}
		} else {
			code = 1
		}
		codeSupplier = "S" + fmt.Sprintf("%06d", code)
	}

	suppliers.Code = codeSupplier
	suppliers.LastUpdate = time.Now()
	suppliers.LastUpdateBy = dto.CurrUser

	res := database.SaveSupplier(suppliers)
	fmt.Println("save : ", suppliers)

	return res
}

// Update Data Supplier
func (s SupplierService) UpdateDataSupplier(supplier *dbmodels.Supplier) models.Response {
	var data dbmodels.Supplier

	data.ID = supplier.ID
	data.Name = supplier.Name
	data.Code = supplier.Code
	data.Alamat = supplier.Alamat
	data.Kota = supplier.Kota
	data.Status = supplier.Status
	data.LastUpdate = time.Now()
	data.LastUpdateBy = dto.CurrUser
	data.PicName = supplier.PicName
	data.PicPhone = supplier.PicPhone
	data.BankAccountName = supplier.BankAccountName

	res := database.UpdateSupplier(&data)
	fmt.Println("update : ", res)

	return res
}

// Get Data Supplier Paging
func (s SupplierService) GetDataSupplierPaging(param dto.FilterName, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetSupplierPaging(param, offset, limit)

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
