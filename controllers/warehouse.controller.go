package controllers

import (
	"distribution-system-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// BrandController ...
type WarehouseController struct {
	DB *gorm.DB
}

// WarehouseService ...
var WarehouseService = new(services.WarehouseService)

// GetBrand ...
func (h *WarehouseController) GetWarehouse(c *gin.Context) {

	res := WarehouseService.GetAllWarehouse()

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
