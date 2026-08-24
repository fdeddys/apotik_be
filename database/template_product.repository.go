package database

import (
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	"fmt"
	"time"
)

// GetTemplateProducts ...
func GetTemplateProducts() ([]dbmodels.TemplateProduct, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var list []dbmodels.TemplateProduct
	err := db.Order("id ASC").Find(&list).Error
	return list, err
}

// ClearTemplateProducts ...
func ClearTemplateProducts() models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	// We can delete all or truncate
	err := db.Exec("TRUNCATE TABLE public.template_product RESTART IDENTITY").Error
	if err != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed to clear template product table: " + err.Error()
		return res
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"
	return res
}

// SaveTemplateProduct ...
func SaveTemplateProduct(tp dbmodels.TemplateProduct) error {
	db := GetDbCon()
	db.Debug().LogMode(true)

	return db.Save(&tp).Error
}

// ProcessTemplateToProduct ...
func ProcessTemplateToProduct() models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	var templates []dbmodels.TemplateProduct
	err := db.Find(&templates).Error
	if err != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed to fetch template products: " + err.Error()
		return res
	}

	for _, temp := range templates {
		if temp.Code == "" {
			continue
		}
		var product dbmodels.Product
		errFind := db.Where("code = ?", temp.Code).First(&product).Error
		if errFind != nil {
			fmt.Printf("Product with Code %s not found\n", temp.Code)
			continue
		}

		product.SellPrice = temp.HargaJual
		product.Status = temp.Status
		if temp.Nama != "" {
			product.Name = temp.Nama
		}
		if temp.Plu != "" {
			product.PLU = temp.Plu
		}
		product.LastUpdateBy = "system upload"
		product.LastUpdate = time.Now()

		errSave := db.Save(&product).Error
		if errSave != nil {
			fmt.Printf("Failed to save product %s: %s\n", temp.Code, errSave.Error())
		}
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"
	return res
}
