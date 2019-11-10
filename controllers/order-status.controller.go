package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"oasis-be/models"
	dto "oasis-be/models/dto"
	"oasis-be/services"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// OrderStatusController ...
type OrderStatusController struct {
	DB *gorm.DB
}

var orderStatusService = new(services.OrderStatusService)

// GetListStatus ...
func (s *OrderStatusController) GetListStatus(c *gin.Context) {

	req := dto.FilterOrderDetail{}
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

	log.Println("is release ", req)

	res = orderStatusService.GetDataOrderStatusPage(req, page, count)
	// orderStatusService.GetDataOrderStatusPage(req, page, count)

	c.JSON(http.StatusOK, res)

	return
}
