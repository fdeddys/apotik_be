package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oasis-be/models"
	dto "oasis-be/models/dto"
	"oasis-be/services"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"oasis-be/models/dbModels"
	"log"
	"oasis-be/constants"
)


// MerchantController ...
type MerchantController struct {
	DB *gorm.DB
}


// MerchantService ...
var MerchantService = new(services.MerchantService)


// Save Data Merchant
func (m *MerchantController) SaveDataMerchant(c *gin.Context) {
	MerchantReq := dbmodels.Merchant{} 
	res := models.Response{}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &MerchantReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	c.JSON(http.StatusOK,MerchantService.SaveDataMerchant(&MerchantReq))
	
	return 
}


// List and Paging Merchant
func (m *MerchantController) FilterDataMerchant(c *gin.Context) {
	req := dto.FilterName{}
	res := models.ResponsePagination{}

	page, errPage := strconv.Atoi(c.Param("page"))
	if errPage != nil {
		logs.Info("error", errPage)
		res.Error = errPage.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	count, errCount := strconv.Atoi(c.Param("count"))
	if errCount != nil {
		logs.Info("error", errPage)
		res.Error = errCount.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	log.Println("page->", page, "count->", count)

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	temp, _ := json.Marshal(req)
	log.Println("searchName-->", string(temp))
	res = MerchantService.GetDataMerchantPaging(req, page, count)

	c.JSON(http.StatusOK, res)
	return
}

// Edit Data Merchant
func (m *MerchantController) EditDataMerchant(c *gin.Context) {
	req := dbmodels.Merchant{}
	res := models.Response{}
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	fmt.Println("Edit data")
	c.JSON(http.StatusOK, MerchantService.UpdateDataMerchant(&req))
}


// List Data Merchant
func (m *MerchantController) ListDataMerchant(c *gin.Context) {
	res := models.ResponseMerchant{}
	res = MerchantService.GetDataMerchantList()

	c.JSON(http.StatusOK, res)
	return
}


func (m *MerchantController) ListDataMerchantByName(c *gin.Context) {
	res := models.ContentResponse{}

	name := c.Query("search")
	if name == "" {
		logs.Info("error", "can't found the name string")
		c.JSON(http.StatusOK, res)
		c.Abort()
		return
	}

	res = MerchantService.GetDataMerchantListByName(name)
	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}


// func (m *MerchantController) CheckOrderMerchantSupplier(c *gin.Context) {
// 	var checkSupplierReq dbmodels.MerchantCheckSupplier
// 	var res models.ResponseSFA
// 	var data string

// 	body := c.Request.Body
// 	dataBodyReq, _ := ioutil.ReadAll(body)
	
// 	if err := json.Unmarshal(dataBodyReq, &checkSupplierReq); err != nil {
// 		fmt.Println("Error, body Request")
// 		res.Data = ""
// 		res.Meta.Status = false
// 		res.Meta.Code = 400
// 		res.Meta.Message = "Terjadi Kesalahan"
// 		c.JSON(http.StatusBadRequest, res)
// 		c.Abort()
// 		return
// 	} 

// 	checkSupplier := MerchantService.GetDataCheckOrder(checkSupplierReq.SupplierID, checkSupplierReq.MerchantID)
	
// 	if len(checkSupplier) > 0 {
// 		data = "Data tersedia"
// 	}else {
// 		data = "Data tidak tersedia"
// 	}

// 	res.Data = data
// 	res.Meta.Status = true
// 	res.Meta.Code = 200
// 	res.Meta.Message = "OK"
	
// 	c.JSON(http.StatusOK, res)
// 	return
// }
