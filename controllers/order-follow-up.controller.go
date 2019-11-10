package controllers

import (
	"net/http"
	models "oasis-be/models"
	dto "oasis-be/models/dto"
	"oasis-be/services"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type FollowUpOrderController struct {
	DB *gorm.DB
}

var FollowUpOrderService = new(services.FollowUpOrderService)

func (f *FollowUpOrderController) GetFollowUpOrder(c *gin.Context) {
	req := dto.FilterOrder{}
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
	log.Println("is release ", req.InternalStatus)

	intStatus := 0
	if intVal, errconv := strconv.Atoi(req.InternalStatus); errconv == nil {
		intStatus = intVal
	}
	res = FollowUpOrderService.GetFollowOrder(req, page, count, intStatus)

	c.JSON(http.StatusOK, res)

	return
}
