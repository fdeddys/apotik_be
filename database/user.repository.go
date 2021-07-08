package database

import (
	"crypto/sha256"
	"distribution-system-be/constants"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"encoding/hex"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
)

// type TotalRows struct {
// 	Total int `gorm:"column(count)"`
// }

// GetUserTransaction ...
func GetUserTransaction(param dto.FilterUser, offset int, limit int) ([]dbmodels.User, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var user []dbmodels.User
	var total int

	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&user).Error
		if err != nil {
			return user, 0, err
		}
		return user, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerys(db, offset, limit, &user, param, errQuery)
	go AsyncQueryCounts(db, &total, param, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()
	// wg.Done()

	if resErrQuery != nil {
		return user, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return user, 0, resErrCount
	}
	return user, total, nil
}

// AsyncQueryCounts ...
func AsyncQueryCounts(db *gorm.DB, total *int, param dto.FilterUser, resChan chan error) {
	// var criteriaUserName = "%"
	// if strings.TrimSpace(param.Username) != "" {
	// 	criteriaUserName = param.Username + criteriaUserName
	// }

	criteriaUserName := param.Username
	if criteriaUserName == "" {
		criteriaUserName = "%"
	} else {
		criteriaUserName = "%" + param.Username + "%"
	}

	err := db.Preload("Role").Model(&dbmodels.User{}).Where("user_name ilike ?", criteriaUserName).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// AsyncQuerys ...
func AsyncQuerys(db *gorm.DB, offset int, limit int, user *[]dbmodels.User, param dto.FilterUser, resChan chan error) {

	// var criteriaUserName = "%"
	// if strings.TrimSpace(param.Username) != "" {
	// 	criteriaUserName = param.Username + criteriaUserName
	// }

	criteriaUserName := param.Username
	if criteriaUserName == "" {
		criteriaUserName = "%"
	} else {
		criteriaUserName = "%" + param.Username + "%"
	}

	err := db.Preload("Role").Order("user_name ASC").Offset(offset).Limit(limit).Find(&user, "user_name ilike ?", criteriaUserName).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// GetUserByName ...
func GetUserByName(username string) (dbmodels.User, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var user dbmodels.User
	var err error

	db.Where("user_name = ?", username).Find(&user)

	fmt.Println("User => ", user)
	return user, err

}

// GetByUsername ...
func GetByUsername(username string) (dbmodels.User, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var user dbmodels.User
	var err error

	db.Where("user_name = ?", username).First(&user)

	fmt.Println("User => ", user)
	return user, err

}

// UpdatePassword ...
func UpdatePassword(username, password string) (dbmodels.User, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var user dbmodels.User
	var err error

	db.Model(&user).Where("user_name = ?", username).Update("password", password)

	fmt.Println("User => ", user)
	return user, err

}

// SaveUser ...
func SaveUser(user dbmodels.User) models.ContentResponse {
	var res models.ContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	pass := GeneratePassword(8)
	fmt.Println("password =>", pass)
	enc := Encrypt(pass)
	fmt.Printf("encrypt password =>%s\n", enc)

	dec := Decrypt(enc)
	fmt.Printf("decrypt password =>%s\n", dec)
	hashPassword := sha256.Sum256([]byte(user.UserName + pass))
	user.Password = hex.EncodeToString(hashPassword[:])

	fmt.Printf("hash password =>%s\n", hex.EncodeToString(hashPassword[:]))

	if r := db.Save(&user); r.Error != nil {
		res.ErrCode = constants.ERR_CODE_51
		res.ErrDesc = constants.ERR_CODE_51_MSG
		return res
	}

	// byt := []byte(`{"enc_pass":"` + enc + `"}`)
	// var dat map[string]interface{}
	// if err := json.Unmarshal(byt, &dat); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(dat)

	user.Password = enc
	// user.LastName = pass
	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG
	res.Contents = user
	return res
}

// UpdateUser ...
func UpdateUser(updateduser dbmodels.User) models.NoContentResponse {
	var res models.NoContentResponse
	db := GetDbCon()
	db.Debug().LogMode(true)

	var user dbmodels.User
	err := db.Model(&dbmodels.User{}).Where("id=?", &updateduser.ID).First(&user).Error
	if err != nil {
		res.ErrCode = constants.ERR_CODE_51
		res.ErrDesc = constants.ERR_CODE_51_MSG
		return res
	}

	user.UserName = updateduser.UserName
	user.Email = updateduser.Email
	user.LastUpdateBy = updateduser.LastUpdateBy
	user.LastUpdate = updateduser.LastUpdate
	user.SupplierCode = updateduser.SupplierCode
	user.FirstName = updateduser.FirstName
	user.LastName = updateduser.LastName
	user.IsAdmin = updateduser.IsAdmin
	user.Status = updateduser.Status
	user.RoleID = updateduser.RoleID

	err2 := db.Save(&user)
	if err2 != nil {
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		return res
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	return res
}
