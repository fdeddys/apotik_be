package database

import (
	"fmt"
	"log"
	constants "oasis-be/constants"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"

	// "oasis-be/utils/http"
	"strconv"
	"strings"
	"sync"

	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

// GetProductDetails ...
func GetProductDetails(id int) ([]dbmodels.Product, string, string, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var product []dbmodels.Product
	err := db.Model(&dbmodels.Product{}).Preload("ProductGroup").Preload("Brand").Preload("UomLookup", "lookup_group=?", "UOM").Where("id = ?", &id).First(&product).Error
	// .Preload("StockLookup", "lookup_group=?", "STOCK_STATUS")

	if err != nil {
		return nil, "02", "Error query data to DB", err
	}
	// else {
	return product, "00", "success", nil
	// }
}

// GetProductListPaging ...
func GetProductListPaging(param dto.FilterProduct, offset int, limit int) ([]dbmodels.Product, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var product []dbmodels.Product
	// var uom []dbmodels.Lookup
	var total int
	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&product).Error
		if err != nil {
			return product, 0, err
		}
		return product, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go ProductQuerys(db, offset, limit, &product, param, errQuery)
	go AsyncQueryCount(db, &total, param, &dbmodels.Product{}, "name", errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		return product, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("err-->", resErrCount)
		return product, 0, resErrCount
	}

	for i := 0; i < len(product); i++ {
		isExist := CheckImage(product[i].Code, "product")
		if isExist {
			product[i].IMG1 = GetImage(product[i].Code, "product")
		} else {
			product[i].IMG1 = GetImage("no_image", "product")
		}
	}

	return product, total, nil
}

// UpdateProduct ...
func UpdateProduct(updatedProduct dbmodels.Product) models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	var product dbmodels.Product
	err := db.Model(&dbmodels.Product{}).Where("id=?", &updatedProduct.ID).First(&product).Error
	if err != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error select data to DB"
	}

	product.ID = updatedProduct.ID
	product.Name = updatedProduct.Name
	product.Status = updatedProduct.Status
	product.LastUpdateBy = updatedProduct.LastUpdateBy
	product.LastUpdate = updatedProduct.LastUpdate
	product.Code = updatedProduct.Code
	product.ProductGroupID = updatedProduct.ProductGroupID
	product.BrandID = updatedProduct.BrandID
	// product.StockStatus = updatedProduct.StockStatus
	product.UOM = updatedProduct.UOM

	err2 := db.Save(&product)
	if err2 != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error update data to DB"
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"

	return res
}

//SaveProduct ...
func SaveProduct(product dbmodels.Product) models.ContentResponse {
	var res models.ContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	prefix := product.Name[:3]
	product.Code = GenerateProductCode(strings.ToUpper(prefix))

	if r := db.Save(&product); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error save data to DB"
		res.Contents = r.Error
	}

	// url := beego.AppConfig.String("kafka.rest-server")
	// var kafka dbmodels.KafkaReq
	// var productVendor dbmodels.ProductVendor

	// productVendor.Code = product.Code
	// productVendor.Name = product.Name
	// productVendor.Uom = product.UOM

	// // kafka.Topic = "oasis-add-product-uki-topic"
	// kafka.Data = productVendor

	// _, err := http.HttpPost(url, kafka, "60s", 1)
	// if err != nil {
	// 	logs.Error("Error hit kafka")
	// }

	res.ErrCode = "00"
	res.ErrDesc = "Success"
	res.Contents = product
	return res
}

// ProductQuerys ...
func ProductQuerys(db *gorm.DB, offset int, limit int, product *[]dbmodels.Product, param dto.FilterProduct, resChan chan error) {

	// var criteriaUserName = "%"
	// if strings.TrimSpace(param.Name) != "" {
	criteriaUserName := param.Name //+ criteriaUserName
	// }

	// err := db.Set("gorm:auto_preload", true).Order("name ASC").Offset(offset).Limit(limit).Find(&user, "name like ?", criteriaUserName).Error
	err := db.Preload("Brand").Preload("ProductGroup").Preload("UomLookup", "lookup_group=?", "UOM").Order("name ASC").Offset(offset).Limit(limit).Find(&product, "name ~* ?", criteriaUserName).Error
	// .Preload("StockLookup", "lookup_group=?", "STOCK_STATUS")
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

func ProductList() []dbmodels.Product {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var product []dbmodels.Product
	err := db.Preload("Brand").Preload("ProductGroup").Preload("UomLookup", "lookup_group=?", "UOM").Order("name ASC").Find(&product).Error
	// .Preload("StockLookup", "lookup_group=?", "STOCK_STATUS")

	if err != nil {
		return product
	}
	return product

}

// GenerateProductCode ...
func GenerateProductCode(prefix string) string {
	db := GetDbCon()
	db.Debug().LogMode(true)

	// err := db.Order(order).First(&models)

	sprefix := prefix
	if prefix == "" {
		sprefix = "%"
	} else {
		sprefix = "%" + prefix + "%"
	}

	var product []dbmodels.Product
	err := db.Model(&dbmodels.Product{}).Order("id desc").Where("code ILIKE ?", sprefix).First(&product).Error
	// err := db.Model(&dbmodels.Brand{}).Where("id = 200000").Order("id desc").First(&brand).Error

	if err != nil {
		return prefix + "000001"
	}
	if len(product) > 0 {
		// fmt.Printf("ini latest code nya : %s \n", brand[0].Code)
		woprefix := strings.TrimPrefix(product[0].Code, prefix)
		latestCode, err := strconv.Atoi(woprefix)
		if err != nil {
			fmt.Printf("error")
			return prefix + "000001"
		}
		// fmt.Printf("ini latest code nya : %d \n", latestCode)
		wpadding := fmt.Sprintf("%06s", strconv.Itoa(latestCode+1))
		// fmt.Printf("ini pake padding : %s \n", "B"+wpadding)
		return prefix + wpadding
	}
	return prefix + "000001"

}

//GetProductLike ...
func GetProductLike(productterms string) ([]dbmodels.Product, string, string, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var product []dbmodels.Product
	err := db.Model(&dbmodels.Product{}).Where("name ~* ?", &productterms).Find(&product).Error

	if err != nil {
		return nil, constants.ERR_CODE_51, constants.ERR_CODE_51_MSG, err
	}
	return product, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, nil
}