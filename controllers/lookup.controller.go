package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	dto "oasis-be/models/dto"
	"oasis-be/services"
	"strconv"
	"strings"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"oasis-be/constants"
)

//LookupController ...
type LookupController struct {
	DB *gorm.DB
}

//LookupService ...
var lookupService = new(services.LookupService)

//GetLookupByGroup ...
func (h *LookupController) GetLookupByGroup(c *gin.Context) {
	res := models.ContentResponse{}

	lookupstr := strings.ToUpper(c.Query("terms"))

	if lookupstr == "" {
		logs.Info("error", "can't found the lookup string")
		c.JSON(http.StatusOK, res)
		c.Abort()
		return
	}

	fmt.Sprintf("ini lookupstr = " + lookupstr)

	res = lookupService.GetLookupByGroup(lookupstr)
	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// GetLookupPaging ...
func (h *LookupController) GetLookupPaging(c *gin.Context) {
	req := dto.FilterLookup{}
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

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = lookupService.GetPagingLookup(req, page, count)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// GetLookupFilter ...
func (h *LookupController) GetLookupFilter(c *gin.Context) {
	res := models.ContentResponse{}

	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		logs.Info("error", errID)
		// res.Error = errID.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = lookupService.GetLookupFilter(id)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// SaveLookup ...
func (h *LookupController) SaveLookup(c *gin.Context) {

	req := dbmodels.Lookup{}
	res := models.NoContentResponse{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.ErrCode = "03"
		res.ErrDesc = "Error, unmarshall body Request"
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, lookupService.SaveLookup(&req))
	return
}


// GetDistinctLookup ...
func (h *LookupController) GetDistinctLookup(c *gin.Context) {
	res := models.ContentResponse{}
	res = lookupService.GetDistinctLookup()

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// UpdateLookup ...
func (h *LookupController) UpdateLookup(c *gin.Context) {
	req := dbmodels.Lookup{}
	res := models.NoContentResponse{}

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

	c.JSON(http.StatusOK, lookupService.UpdateLookup(&req))
	return
}