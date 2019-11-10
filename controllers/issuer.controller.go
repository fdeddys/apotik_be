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

	"log"
	"oasis-be/models/dbModels"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IssuerController ...
type IssuerController struct {
	DB *gorm.DB
}

// IssuerService ...
var IssuerService = new(services.IssuerService)

// Save Data Issuer
func (i *IssuerController) SaveDataIssuer(c *gin.Context) {
	IssuerReq := dbmodels.Issuer{}
	res := models.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &IssuerReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = "03"
		res.ErrDesc = "Error, unmarshall body Request"
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, IssuerService.SaveDataIssuer(&IssuerReq))

	return
}

// List and Paging Issuer
func (i *IssuerController) FilterDataIssuer(c *gin.Context) {
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
	res = IssuerService.GetDataIssuerPaging(req, page, count)

	c.JSON(http.StatusOK, res)

	return
}

// Edit Data Issuer
func (i *IssuerController) EditDataIssuer(c *gin.Context) {
	req := dbmodels.Issuer{}
	res := models.Response{}
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

	fmt.Println("Edit data")
	c.JSON(http.StatusOK, IssuerService.UpdateDataIssuer(&req))
}

func (i *IssuerController) ListDataIssuer(c *gin.Context) {
	res := models.ResponseIssuer{}
	res = IssuerService.GetDataIssuerList()

	c.JSON(http.StatusOK, res)
	return
}


func (i *IssuerController) ListDataIssuerByName(c *gin.Context) {
	res := models.ContentResponse{}

	name := c.Query("search")
	if name == "" {
		logs.Info("error", "can't found the name string")
		c.JSON(http.StatusOK, res)
		c.Abort()
		return
	}

	res = IssuerService.GetDataIssuerListByName(name)
	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}



