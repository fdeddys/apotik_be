package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oasis-be/constants"
	"oasis-be/models"
	"oasis-be/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// VendorController ...
type VendorController struct {
	DB *gorm.DB
}

// VendorService ...
var VendorService = new(services.VendorService)

// Login ...
func (h *VendorController) Login(c *gin.Context) {

	res := VendorService.Login()
	c.JSON(http.StatusOK, res)
}

// GetSo ...
func (h *VendorController) GetSo(c *gin.Context) {

	res := VendorService.GetSalesOrder()
	c.JSON(http.StatusOK, res)
}

// UpdateStatus ...
func (h *VendorController) UpdateStatus(c *gin.Context) {

	fmt.Println("Update status from UKIrame ")
	fmt.Println("======================== ")
	req := models.RequestUpdateStatus{}
	res := models.ContentResponse{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	err := json.Unmarshal(dataBodyReq, &req)
	if err != nil {
		fmt.Println("Error, body Request ")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, VendorService.UpdateStatus(req))
	fmt.Println("============================================ ")
}
