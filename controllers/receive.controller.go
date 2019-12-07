package controllers

import (
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	"distribution-system-be/models/dto"
	"distribution-system-be/services"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ReceiveController ...
type ReceiveController struct {
	DB *gorm.DB
}

var receiveService = new(services.ReceiveService)

// FilterData ...
func (s *ReceiveController) FilterData(c *gin.Context) {
	req := dto.FilterReceive{}
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

	status := -1
	if intVal, errconv := strconv.Atoi(req.Status); errconv == nil {
		status = intVal
	}

	res = receiveService.GetDataPage(req, page, count, status)

	c.JSON(http.StatusOK, res)

	return
}

// GetByOrderId ...
func (s *ReceiveController) GetByReceiveId(c *gin.Context) {

	res := dbmodels.Receive{}

	orderID, errPage := strconv.ParseInt(c.Param("id"), 10, 64)
	if errPage != nil {
		logs.Info("error", errPage)
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = receiveService.GetDataReceiveByID(orderID)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return

}
