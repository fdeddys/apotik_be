package controllers

import (
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"distribution-system-be/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//ProductController ...
type ProductController struct {
	DB *gorm.DB
}

//ProductService ...
var ProductService = new(services.ProductService)

//GetProductListPaging ...
func (h *ProductController) GetProductListPaging(c *gin.Context) {
	req := dto.FilterProduct{}
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

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = ProductService.GetProductFilterPaging(req, page, count)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// GetProductDetails ...
func (h *ProductController) GetProductDetails(c *gin.Context) {
	res := models.ContentResponse{}

	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		logs.Info("error", errID)
		// res.Error = errID.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = ProductService.GetProductDetails(id)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// SaveProduct ...
func (h *ProductController) SaveProduct(c *gin.Context) {

	req := dbmodels.Product{}
	res := models.NoContentResponse{}

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

	c.JSON(http.StatusOK, ProductService.SaveProduct(&req))
	return
}

// UpdateProduct ...
func (h *ProductController) UpdateProduct(c *gin.Context) {
	req := dbmodels.Product{}
	res := models.NoContentResponse{}

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

	c.JSON(http.StatusOK, ProductService.UpdateProduct(&req))
	return
}

func (h *ProductController) ProductList(c *gin.Context) {
	// res := []dbmodels.Product{}

	c.JSON(http.StatusOK, ProductService.ProductList())
	return
}

//GetProductLike ...
func (h *ProductController) GetProductLike(c *gin.Context) {
	res := models.ContentResponse{}

	productterms := c.Query("terms")

	if productterms == "" {
		logs.Info("error", "can't found the brand string")
		c.JSON(http.StatusOK, res)
		c.Abort()
		return
	}

	// fmt.Sprintf("ini lookupstr = " + lookupstr)

	res = ProductService.GetProductLike(productterms)
	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
