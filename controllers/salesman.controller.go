package controllers

import (
	"distribution-system-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SalesmanController ...
type SalesmanController struct {
	DB *gorm.DB
}

// SalesmanService ...
var SalesmanService = new(services.SalesmanService)

// SalesmanController ...
func (h *SalesmanController) GetSalesman(c *gin.Context) {

	res := SalesmanService.GetAllSalesman()

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
