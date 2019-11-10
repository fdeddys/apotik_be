package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oasis-be/models"
	"oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"oasis-be/services"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"oasis-be/constants"
	"github.com/astaxie/beego"
	"bytes"
	"time"
)


// NooDocController ...
type NooDocController struct {
	DB *gorm.DB
}


var NooDocService = new(services.NooDocService)


func (s *NooDocController) FilterDataNooDoc(c *gin.Context) {
	req := dto.FilterSupplierNooDocDto{}
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
	log.Println("searchCode-->", string(temp))
	res = NooDocService.GetDataNooDocPaging(req, page, count)

	c.JSON(http.StatusOK, res)

	return
}


func (s *NooDocController) UpdateDataNooDoc(c *gin.Context){
	nooDocReq   := dbmodels.SupplierNooDoc{}
	res := models.Response{}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &nooDocReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	if nooDocReq.ApprovalStatus == "1" {
		res := SupplierNooDocService.UpdateDataSupplierNooDoc(&nooDocReq)
		if res.ErrCode == constants.ERR_CODE_00 {
			url := beego.AppConfig.DefaultString("kafka.rest-server", "")
			var jsonStr = []byte(fmt.Sprintf(`{"topic":"oasis-add-merchant-uki-topic","data":{"access_token":"%s", "name":"%s", "state_action":"activate", "code":"%s"}}`, 
			"", nooDocReq.Merchant.Name, nooDocReq.Merchant.Code))
			req, err:= http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
			if err != nil {
				fmt.Println(err)
				return
			}
	
			req.Header.Set("Timestamp", fmt.Sprintf("%d",time.Now().Unix()))
			req.Header.Set("Content-Type", "application/json")
	
			client := &http.Client{}
			resp, err := client.Do(req)
			bufRespNooDocReq, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(bufRespNooDocReq))
		}else{
			fmt.Println("Failed updated")
			res.ErrCode = constants.ERR_CODE_03
			res.ErrDesc = constants.ERR_CODE_03_MSG
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG
	
	// fmt.Println("Edit data")
	c.JSON(http.StatusOK, res)
	return
}


func (s *NooDocController) GetStatusApproveNooByMerchantAndSupplier(c *gin.Context){
	var res models.ResponseStatus
	var reqSupplierNooDoc dbmodels.ReqNooDocByStatus
	var status bool

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &reqSupplierNooDoc); err != nil {
		fmt.Println("Error, body Request ")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	resp := SupplierNooDocService.GetStatusApproveByMerchantAndSupplier(reqSupplierNooDoc.SupplierID, reqSupplierNooDoc.MerchantCode)
	if resp.ErrCode == constants.ERR_CODE_00 {
		d := resp.Contents.([]dbmodels.SupplierNooDoc)
		for i:=0;i<len(d);i++{
			if d[i].ApprovalStatus == "1" {
				status = true
			}else{
				status = false
			}
		}
		res.ErrCode 		= constants.ERR_CODE_00
		res.ErrDesc 		= constants.ERR_CODE_00_MSG
		res.ApprovalStatus  = status
	}else{
		res.ErrCode 		= constants.ERR_CODE_40
		res.ErrDesc 		= constants.ERR_CODE_40_MSG
		res.ApprovalStatus  = status
	}

	c.JSON(http.StatusOK, res)
	return
}