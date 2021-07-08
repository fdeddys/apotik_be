package database

import (
	constants "distribution-system-be/constants"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"distribution-system-be/utils/util"
	"fmt"
	"log"
	"reflect"

	// "distribution-system-be/utils/http"
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
	err := db.Model(&dbmodels.Product{}).Preload("ProductGroup").Preload("Brand").Preload("BigUom").Where("id = ?", &id).First(&product).Error
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

	err2 := db.Save(&product)
	if err2 != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Error update data to DB"
		return res
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"

	return res
}

//SaveProduct ...
func SaveProduct(product dbmodels.Product) models.ContentResponse {
	var res models.ContentResponse
	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	newProduct := false
	db := GetDbCon()
	tx := db.Begin()
	db.Debug().LogMode(true)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// prefix := product.Name[:3]
	if product.ID < 1 {
		product.Code = GenerateProductCode(product.Name)
		newProduct = true
	}
	// product.Code = GenerateProductCode(strings.ToUpper(prefix))

	if r := tx.Save(&product); r.Error != nil {
		res.ErrCode = constants.ERR_CODE_30
		res.ErrDesc = constants.ERR_CODE_30_MSG
		res.Contents = r.Error
		tx.Rollback()
		return res
	}
	res.Contents = product

	if !newProduct {
		tx.Commit()
		return res
	}

	// InitStock(productId)
	productId := product.ID
	warehouses, errCode, _ := GetAllWarehouse()
	if errCode != constants.ERR_CODE_00 {
		tx.Commit()
		return res
	}

	for _, warehouse := range warehouses {
		var stock dbmodels.Stock
		stock.ProductID = productId
		stock.LastUpdateBy = dto.CurrUser
		stock.LastUpdate = util.GetCurrDate()
		stock.Qty = 0
		stock.WarehouseID = warehouse.ID

		err := tx.Save(&stock).Error
		if err != nil {
			tx.Rollback()
			return res
		}

		var history dbmodels.HistoryStock
		history.Code = product.Code
		history.WarehouseID = warehouse.ID
		history.Debet = 0
		history.Description = "INIT STOCK"
		history.Hpp = 0
		history.Kredit = 0
		history.LastUpdate = util.GetCurrDate()
		history.LastUpdateBy = dto.CurrUser
		history.Name = product.Name
		history.ReffNo = ""
		history.Price = 0
		history.Saldo = 0
		history.TransDate = util.GetCurrFormatDate()
		tx.Save(&history)

	}

	tx.Commit()

	return res
}

// func InitStock(productID int64) {

// 	warehouses, errCode, _ := GetAllWarehouse()
// 	if errCode != constants.ERR_CODE_00 {
// 		return
// 	}

// 	for _, warehouse := range warehouses {
// 		SaveStock(productID, warehouse.ID)
// 	}

// }

// ProductQuerys ...
func ProductQuerys(db *gorm.DB, offset int, limit int, product *[]dbmodels.Product, param dto.FilterProduct, resChan chan error) {

	// var criteriaUserName = "%"
	// if strings.TrimSpace(param.Name) != "" {
	// criteriaUserName := param.Name + '%' //+ criteriaUserName
	varInterface := reflect.ValueOf(param)
	strQuery := varInterface.Field(0).Interface().(string)

	criteriaName := strQuery
	if criteriaName == "" {
		criteriaName = "%"
	} else {
		criteriaName = "%" + strQuery + "%"
	}

	// }

	// err := db.Set("gorm:auto_preload", true).Order("name ASC").Offset(offset).Limit(limit).Find(&user, "name like ?", criteriaUserName).Error
	err := db.Preload("Brand").Preload("ProductGroup").Preload("BigUom").Preload("SmallUom").Order("name ASC").Offset(offset).Limit(limit).Find(&product, "name ilike ?", criteriaName).Error
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
	err := db.Preload("Brand").Preload("ProductGroup").Preload("BigUom").Order("name ASC").Find(&product).Error
	// .Preload("StockLookup", "lookup_group=?", "STOCK_STATUS")

	if err != nil {
		return product
	}
	return product

}
func GenerateProductCode(name string) string {
	db := GetDbCon()
	db.Debug().LogMode(true)

	defaultCode := "X0001"
	header := name[:1]
	header = strings.ToUpper(header)

	var product []dbmodels.Product
	err := db.Where("substring(code,1,1) = ? ", header).Order("id desc").Find(&product).Error

	if err != nil {
		fmt.Println("Error not found ", err.Error())
		return header + "0001"
	}

	if len(product) > 0 {
		woprefix := strings.TrimPrefix(product[0].Code, header)
		latestCode, err := strconv.Atoi(woprefix)
		if err != nil {
			fmt.Printf("error")
			return defaultCode
		}
		wpadding := fmt.Sprintf("%04s", strconv.Itoa(latestCode+1))
		return header + wpadding
	}
	return header + "0001"
}

