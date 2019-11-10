package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oasis-be/models"
	dto "oasis-be/models/dto"
	"oasis-be/services"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"oasis-be/models/dbModels"
	"oasis-be/constants"
	"bytes"
    "encoding/base64"

    "image"
    _ "image/gif"
    _ "image/jpeg"
	_ "image/png"
	"os"
    "log"
	"strings"
	"time"
	"github.com/astaxie/beego"
)


// SupplierController ...
type SupplierController struct {
	DB *gorm.DB
}


// SupplierService ...
var SupplierService = new(services.SupplierService)
var SupplierMerchantService = new(services.SupplierMerchantService)
var SupplierWarehouseService = new(services.SupplierWarehouseService)
var SupplierPriceService = new(services.SupplierPriceService)
var SupplierNooChecklistService = new(services.SupplierNooChecklistService)
var SupplierNooDocService = new(services.SupplierNooDocService)
var ProductServices = new(services.ProductService)
var MerchantServices = new(services.MerchantService)


/* ------------------------------------------- Begin Supplier ----------------------------------------------- */

func (s *SupplierController) SaveDataSupplier(c *gin.Context) {
	SupplierReq := dbmodels.Supplier{} 
	res := models.ResponseSupplier{}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &SupplierReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		res.Code    = ""
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	c.JSON(http.StatusOK,SupplierService.SaveDataSupplier(&SupplierReq))
	
	return 
}


func (s *SupplierController) FilterDataSupplier(c *gin.Context) {
	req := dto.FilterName{}
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
	res = SupplierService.GetDataSupplierPaging(req, page, count)

	c.JSON(http.StatusOK, res)

	return
}


func (s *SupplierController) EditDataSupplier(c *gin.Context) {
	req := dbmodels.Supplier{}
	res := models.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	fmt.Println("Edit data")
	c.JSON(http.StatusOK, SupplierService.UpdateDataSupplier(&req))
}

func (s *SupplierController) UploadImageSupplier(c *gin.Context) {
	file, header, err := c.Request.FormFile("logo")
	filename := header.Filename

	fmt.Printf(filename)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, SupplierService.UploadImageSupplier(file, filename))
	return
}


func (s *SupplierController) UploadImageKtp(c *gin.Context) {
	file, header, err := c.Request.FormFile("ktp")
	filename := header.Filename

	fmt.Printf(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	c.JSON(http.StatusOK, SupplierService.UploadImageSupplier(file, "ktp" + "/" + filename))
	return
}

func (s *SupplierController) UploadImageNpwp(c *gin.Context) {
	file, header, err := c.Request.FormFile("npwp")
	filename := header.Filename

	fmt.Printf(filename)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, SupplierService.UploadImageSupplier(file, "npwp" + "/" + filename))
	return
}

func (s *SupplierController) SaveDataMerchantPicture(c *gin.Context) {
	merchantPictureReq := dbmodels.MerchantPict{} 
	res := models.Response{}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &merchantPictureReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	c.JSON(http.StatusOK,SupplierService.SaveDataMerchantPicture(&merchantPictureReq))
	
	return 
}

