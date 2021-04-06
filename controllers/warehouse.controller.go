package controllers

import (
	"distribution-system-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// WarehouseController ...
type WarehouseController struct {
	DB *gorm.DB
}

// WarehouseService ...
var WarehouseService = new(services.WarehouseService)

// GetWarehouse ...
func (h *WarehouseController) GetWarehouse(c *gin.Context) {

	res := WarehouseService.GetAllWarehouse()

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
