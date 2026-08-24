package database

import (
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"log"
	"strings"
	"sync"

	"distribution-system-be/constants"

	"github.com/jinzhu/gorm"
)

// GetPelangganPaging ...
func GetPelangganPaging(param dto.FilterName, offset int, limit int) ([]dbmodels.Pelanggan, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var pelanggan []dbmodels.Pelanggan
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&pelanggan).Error
		if err != nil {
			return pelanggan, 0, err
		}
		return pelanggan, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerysPelanggan(db, offset, limit, &pelanggan, param, errQuery)
	go AsyncQueryCountsPelanggan(db, &total, param, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return pelanggan, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return pelanggan, 0, resErrCount
	}
	return pelanggan, total, nil
}

// AsyncQueryCountsPelanggan ...
func AsyncQueryCountsPelanggan(db *gorm.DB, total *int, param dto.FilterName, resChan chan error) {
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = "%" + param.Name + "%"
	}
	err := db.Model(&dbmodels.Pelanggan{}).Where("nama ilike ?", searchName).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQuerysPelanggan ...
func AsyncQuerysPelanggan(db *gorm.DB, offset int, limit int, pelanggan *[]dbmodels.Pelanggan, param dto.FilterName, resChan chan error) {
	var searchName = "%"
	if strings.TrimSpace(param.Name) != "" {
		searchName = "%" + param.Name + "%"
	}

	err := db.Order("nama ASC").Offset(offset).Limit(limit).Find(pelanggan, "nama ilike ?", searchName).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// SavePelanggan ...
func SavePelanggan(pelanggan *dbmodels.Pelanggan) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response
	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	r := db.Save(&pelanggan)
	if r.Error != nil {
		res.ErrCode = constants.ERR_CODE_30
		res.ErrDesc = constants.ERR_CODE_30_MSG + " " + r.Error.Error()
	}

	return res
}

// GetPelangganById ...
func GetPelangganById(id int64) dbmodels.Pelanggan {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var pelanggan dbmodels.Pelanggan
	err := db.Where("id = ?", id).Find(&pelanggan).Error
	if err != nil {
		return pelanggan
	}
	return pelanggan
}

// DeletePelanggan ...
func DeletePelanggan(id int64) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response
	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	err := db.Where("id = ?", id).Delete(&dbmodels.Pelanggan{}).Error
	if err != nil {
		res.ErrCode = constants.ERR_CODE_30
		res.ErrDesc = err.Error()
	}
	return res
}
