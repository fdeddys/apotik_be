package services

import (
	"distribution-system-be/database"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"strings"
)

// PelangganService ...
type PelangganService struct {
}

// GetPelangganPaging ...
func (PelangganService) GetPelangganPaging(param dto.FilterName, page int, limit int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetPelangganPaging(param, offset, limit)
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

// SavePelanggan ...
func (PelangganService) SavePelanggan(pelanggan *dbmodels.Pelanggan) models.Response {
	// Validation
	if strings.TrimSpace(pelanggan.Nama) == "" {
		var res models.Response
		res.ErrCode = "03"
		res.ErrDesc = "Nama pelanggan tidak boleh kosong"
		return res
	}

	res := database.SavePelanggan(pelanggan)
	return res
}

// DeletePelanggan ...
func (PelangganService) DeletePelanggan(id int64) models.Response {
	res := database.DeletePelanggan(id)
	return res
}

// GetPelangganById ...
func (PelangganService) GetPelangganById(id int64) dbmodels.Pelanggan {
	return database.GetPelangganById(id)
}
