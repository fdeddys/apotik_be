package database

import (
	"distribution-system-be/constants"
	dbmodels "distribution-system-be/models/dbModels"
)

// GetStockByProductAndWarehouse ...
func GetStockByProductAndWarehouse(productID, warehouseID int64) (dbmodels.Stock, string, string) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var stock dbmodels.Stock
	err := db.Where("warehouse_id = ? and product_id = ? ", warehouseID, productID).First(&stock).Error

	if err != nil {
		return stock, constants.ERR_CODE_81, constants.ERR_CODE_81_MSG + " " + err.Error()
	}
	return stock, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG

}
