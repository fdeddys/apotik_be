package controllers

import (
	"distribution-system-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// OrderController ...
type ParameterController struct {
	DB *gorm.DB
}

// ParameterService ...
var ParameterService = new(services.ParameterService)

// GetByName ...
func (s *ParameterController) GetByName(c *gin.Context) {

	// res := dbmodels.Parameter{}

	paramName := c.Param("param-name")
	res := ParameterService.GetByName(paramName)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return

}
