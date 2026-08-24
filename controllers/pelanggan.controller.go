package controllers

import (
	"distribution-system-be/models"
	dto "distribution-system-be/models/dto"
	"distribution-system-be/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"distribution-system-be/constants"
	dbmodels "distribution-system-be/models/dbModels"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// PelangganController ...
type PelangganController struct {
	DB *gorm.DB
}

// PelangganService ...
var PelangganService = new(services.PelangganService)

// SaveDataPelanggan ...
func (m *PelangganController) SaveDataPelanggan(c *gin.Context) {
	pelangganReq := dbmodels.Pelanggan{}
	res := models.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &pelangganReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, PelangganService.SavePelanggan(&pelangganReq))
	return
}

// FilterDataPelanggan ...
func (m *PelangganController) FilterDataPelanggan(c *gin.Context) {
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

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = PelangganService.GetPelangganPaging(req, page, count)
	c.JSON(http.StatusOK, res)
	return
}

// DeleteDataPelanggan ...
func (m *PelangganController) DeleteDataPelanggan(c *gin.Context) {
	res := models.Response{}

	id, errID := strconv.ParseInt(c.Param("id"), 10, 64)
	if errID != nil {
		logs.Info("error", errID)
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = errID.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = PelangganService.DeletePelanggan(id)
	c.JSON(http.StatusOK, res)
	return
}
