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

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//SupplierGroupController ...
type SupplierGroupController struct {
	DB *gorm.DB
}

//SupplierGroupService ...
var SupplierGroupService = new(services.SupplierGroupService)

//GetSupplierGroupPaging ...
func (h *SupplierGroupController) GetSupplierGroupPaging(c *gin.Context) {
	req := dto.FilterSupplierGroup{}
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

	res = SupplierGroupService.GetSupplierGroupPaging(req, page, count)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// GetSupplierGroupDetails ...
func (h *SupplierGroupController) GetSupplierGroupDetails(c *gin.Context) {
	res := models.ContentResponse{}

	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		logs.Info("error", errID)
		// res.Error = errID.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = SupplierGroupService.GetSupplierGroupDetails(id)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// SaveSupplierGroup ...
func (h *SupplierGroupController) SaveSupplierGroup(c *gin.Context) {

	req := dbmodels.SupplierGroup{}
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

	c.JSON(http.StatusOK, SupplierGroupService.SaveSupplierGroup(&req))
	return
}

// UpdateSupplierGroup ...
func (h *SupplierGroupController) UpdateSupplierGroup(c *gin.Context) {
	req := dbmodels.SupplierGroup{}
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

	c.JSON(http.StatusOK, SupplierGroupService.UpdateSupplierGroup(&req))
	return
}
// GetListSupplierGroup
func (h *SupplierGroupController) GetListSupplierGroup(c *gin.Context) {
	res := models.ResponseSupplierGroup{}
	res = SupplierGroupService.GetListSupplierGroup()

	c.JSON(http.StatusOK, res)
	return
}