// // GenerateProductCode ...
// func GenerateProductCode(prefix string) string {
// 	db := GetDbCon()
// 	db.Debug().LogMode(true)

// 	// err := db.Order(order).First(&models)

// 	sprefix := prefix
// 	if prefix == "" {
// 		sprefix = "%"
// 	} else {
// 		sprefix = "%" + prefix + "%"
// 	}

// 	var product []dbmodels.Product
// 	err := db.Model(&dbmodels.Product{}).Order("id desc").Where("code ILIKE ?", sprefix).First(&product).Error
// 	// err := db.Model(&dbmodels.Brand{}).Where("id = 200000").Order("id desc").First(&brand).Error

// 	if err != nil {
// 		return prefix + "000001"
// 	}
// 	if len(product) > 0 {
// 		// fmt.Printf("ini latest code nya : %s \n", brand[0].Code)
// 		woprefix := strings.TrimPrefix(product[0].Code, prefix)
// 		latestCode, err := strconv.Atoi(woprefix)
// 		if err != nil {
// 			fmt.Printf("error")
// 			return prefix + "000001"
// 		}
// 		// fmt.Printf("ini latest code nya : %d \n", latestCode)
// 		wpadding := fmt.Sprintf("%06s", strconv.Itoa(latestCode+1))
// 		// fmt.Printf("ini pake padding : %s \n", "B"+wpadding)
// 		return prefix + wpadding
// 	}
// 	return prefix + "000001"

// }

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

// FindProductByID ...
func FindProductByID(prodID int64) (dbmodels.Product, string, string) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var product dbmodels.Product
	db.Where("id = ? and status = 1 ", prodID).First(&product)
	// fmt.Println("isi err prod ", err)
	// if err != nil {
	// 	return product, constants.ERR_CODE_51, err.Error.Error()
	// }
	return product, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// UpdateStockProductByID ...
func UpdateStockProductByID(prodID, qty, warehouseID int64) (string, string) {

	fmt.Println("Update stock", prodID, "qty ", qty)

	db := GetDbCon()
	db.Debug().LogMode(true)

	var stock dbmodels.Stock
	result := db.Where("product_id=? and warehouse_id = ?", prodID, warehouseID).Find(&stock)

	fmt.Println("Find stock for update ", prodID, " wh", warehouseID, " err => ", result)
	if stock.ID == 0 {
		fmt.Println("Find stock for update -- create new record ")
		stock.ProductID = prodID
		stock.WarehouseID = warehouseID
		stock.LastUpdateBy = dto.CurrUser
		stock.LastUpdate = util.GetCurrDate()
		stock.Qty = qty
		db.Save(&stock)
		return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
	}

	fmt.Println("UpdateStockProductByID -- Update  ")
	err := db.Model(&dbmodels.Stock{}).Where("id = ?", stock.ID).Update("qty", qty).Error
	if err != nil {
		return constants.ERR_CODE_80, err.Error()
	}
	// err := db.Model(&product).Where("id = ? ", prodID).Update("qty_stock = ?", qty)
	// fmt.Println("err prod => ", err.Error)
	// if err != nil {
	// 	fmt.Println("err prod ", err.Error)
	// 	return product, constants.ERR_CODE_51, constants.ERR_CODE_51_MSG
	// }
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// UpdateStockAndHppProductByID ...
func UpdateStockAndHppProductByID(prodID, warehouseID int64, qty int64, newHpp float32) (dbmodels.Stock, string, string) {

	fmt.Println("Update stock", prodID, "qty ", qty)

	db := GetDbCon()
	db.Debug().LogMode(true)

	// var product dbmodels.Product
	// db.Model(&dbmodels.Product{}).Where("id=?", prodID).First(&product)
	var stock dbmodels.Stock
	db.Where("product_id=? and warehouse_id", prodID, warehouseID).First(&stock)
	stock.Qty = qty
	db.Save(&stock)

	var product dbmodels.Product
	db.Where("product_id=? ", prodID).First(&product)
	product.Hpp = newHpp
	db.Save(&product)

	return stock, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}
