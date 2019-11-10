package database

import (
	// "log"
	"oasis-be/models"
	"oasis-be/models/dbModels"
	// dto "oasis-be/models/dto"
	_ "strconv"
	// "strings"
	// "sync"

	// "github.com/jinzhu/gorm"
	"oasis-be/constants"
	// "fmt"
	// "strings"
)


func SaveMerchantPicture(merchantPict *dbmodels.MerchantPict) models.Response {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var res models.Response

	if r := db.Save(&merchantPict); r.Error != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed save data to DB"
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}


func GetMerchantPictureByCode(merchant_code string)([]dbmodels.MerchantPict) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplier []dbmodels.MerchantPict
	var err error

	err = db.Where("pict_path ilike ?", merchant_code + "%").Find(&supplier).Error
	if err != nil {
		return supplier
	}

	// lookup, _, _, _ := GetLookupByGroup("NOO_DOCUMENT")
	// for i := 0; i < len(supplier); i++ {
	// 	// if supplier[i].LookupCode == "001" {
	// 	// 	isExist := CheckImage("ktp/" + supplier[i].PictPath, "supplier")
	// 	// 	if isExist {
	// 	// 		supplier[i].PictPath = GetImage("ktp/" + supplier[i].PictPath, "supplier")
	// 	// 	} else {
	// 	// 		supplier[i].PictPath = GetImage("no_image", "supplier")
	// 	// 	}
	// 	// }else if supplier[i].LookupCode == "002"{
	// 	// 	isExist := CheckImage("npwp/" + supplier[i].PictPath, "supplier")
	// 	// 	if isExist {
	// 	// 		supplier[i].PictPath = GetImage("npwp/" + supplier[i].PictPath, "supplier")
	// 	// 	} else {
	// 	// 		supplier[i].PictPath = GetImage("no_image", "supplier")
	// 	// 	}
	// 	// }
	// 	if supplier[i].LookupCode == lookup[i].Code {
	// 		// splitName := strings.Split(lookup[i].Name, " ")
	// 		fmt.Println("merchant supplier:",supplier[i].LookupCode)
	// 	}
	// }

	return supplier
}

func GetMerchantPictureByLoookupAndCode(lookup_code string, merchant_code string)([]dbmodels.MerchantPict){
	db := GetDbCon()
	db.Debug().LogMode(true)

	var supplier []dbmodels.MerchantPict
	var err error

	err = db.Where("pict_path ilike ? and lookup_code = ?", merchant_code+"%", lookup_code).Find(&supplier).Error
	if err != nil {
		return supplier
	}
	return supplier
}