func (s *SupplierController) ListMerchantPicture(c *gin.Context) {
	res := []dbmodels.MerchantPict{}
	var merchant_code string

	merchant_code = c.Params.ByName("merchant_code")
	if merchant_code == "" {
		logs.Info("error", "can't found the name string")
		c.JSON(http.StatusOK, res)
		c.Abort()
		return
	}

	res = SupplierService.GetMerchantPictures(merchant_code)
	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

/* ----------------------------------- End Supplier --------------------------------------------------------- */


/* --------------------------------- Begin Supplier Merchant ------------------------------------------------ */

func (s *SupplierController) SaveDataSupplierMerchant(c *gin.Context){
	supplierMerchantReq := dbmodels.SupplierMerchant{}

	res := models.Response{}

	id, err := strconv.ParseInt(c.Params.ByName("supplier_id"), 10, 64)

	if err != nil {
		fmt.Println("Error")
	}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierMerchantReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	supplierMerchantReq.SupplierId = id

	c.JSON(http.StatusOK,SupplierMerchantService.SaveDataSupplierMerchant(&supplierMerchantReq))
	
	return 
}


func (s *SupplierController) FilterDataSupplierMerchant(c *gin.Context) {
	req := dto.FilterSupplierMerchantDto{}
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

	supplier_id, errSupplierId := strconv.Atoi(c.Param("supplier_id"))
	if errSupplierId != nil {
		logs.Info("error", errPage)
		res.Error = errSupplierId.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	log.Println("page->", page, "count->", count, "supplier_id->", supplier_id)


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
	log.Println("searchCode-->", string(temp))
	res = SupplierMerchantService.GetDataSupplierMerchantPaging(req, page, count, int64(supplier_id))

	c.JSON(http.StatusOK, res)

	return
}

func (s *SupplierController) UpdateDataSupplierMerchant(c *gin.Context){
	supplierMerchantReq := dbmodels.SupplierMerchant{}

	res := models.Response{}

	id, err := strconv.ParseInt(c.Params.ByName("supplier_id"), 10, 64)

	if err != nil {
		fmt.Println("Error")
	}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierMerchantReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	supplierMerchantReq.SupplierId = id

	fmt.Println("Edit data")
	c.JSON(http.StatusOK, SupplierMerchantService.UpdateDataSupplierMerchant(&supplierMerchantReq))
}


func (s *SupplierController) ListDataSupplierById(c *gin.Context) {
	res := dbmodels.Supplier{}

	id, err := strconv.Atoi(c.Params.ByName("supplier_id"))

	if err != nil {
		fmt.Println("Error")
	}
	
	res = SupplierService.GetDataSupplierById(int64(id));

	c.JSON(http.StatusOK, res)
	return
}

/* ---------------------------------------- End Supplier Merchant ------------------------------------- */


/* ---------------------------------------- Begin Supplier Warehouse ---------------------------------- */

func (s *SupplierController) SaveDataSupplierWarehouse(c *gin.Context){
	var supplierWarehouse dbmodels.SupplierWarehouse
	var code int64
	var codeWarehouse string
	// var state_action string

	supplierWarehouseReq := dbmodels.SupplierWarehouse{}
	res := models.Response{}

	id, err := strconv.ParseInt(c.Params.ByName("supplier_id"), 10, 64)

	if err != nil {
		fmt.Println("Error")
	}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierWarehouseReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	supplierWarehouse = SupplierWarehouseService.GetDataLastSupplierWarehouse()

	if supplierWarehouse != (dbmodels.SupplierWarehouse{}){
		if supplierWarehouse.Code == "" {
			code = 1
		}else{
			codeWarehouse = strings.TrimPrefix(supplierWarehouse.Code, "SW")
			code, err = strconv.ParseInt(codeWarehouse, 10, 64) 
			code = code + 1
		}
	} else {
		code = 1
	}

	codeWarehouse = "SW" + fmt.Sprintf("%06d", code)

	supplierWarehouseReq.SupplierId = id
	supplierWarehouseReq.Code = codeWarehouse

	res = SupplierWarehouseService.SaveDataSupplierWarehouse(&supplierWarehouseReq)

	if res.ErrCode == "00" {
		url := beego.AppConfig.DefaultString("kafka.rest-server", "")
		var jsonStr = []byte(fmt.Sprintf(`{"topic":"oasis-add-warehouse-uki-topic","data" :{"access_token" : "", "name":"%s", "state_action":"activate", "code":"%s", "supplier_id":"%d"}}`, 
					supplierWarehouseReq.Description, supplierWarehouseReq.Code, supplierWarehouseReq.SupplierId))
		req, err:= http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		// fmt.Println("post:", req)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("Timestamp", fmt.Sprintf("%d",time.Now().Unix()))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		bufRespWarehouse, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bufRespWarehouse))
	}

	c.JSON(http.StatusOK, res)
	return 
}

func (s *SupplierController) FilterDataSupplierWarehouse(c *gin.Context) {
	req := dto.FilterSupplierWarehouseDto{}
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

	supplier_id, errSupplierId := strconv.Atoi(c.Param("supplier_id"))
	if errSupplierId != nil {
		logs.Info("error", errPage)
		res.Error = errSupplierId.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	log.Println("page->", page, "count->", count, "supplier_id->", supplier_id)


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
	log.Println("searchCode-->", string(temp))
	res = SupplierWarehouseService.GetDataSupplierWarehousePaging(req, page, count, int64(supplier_id))

	c.JSON(http.StatusOK, res)

	return
}

func (s *SupplierController) UpdateDataSupplierWarehouse(c *gin.Context){
	supplierWarehouseReq := dbmodels.SupplierWarehouse{}

	res := models.Response{}

	id, err := strconv.ParseInt(c.Params.ByName("supplier_id"), 10, 64)

	if err != nil {
		fmt.Println("Error")
	}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierWarehouseReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	supplierWarehouseReq.SupplierId = id

	fmt.Println("Edit data")
	c.JSON(http.StatusOK, SupplierWarehouseService.UpdateDataSupplierWarehouse(&supplierWarehouseReq))
}

func (s *SupplierController) GetDataWarehouseBySupplierId(c *gin.Context){
	req := dbmodels.SupplierWarehouseReq{}
	res := models.ResponseSupplierWarehouse{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	// temp, _ := json.Marshal(supplierWarehouseReq)
	res = SupplierWarehouseService.GetDataWarehouseBySupplier(req.SupplierCode)
	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG

	c.JSON(http.StatusOK, res)

	return
}

/* --------------------------------------- End Supplier Warehouse -------------------------------------- */


/* ---------------------------------------- Begin Supplier Price ---------------------------------- */

func (s *SupplierController) SaveDataSupplierPrice(c *gin.Context){
	supplierPriceReq := dbmodels.SupplierPrice{}
	res := models.Response{}

	id, err := strconv.ParseInt(c.Params.ByName("supplier_id"), 10, 64)

	if err != nil {
		fmt.Println("Error")
	}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierPriceReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	supplierPriceReq.SupplierId = id
	
	res = SupplierPriceService.SaveDataSupplierPrice(&supplierPriceReq)

	if res.ErrCode == constants.ERR_CODE_00 {
		url := beego.AppConfig.DefaultString("kafka.rest-server", "")
		var jsonStr = []byte(fmt.Sprintf(`{"topic":"oasis-add-product-uki-topic","data" :{"access_token" : "", "name":"%s", "unit_of_measurement_code": "PC", "state_action":"activate", "code":"%s", "supplier_id":"%d"}}`, 
					supplierPriceReq.Product.Name, supplierPriceReq.ProductCode, supplierPriceReq.SupplierId))
		req, err:= http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		// fmt.Println("post:", req)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("Timestamp", fmt.Sprintf("%d",time.Now().Unix()))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		bufRespSupplierProduct, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bufRespSupplierProduct))
	}

	c.JSON(http.StatusOK,res)
	return 
}

func (s *SupplierController) FilterDataSupplierPrice(c *gin.Context) {
	req := dto.FilterSupplierPriceDto{}
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

	supplier_id, errSupplierId := strconv.Atoi(c.Param("supplier_id"))
	if errSupplierId != nil {
		logs.Info("error", errPage)
		res.Error = errSupplierId.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	log.Println("page->", page, "count->", count, "supplier_id->", supplier_id)


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
	log.Println("searchCode-->", string(temp))
	res = SupplierPriceService.GetDataSupplierPricePaging(req, page, count, int64(supplier_id))

	c.JSON(http.StatusOK, res)

	return
}

func (s *SupplierController) UpdateDataSupplierPrice(c *gin.Context){
	supplierPriceReq := dbmodels.SupplierPrice{}

	res := models.Response{}

	id, err := strconv.ParseInt(c.Params.ByName("supplier_id"), 10, 64)

	if err != nil {
		fmt.Println("Error")
	}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierPriceReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	supplierPriceReq.SupplierId = id

	fmt.Println("Edit data")
	c.JSON(http.StatusOK, SupplierPriceService.UpdateDataSupplierPrice(&supplierPriceReq))
}

/* --------------------------------------- End Supplier Price -------------------------------------- */


/* --------------------------------------- Supplier Noo Checklist ---------------------------------- */
func (s *SupplierController) FilterDataSupplierNooChecklist(c *gin.Context) {
	req := dto.FilterSupplierNooChecklistDto{}
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

	supplier_id, errSupplierId := strconv.Atoi(c.Param("supplier_id"))
	if errSupplierId != nil {
		logs.Info("error", errPage)
		res.Error = errSupplierId.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	log.Println("page->", page, "count->", count, "supplier_id->", supplier_id)


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
	res = SupplierNooChecklistService.GetDataSupplierNooChecklistPaging(req, page, count, int64(supplier_id))
	fmt.Println(res)
	c.JSON(http.StatusOK, res)

	return
}

func (s *SupplierController) SaveDataSupplierNooChecklist(c *gin.Context){
	supplierNooChecklistReq := dbmodels.SupplierNooChecklist{}
	res := models.Response{}

	id, err := strconv.ParseInt(c.Params.ByName("supplier_id"), 10, 64)

	if err != nil {
		fmt.Println("Error")
	}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierNooChecklistReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	supplierNooChecklistReq.SupplierID = id

	c.JSON(http.StatusOK,SupplierNooChecklistService.SaveDataSupplierNooChecklist(&supplierNooChecklistReq))
	
	return 
}

func (s *SupplierController) UpdateDataSupplierNooChecklist(c *gin.Context){
	supplierNooChecklistReq := dbmodels.SupplierNooChecklist{}

	res := models.Response{}

	id, err := strconv.ParseInt(c.Params.ByName("supplier_id"), 10, 64)

	if err != nil {
		fmt.Println("Error")
	}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierNooChecklistReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	supplierNooChecklistReq.SupplierID = id

	fmt.Println("Edit data")
	c.JSON(http.StatusOK, SupplierNooChecklistService.UpdateDataSupplierNooChecklist(&supplierNooChecklistReq))
}
/* --------------------------------------- Supplier Noo Checklist ---------------------------------- */



/* --------------------------------------- Supplier Noo Doc ---------------------------------- */
func (s *SupplierController) SaveDataSupplierNooDoc(c *gin.Context){
	supplierNooDocReq := dbmodels.SupplierNooDoc{}
	res := models.Response{}

	id, err := strconv.ParseInt(c.Params.ByName("supplier_id"), 10, 64)

	if err != nil {
		fmt.Println("Error")
	}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierNooDocReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	supplierNooDocReq.SupplierID = id

	c.JSON(http.StatusOK,SupplierNooDocService.SaveDataSupplierNooDoc(&supplierNooDocReq))
	
	return 
}

func (s *SupplierController) UpdateDataSupplierNooDoc(c *gin.Context){
	supplierNooDocReq := dbmodels.SupplierNooDoc{}

	res := models.Response{}

	id, err := strconv.ParseInt(c.Params.ByName("supplier_id"), 10, 64)

	if err != nil {
		fmt.Println("Error")
	}
	
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierNooDocReq); err != nil {
		fmt.Println("Error, body Request")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	supplierNooDocReq.SupplierID = id

	fmt.Println("Edit data")
	c.JSON(http.StatusOK, SupplierNooDocService.UpdateDataSupplierNooDoc(&supplierNooDocReq))
}

// List Approve Supplier NOO Doc
func (s *SupplierController) FilterDataSupplierNooDocApprove(c *gin.Context) {
	req := dto.FilterSupplierNooDocDto{}
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

	supplier_id, errSupplierId := strconv.Atoi(c.Param("supplier_id"))
	if errSupplierId != nil {
		logs.Info("error", errPage)
		res.Error = errSupplierId.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	log.Println("page->", page, "count->", count, "supplier_id->", supplier_id)


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
	log.Println("searchCode-->", string(temp))
	res = SupplierNooDocService.GetDataSupplierNooDocApprovePaging(req, page, count, int64(supplier_id))

	c.JSON(http.StatusOK, res)

	return
}
/* --------------------------------------- Supplier Noo Doc ---------------------------------- */


/* --------------------------------------- SFA ----------------------------------------------- */

// Get List Data Supplier for SFA
func (s *SupplierController) ListDataSupplier(c *gin.Context){
	var suppliers []dbmodels.Supplier
	var res models.ResponseSFA

	data := SupplierService.GetDataListSupplier()
	temp, _ := json.Marshal(&data)
	if err := json.Unmarshal([]byte(temp), &suppliers); err != nil {
		fmt.Println("Error, body Request ")
		res.Meta.Status = false
		res.Meta.Code = 401
		res.Meta.Message = "Error, Body Request"
		res.Data = ""
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res.Data = suppliers
	res.Meta.Status = true
	res.Meta.Code = 200
	res.Meta.Message = "OK"

	c.JSON(http.StatusOK, res)
	return 
}

// Check Merchant in Supplier
func (s *SupplierController) CheckMerchantBySupplier(c *gin.Context) {
	var supplierMerchant []dbmodels.SupplierMerchant
	var checkSupplierReq dbmodels.SupplierMerchantReq
	var respSupplierNooChecklist []dbmodels.ResponseSupplierNooChecklist
	var respMerchantBySupplier models.ResponseMerchantBySupplier
	var techRequirement []string
	var docRequirement []string
	var res models.ResponseSFA

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &checkSupplierReq); err != nil {
		fmt.Println("Error, body Request ")
		res.Meta.Status = false
		res.Meta.Code = 401
		res.Meta.Message = "Error, Body Request"
		res.Data = ""
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	supplierMerchant = SupplierMerchantService.GetDataListMerchantBySupplier(checkSupplierReq.SupplierID, checkSupplierReq.MerchantID)
	respSupplierNooChecklist = SupplierNooChecklistService.GetDataNooCheckListBySupplier(checkSupplierReq.SupplierID)

	if len(supplierMerchant) > 0 {
		checkNooDoc := SupplierNooDocService.CheckDataNooMerchantBySupplier(checkSupplierReq.SupplierID, checkSupplierReq.MerchantID)
		if len(checkNooDoc) > 0 {
			for i:=0; i<len(checkNooDoc); i++ {
				if checkNooDoc[i].ApprovalStatus == "0" {
					respMerchantBySupplier.Message = "Data merchant ada"
					respMerchantBySupplier.Description = "Merchant sudah terdaftar, tetapi NOO belum diapprove"
				}else{
					respMerchantBySupplier.Message = "Data merchant ada"
					respMerchantBySupplier.Description = "Merchant sudah terdaftar dan NOO sudah diapprove"
				}
			}
		}else{
			if len(respSupplierNooChecklist) > 0 {
				for i:=0; i<len(respSupplierNooChecklist); i++ {
					if respSupplierNooChecklist[i].ChecklistType == "IMG" {
						docRequirement = append(docRequirement, respSupplierNooChecklist[i].Name)
					}else{
						techRequirement = append(techRequirement, respSupplierNooChecklist[i].Name)
					}
				}
			}else{
				docRequirement = []string{}
				techRequirement = []string{}
			}
			
			respMerchantBySupplier.Message = "Data merchant ada"
			respMerchantBySupplier.Description = "Merchant sudah terdaftar, tetapi NOO belum diupload"
			respMerchantBySupplier.SupplierID = checkSupplierReq.SupplierID
			respMerchantBySupplier.TechRequirement = techRequirement
			respMerchantBySupplier.DocRequirement = docRequirement
		}
	}else{
		if len(respSupplierNooChecklist) > 0 {
			for i:=0; i<len(respSupplierNooChecklist); i++ {
				if respSupplierNooChecklist[i].ChecklistType == "IMG" {
					docRequirement = append(docRequirement, respSupplierNooChecklist[i].Name)
				}else{
					techRequirement = append(techRequirement, respSupplierNooChecklist[i].Name)
				}
			}
		}else{
			docRequirement = []string{}
			techRequirement = []string{}
		}

		respMerchantBySupplier.Message = "Data merchant tidak ada"
		respMerchantBySupplier.Description = "Merchant belum terdaftar dan NOO belum diupload"
		respMerchantBySupplier.SupplierID = checkSupplierReq.SupplierID
		respMerchantBySupplier.TechRequirement = techRequirement
		respMerchantBySupplier.DocRequirement = docRequirement
	}

	res.Data = respMerchantBySupplier
	res.Meta.Status = true
	res.Meta.Code = 200
	res.Meta.Message = "OK"
	
	c.JSON(http.StatusOK, res)
	return 
}


// Register Merchant In Supplier
func (s *SupplierController) SubmitNOO(c *gin.Context){
	var supplierRegister dbmodels.RequestSupplierRegister
	var res models.ResponseSFA

	

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	
	if err := json.Unmarshal(dataBodyReq, &supplierRegister); err != nil {
		fmt.Println("Error, body Request")
		res.Data = ""
		res.Meta.Status = true
		res.Meta.Code = 401
		res.Meta.Message = "Error, body Request"
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} 

	merchantSupplier := SupplierMerchantService.GetDataListMerchantBySupplier(supplierRegister.SupplierID, supplierRegister.MerchantID)
	merchant := MerchantService.GetMerchantById(supplierRegister.MerchantID)

	if len(merchantSupplier) > 0 {
		respCheckNoo := CheckAndSubmitNoo(supplierRegister, merchant.Code)
		if respCheckNoo.ErrCode == constants.ERR_CODE_00 {
			res.Data = "Submit NOO berhasil disimpan"
			res.Meta.Code = 200
			res.Meta.Status = true
			res.Meta.Message = "OK"
		}else if respCheckNoo.ErrCode == "01" {
			res.Data = ""
			res.Meta.Code = 401
			res.Meta.Status = false
			res.Meta.Message = "NOO sudah diupload"
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}
	}else{
		// check merchant existing or no
		if merchant != (dbmodels.Merchant{}) {
			var supplierMerchant dbmodels.SupplierMerchant

			supplierMerchant.MerchantCode = merchant.Code
			supplierMerchant.SupplierId = int64(supplierRegister.SupplierID)
			//save supplier merchant
			respSupplierMerchant := SupplierMerchantService.SaveDataSupplierMerchant(&supplierMerchant)
			if respSupplierMerchant.ErrCode == constants.ERR_CODE_00 {
				respCheckNoo := CheckAndSubmitNoo(supplierRegister, merchant.Code)
				if respCheckNoo.ErrCode == constants.ERR_CODE_00 {
					res.Data = "Submit NOO berhasil disimpan"
					res.Meta.Code = 200
					res.Meta.Status = true
					res.Meta.Message = "OK"
				}else{
					res.Data = ""
					res.Meta.Code = 401
					res.Meta.Status = false
					res.Meta.Message = "NOO tidak berhasil disimpan"
					c.JSON(http.StatusBadRequest, res)
					c.Abort()
					return
				}
			}else{
				res.Data = ""
				res.Meta.Code = 401
				res.Meta.Status = false
				res.Meta.Message = "Merchant gagal didaftar"
				c.JSON(http.StatusBadRequest, res)
				c.Abort()
				return
			}
		}else{
			// save data merchant
			respMerchant := MerchantService.SaveDataMerchant(&dbmodels.Merchant{ID:int64(supplierRegister.MerchantID), Name:
			supplierRegister.MerchantName, Status:1, IssuerCode:1, LastUpdateBy:"SFA"})
			fmt.Println("generate merchant : ", respMerchant)
			if respMerchant.ErrCode == constants.ERR_CODE_00 {
				var supplierMerchant dbmodels.SupplierMerchant

				merchants := MerchantService.GetMerchantById(supplierRegister.MerchantID)

				supplierMerchant.MerchantCode = merchants.Code
				supplierMerchant.SupplierId = int64(supplierRegister.SupplierID)
				fmt.Println(merchants.Code)

				respSupplierMerchant := SupplierMerchantService.SaveDataSupplierMerchant(&dbmodels.SupplierMerchant{MerchantCode:merchants.Code,
				SupplierId:int64(supplierRegister.SupplierID)})
				if respSupplierMerchant.ErrCode == constants.ERR_CODE_00 {
					respCheckNoo := CheckAndSubmitNoo(supplierRegister, merchants.Code)
					if respCheckNoo.ErrCode == constants.ERR_CODE_00 {
						res.Data = "Submit NOO berhasil disimpan"
						res.Meta.Code = 200
						res.Meta.Status = true
						res.Meta.Message = "OK"
					}else{
						res.Data = ""
						res.Meta.Code = 401
						res.Meta.Status = false
						res.Meta.Message = "NOO tidak berhasil disimpan"
						c.JSON(http.StatusBadRequest, res)
						c.Abort()
						return
					}
				}else{
					res.Data = ""
					res.Meta.Code = 401
					res.Meta.Status = false
					res.Meta.Message = "Merchant gagal didaftar"
					c.JSON(http.StatusBadRequest, res)
					c.Abort()
					return
				}
			}
		}
	}

	c.JSON(http.StatusOK, res)
	return
}


func CheckAndSubmitNoo(supplierRegister dbmodels.RequestSupplierRegister, merchant_code string) models.Response {
	var res models.Response
	var resUpload models.ResponseUpload

	var pic []string
	var lookup []string
	var subbucket []string

	checkNooDoc := SupplierNooDocService.CheckDataNooMerchantBySupplier(supplierRegister.SupplierID, supplierRegister.MerchantID)
	lookupNooDoc := lookupService.GetLookupByGroup("NOO_DOCUMENT")

	// check noo doc empty or no
	if len(checkNooDoc) > 0 {
		res.ErrCode = "01"
		res.ErrDesc = "NOO Doc sudah ada"
	}else{
		// check doc label from sfa
		if len(supplierRegister.DocLabels) > 0 {
			for i := 0; i < len(supplierRegister.DocLabels); i++ {
				d := lookupNooDoc.Contents.([]dbmodels.Lookup)
				if len(d) > 0 {
					for j := range d {
						if d[j].Name == "img "+strings.ToLower(supplierRegister.DocLabels[i]){
							pic = append(pic, supplierRegister.DocImages[i])
							lookup = append(lookup, d[j].Code)
							subbucket = append(subbucket, strings.ToLower(supplierRegister.DocLabels[i]))
						}
					}
				}
			}
		}

		// make to struct NOO Doc
		nooDoc := dbmodels.SupplierNooDoc{SupplierID: int64(supplierRegister.SupplierID), MerchantCode:merchant_code, LastUpdateBy: dto.CurrUser}
		saveNooDoc := SupplierNooDocService.SaveDataSupplierNooDoc(&nooDoc)

		if saveNooDoc.ErrCode == constants.ERR_CODE_00 {
			if len(pic) > 0 {
				for j:=0; j< len(pic); j++ {
					merchantPict := SupplierService.GetMerchantPicturesByCodeAndLookup(merchant_code, lookup[j])
					if len(merchantPict) <= 0 {
						resUpload = UploadImage(merchant_code, pic[j], subbucket[j])
						fmt.Println(resUpload)
						if resUpload.ErrCode == constants.ERR_CODE_00 {
							respMerchantPict := SupplierService.SaveDataMerchantPicture(&dbmodels.MerchantPict{LookupCode: lookup[j], PictPath: resUpload.FileName, MerchantCode:merchant_code})
							if respMerchantPict.ErrCode == constants.ERR_CODE_00 {
								res.ErrCode = constants.ERR_CODE_00
								res.ErrDesc = constants.ERR_CODE_00_MSG
							}else{
								res.ErrCode = constants.ERR_CODE_03
								res.ErrDesc = constants.ERR_CODE_03_MSG
							}
						}
					}else{

					}
				}
			}
		}
	}

	return res
}


// upload image
func UploadImage(merchant_id string, images string, subbucket string) models.ResponseUpload {
	var res models.ResponseUpload
	idx := strings.Index(images, ";base64,")
	if idx < 0 {
		res.ErrCode = "05"
		res.ErrDesc = "Image tidak boleh null"
		return res
	}
	
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(images[idx+8:]))
	buff := bytes.Buffer{}
	
	_, err := buff.ReadFrom(reader)
	if err != nil {
		res.ErrCode = "05"
		res.ErrDesc = "Gagal membaca file"
		return res
	}
	
	_, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	fileName := merchant_id + "." + fm
	ioutil.WriteFile(fileName, buff.Bytes(), 0644)
	file, err := os.Open(fileName)

	//save image to minion
	SupplierService.UploadImageNooDoc(file, subbucket + "/" + fileName, "supplier")
	removeImage := os.Remove(fileName)
	fmt.Println(removeImage)

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG
	res.FileName = fileName

	return res
}


// Get List Product By Supplier And Approve
func (s *SupplierController) GetDataListProductBySupplierId(c *gin.Context) {
	var res models.ResponseSFA
	var requestSupplierProduct dbmodels.RequestSupplierProduct
	var listProducts []models.ResponseSupplierProduct

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &requestSupplierProduct); err != nil {
		res.Data = ""
		res.Meta.Status = false
		res.Meta.Code = 400
		res.Meta.Message = "Error, body Request"
		c.JSON(http.StatusBadRequest,res)
		c.Abort()
		return
	}

	nooDoc := SupplierNooDocService.CheckDataNooMerchantBySupplier(requestSupplierProduct.SupplierID, requestSupplierProduct.MerchantID)
	if len(nooDoc) <= 0 {
		res.Data = ""
		res.Meta.Code = 401
		res.Meta.Status = false
		res.Meta.Message = "NOO belum diupload. Silahkan upload NOO terlebih dahulu"
	}else{
		for i:=0;i<len(nooDoc);i++{
			if nooDoc[i].ApprovalStatus == "1" {
				listProduct := SupplierPriceService.GetDataListProductBySuppId(int64(requestSupplierProduct.SupplierID))
				if listProduct.ErrCode == constants.ERR_CODE_00 {
					d := listProduct.Contents.([]dbmodels.SupplierPrice)
					if len(d) > 0 {
						for j:=0; j<len(d); j++ {
							listProducts = append(listProducts, models.ResponseSupplierProduct{ID: d[j].Product.ID, 
								Title: d[j].Product.Name, Price: d[j].SellPrice, MainImageUrl: d[j].Product.IMG1})
						}
						res.Data = listProducts
						res.Meta.Status = true
						res.Meta.Code = 200
						res.Meta.Message = "OK"
					}else{
						res.Data = ""
						res.Meta.Status = false
						res.Meta.Code = 404
						res.Meta.Message = "Data tidak ada"
						c.JSON(http.StatusBadRequest,res)
						c.Abort()
						return
					}
				}else{
					res.Data = ""
					res.Meta.Status = false
					res.Meta.Code = 401
					res.Meta.Message = "Terjadi Kesalahan pada Server"
					c.JSON(http.StatusBadRequest,res)
					c.Abort()
					return
				}
				
			}else{
				res.Data = listProducts
				res.Meta.Status = false
				res.Meta.Code = 404
				res.Meta.Message = "NOO belum diapprove"
			}
		}		
	}

	c.JSON(http.StatusOK, res)
	return
}

// approve noo by tim ops
func (s *SupplierController) ApproveNOO(c *gin.Context){
	var res models.ResponseSFA
	var reqApproveNOO dbmodels.ApproveSupplierNooDoc
	var approveNOO dbmodels.SupplierNooDoc

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &reqApproveNOO); err != nil {
		res.Data = ""
		res.Meta.Status = false
		res.Meta.Code = 400
		res.Meta.Message = "Error, body Request"
		c.JSON(http.StatusBadRequest,res)
		c.Abort()
		return
	}

	merchant := MerchantService.GetMerchantById(int64(reqApproveNOO.MerchantID))
	noodoc := SupplierNooDocService.GetDataNOODocById(reqApproveNOO.ID)

	if merchant != (dbmodels.Merchant{}) && len(noodoc) > 0 {
		if reqApproveNOO.ApprovalStatus == "1" {
			approveNOO.ID = reqApproveNOO.ID
			approveNOO.SupplierID = int64(reqApproveNOO.SupplierID)
			approveNOO.MerchantCode = merchant.Code
			approveNOO.ApprovalStatus = "1"

			respApproveNOO := SupplierNooDocService.UpdateDataSupplierNooDoc(&approveNOO)
			if respApproveNOO.ErrCode == constants.ERR_CODE_00 {
				res.Data = "NOO diapprove"
			}else{
				res.Data = ""
				res.Meta.Code = 400
				res.Meta.Status = false
				res.Meta.Message = "Terjadi kesalahan"
				c.JSON(http.StatusBadRequest,res)
				c.Abort()
				return
			}
		}else if reqApproveNOO.ApprovalStatus == "2" {
			approveNOO.ID = reqApproveNOO.ID
			approveNOO.SupplierID = int64(reqApproveNOO.SupplierID)
			approveNOO.MerchantCode = merchant.Code
			approveNOO.ApprovalStatus = "2"

			respApproveNOO := SupplierNooDocService.UpdateDataSupplierNooDoc(&approveNOO)
			if respApproveNOO.ErrCode == constants.ERR_CODE_00 {
				res.Data = "NOO ditolak"
			}else{
				res.Data = ""
				res.Meta.Code = 400
				res.Meta.Status = false
				res.Meta.Message = "Terjadi kesalahan"
				c.JSON(http.StatusBadRequest,res)
				c.Abort()
				return
			}
		}
	}else{
		res.Data = "Data tidak ditemukan"
		res.Meta.Code = 401
		res.Meta.Status = false
		res.Meta.Message = "Not Found"
		c.JSON(http.StatusBadRequest,res)
		c.Abort()
		return
	}

	res.Meta.Code = 200
	res.Meta.Status = true
	res.Meta.Message = "OK"

	c.JSON(http.StatusOK, res)
	return
}
/* --------------------------------------- SFA ----------------------------------------------- */