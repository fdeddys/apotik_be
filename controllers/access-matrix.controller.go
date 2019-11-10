package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oasis-be/models"
	dbmodels "oasis-be/models/dbModels"
	"oasis-be/models/dto"
	"oasis-be/services"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AccessMatrixController ...
type AccessMatrixController struct {
	DB *gorm.DB
}

// RoleMenuService ...
var RoleMenuService = new(services.RoleMenuService)

// GetAllActiveMenu ...
func (h *AccessMatrixController) GetAllActiveMenu(c *gin.Context) {
	res := []dbmodels.Menu{}

	res = RoleMenuService.GetActiveMenu()

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// GetMenuByRoleID ...
func (h *AccessMatrixController) GetMenuByRoleID(c *gin.Context) {
	res := []dto.RoleMenuDto{}

	roleID, errPage := strconv.Atoi(c.Param("roleId"))
	if errPage != nil {
		logs.Info("error", errPage)
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = RoleMenuService.GetMenuByRole(roleID)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// SaveRoleMenu ...
func (h *AccessMatrixController) SaveRoleMenu(c *gin.Context) {
	req := []int{}
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

	roleID, err := strconv.Atoi(c.Param("roleId"))
	if err != nil {
		logs.Info("error", err)
		res.ErrCode = "03"
		res.ErrDesc = "Error, unmarshall body Request"
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = RoleMenuService.SaveMenuByRole(roleID, req)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
