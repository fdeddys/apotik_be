package controllers

import (
	"distribution-system-be/constants"
	"distribution-system-be/database"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"distribution-system-be/services"
	"distribution-system-be/utils/util"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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

//GetProductListPaging ...
func (h *ProductController) GetProductListPagingAllStatus(c *gin.Context) {
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

	res = ProductService.GetProductFilterPagingAllStatus(req, page, count)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

func (h *ProductController) SearchProduct(c *gin.Context) {
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

	res = ProductService.SearchProduct(req, page, count)

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

func (h *ProductController) ProcessCSV(c *gin.Context) {

	fmt.Println("Process")
	fileObat, err := os.Open("obat.csv")
	if err != nil {
		fmt.Println("Error ==>", err)
		panic(err)
	}
	defer fileObat.Close()

	fmt.Println("reader")
	csvReader := csv.NewReader(fileObat)
	csvReader.Comma = ';'
	records, err2 := csvReader.ReadAll()
	if err2 != nil {
		fmt.Println("Error err2 ==>", err2)
		return
	}

	var datas []dto.TemplateObat
	for _, value := range records {
		fmt.Println("data ", value)
		var template dto.TemplateObat
		template.Plu = value[0]
		template.Name = value[1]
		template.Satuan = value[2]
		template.Qty = util.Atoi64(value[3])
		template.Hargabeli = math.Floor(util.AtoFloat64(value[7]) / util.AtoFloat64(value[3]))
		template.Hargajual = util.Atoi64(value[6])
		template.SatuanBesar = value[4]
		template.Kadar = value[8]
		datas = append(datas, template)

	}

	for _, templateObat := range datas {
		fmt.Println(templateObat)
		var produk dbmodels.Product
		var uomID int64
		lookup, errcode, _, _ := database.GetLookupByName(templateObat.Satuan)
		if errcode != constants.ERR_CODE_00 {
			uomID = 30
		} else {
			uomID = lookup.ID
		}

		var uomID2 int64
		lookup2, errcode2, _, _ := database.GetLookupByName(templateObat.SatuanBesar)
		if errcode2 != constants.ERR_CODE_00 {
			uomID2 = 30
		} else {
			uomID2 = lookup2.ID
		}

		produk.BigUomID = uomID
		produk.BrandID = 1
		produk.Code = ""
		produk.Hpp = float32(templateObat.Hargabeli)
		produk.LastUpdate = time.Now()
		produk.LastUpdateBy = "system"
		produk.Name = templateObat.Name
		produk.ProductGroupID = 1
		produk.PLU = templateObat.Plu
		produk.QtyUom = int16(templateObat.Qty)
		produk.SellPrice = float32(templateObat.Hargajual)
		produk.SellPriceType = 0
		produk.SmallUomID = uomID
		produk.BigUomID = uomID2
		produk.SediaanID = 35
		produk.Status = 1
		produk.Composition = templateObat.Kadar

		database.SaveProduct(produk)
	}

	c.JSON(http.StatusOK, "ok")
	c.Abort()
	return
}

func (h *ProductController) ProcessUpdateProd(c *gin.Context) {
	res := database.ProcessTemplateToProduct()
	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// UploadCSV ...
func (h *ProductController) UploadCSV(c *gin.Context) {
	var res models.ContentResponse
	file, err := c.FormFile("file")
	if err != nil {
		res.ErrCode = "06"
		res.ErrDesc = "Failed to upload file: " + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	tempFilePath := "update-product.csv"
	err = c.SaveUploadedFile(file, tempFilePath)
	if err != nil {
		res.ErrCode = "06"
		res.ErrDesc = "Failed to save file on server: " + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	fileObat, err := os.Open(tempFilePath)
	if err != nil {
		res.ErrCode = "06"
		res.ErrDesc = "Failed to open saved file: " + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	defer fileObat.Close()

	csvReader := csv.NewReader(fileObat)
	csvReader.Comma = ';'
	
	// Auto-detect delimiter
	firstLineBytes, errRead := ioutil.ReadFile(tempFilePath)
	if errRead == nil && len(firstLineBytes) > 0 {
		snippet := string(firstLineBytes)
		if len(snippet) > 1000 {
			snippet = snippet[:1000]
		}
		semicolonCount := strings.Count(snippet, ";")
		commaCount := strings.Count(snippet, ",")
		if commaCount > semicolonCount {
			csvReader.Comma = ','
		}
	}

	records, err2 := csvReader.ReadAll()
	if err2 != nil {
		res.ErrCode = "06"
		res.ErrDesc = "Failed to read CSV: " + err2.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	clearRes := database.ClearTemplateProducts()
	if clearRes.ErrCode != "00" {
		res.ErrCode = clearRes.ErrCode
		res.ErrDesc = clearRes.ErrDesc
		c.JSON(http.StatusOK, res)
		return
	}

	for i, value := range records {
		var code, name, plu, priceStr, statusStr string

		if len(value) >= 9 {
			// Layout A: Full Export Layout (10 columns)
			// No;Name;Code;Big UOM;Small UOM;Qty UOM;Status;Sell Price;PLU;Composition
			name = strings.TrimSpace(value[1])
			code = strings.TrimSpace(value[2])
			statusStr = strings.TrimSpace(value[6])
			priceStr = strings.TrimSpace(value[7])
			plu = strings.TrimSpace(value[8])
		} else if len(value) >= 5 {
			// Layout B: Short Layout (5 columns)
			// Code;Name;PLU;HargaJual;Status
			code = strings.TrimSpace(value[0])
			name = strings.TrimSpace(value[1])
			plu = strings.TrimSpace(value[2])
			priceStr = strings.TrimSpace(value[3])
			statusStr = strings.TrimSpace(value[4])
		} else {
			continue
		}

		sellPrice := util.AtoFloat32(priceStr)
		if i == 0 && priceStr != "" {
			_, parseErr := strconv.ParseFloat(priceStr, 32)
			if parseErr != nil {
				// Skip header row
				continue
			}
		}

		status := 1
		if statusStr != "" {
			if strings.ToLower(statusStr) == "inactive" || statusStr == "0" || strings.ToLower(statusStr) == "tidak aktif" {
				status = 0
			}
		}

		var tp dbmodels.TemplateProduct
		tp.Code = code
		tp.Plu = plu
		tp.Nama = name
		tp.HargaJual = sellPrice
		tp.Status = status

		errSave := database.SaveTemplateProduct(tp)
		if errSave != nil {
			fmt.Println("Error saving to template_product:", errSave)
		}
	}

	list, errList := database.GetTemplateProducts()
	if errList != nil {
		res.ErrCode = "02"
		res.ErrDesc = "CSV processed, but failed to fetch display list: " + errList.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success upload and import CSV"
	res.Contents = list
	c.JSON(http.StatusOK, res)
}

// GetTemplateProducts ...
func (h *ProductController) GetTemplateProducts(c *gin.Context) {
	var res models.ContentResponse
	list, err := database.GetTemplateProducts()
	if err != nil {
		res.ErrCode = "02"
		res.ErrDesc = "Failed to fetch template products: " + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	res.ErrCode = "00"
	res.ErrDesc = "Success"
	res.Contents = list
	c.JSON(http.StatusOK, res)
}

// ClearTemplate ...
func (h *ProductController) ClearTemplate(c *gin.Context) {
	res := database.ClearTemplateProducts()
	c.JSON(http.StatusOK, res)
}
