package services

import (
	"encoding/json"
	"fmt"
	kons "oasis-be/constants"
	"oasis-be/models"
	dto "oasis-be/models/dto"

	db "oasis-be/database"

	model "oasis-be/models/dtoVendors"
	"oasis-be/utils/http"

	"github.com/astaxie/beego"
)

var (
	// Token ...
	Token                 string
	address               string
	password              string
	serverURL             string
	loginPath             string
	getSoByNumb           string
	sendOrderToKafkaTopik string
	kafkaRestServer       string
)

// VendorService ...
type VendorService struct {
}

func init() {
	address = beego.AppConfig.DefaultString("vendor.login.username", "")
	password = beego.AppConfig.DefaultString("vendor.login.password", "")
	serverURL = beego.AppConfig.DefaultString("vendor.url.server", "")
	loginPath = beego.AppConfig.DefaultString("vendor.url.login-path", "")
	getSoByNumb = beego.AppConfig.DefaultString("vendor.url.get-so-by-number", "")

	sendOrderToKafkaTopik = beego.AppConfig.DefaultString("kafka.topics", "")
	kafkaRestServer = beego.AppConfig.DefaultString("kafka.rest-server", "")

}

// Login ...
func (v VendorService) Login() model.LoginResponse {

	return loginUkirama()
}

// GetSalesOrder ...
func (v VendorService) GetSalesOrder() dto.FilterOrderResult {

	return GetBySalesOrder(1, 5)
}

// UpdateStatus ...
func (v VendorService) UpdateStatus(req models.RequestUpdateStatus) models.ContentResponse {

	res := models.ContentResponse{}

	order, err := db.GetOrderByOrderNo(req.SoNumber)

	if err != nil {
		res.ErrCode = kons.ERR_CODE_81
		res.ErrDesc = err.Error()
		return res
	}

	if order.InternalStatus != 1 {
		curStatus := ""
		switch order.InternalStatus {
		case 0:
			curStatus = "NEW ORDER"
			break
		case 2:
			curStatus = "READY TO PAYMENT"
			break
		case 3:
			curStatus = "ORDER COMPLETE"
			break
		}
		res.ErrCode = kons.ERR_CODE_81
		res.ErrDesc = "status not IN PROGRESS [" + curStatus + "]"
		return res
	}

	res.ErrCode = kons.ERR_CODE_00
	res.ErrDesc = kons.ERR_CODE_00_MSG
	res.Contents = order
	return res
}

// ProduceToKafkaServer ...
func ProduceToKafkaServer(salesOrder dto.SalesOrder) {

	var kafkaBody dto.KafkaBody
	kafkaBody.Topic = sendOrderToKafkaTopik
	kafkaBody.Data = salesOrder
	sendHTTP(kafkaRestServer, kafkaBody)
}

// GetBySalesOrder ...
func GetBySalesOrder(page, count int) dto.FilterOrderResult {

	result := dto.FilterOrderResult{}
	getToken := loginUkirama()

	if getToken.ErrCode != "00" {
		result.ErrDesc = getToken.ErrDesc
		return result
	}

	param := model.ParamGetSoDto{}
	param.AccessToken = getToken.Token
	param.Page = string(page)
	param.PerPage = string(count)
	// param.SalesOrderCodes = "SO19030017"

	url := serverURL + getSoByNumb
	// dataURL := fmt.Sprintf("access_token=%s&page=%s&per_page=%s&sales_order_codes=%s", param.AccessToken, param.Page, param.PerPage, param.SalesOrderCodes)
	dataURL := fmt.Sprintf("access_token=%s&page=%s&per_page=%s", param.AccessToken, param.Page, param.PerPage)

	newURL := fmt.Sprintf("%s?%s", url, dataURL)
	result.ErrDesc = newURL

	res, err := sendGetHTTP(newURL)
	if err == nil {
		respVendor := model.SalesOrderResultDto{}
		err2 := json.Unmarshal(res, &respVendor)
		if err2 == nil {
			fmt.Println("Hasil ==> ", respVendor.Success)
			result.Data = respVendor.Success.Data
			result.ErrCode = kons.ERR_CODE_00
			result.ErrDesc = kons.ERR_CODE_00_MSG
			result.Page = respVendor.Success.Page
			result.PerPage = respVendor.Success.PerPage
			result.TotalPages = respVendor.Success.TotalPages
			result.TotalRecords = respVendor.Success.TotalRecords
		} else {
			result.ErrCode = kons.ERR_CODE_03
			result.ErrDesc = kons.ERR_CODE_03_MSG
		}
	} else {
		result.ErrCode = kons.ERR_CODE_70
		result.ErrDesc = kons.ERR_CODE_70_MSG + err.Error()
	}
	return result
}

func loginUkirama() model.LoginResponse {
	var req model.LoginRequest
	result := model.LoginResponse{}
	respVendor := model.LoginResponseVendor{}

	fmt.Println("===== LOGIN UKIRAMA =========================================================")

	url := serverURL + loginPath
	req.Username = address
	req.Password = password

	res, err := sendHTTP(url, req)

	if err == nil {
		fmt.Println("Login vendor success ....")
		err2 := json.Unmarshal(res, &respVendor)
		if err2 == nil {
			// result.AccessToken = res.AccessToken
			fmt.Println("token ==> ", respVendor.AccessToken)

			if len(respVendor.AccessToken) < 1 {
				result.ErrCode = kons.ERR_CODE_53
				result.ErrDesc = kons.ERR_CODE_53_MSG + " [Vendor] " + respVendor.Error
				fmt.Println("failed from vendor ", respVendor)
			} else {
				Token = respVendor.AccessToken
				result.ErrCode = kons.ERR_CODE_00
				result.ErrDesc = kons.ERR_CODE_00_MSG + " [Vendor] "
				result.Token = respVendor.AccessToken
			}
			fmt.Println("===== END-LOGIN UKIRAMA =========================================================")
			return result
		}
		fmt.Println("failed unmarshal ", err2)
		result.ErrCode = kons.ERR_CODE_03
		result.ErrDesc = kons.ERR_CODE_03_MSG + " [Vendor]" + err2.Error()
		fmt.Println("===== END-LOGIN UKIRAMA =========================================================")
		return result
	}
	fmt.Println("Login failed.. ")
	result.ErrCode = kons.ERR_CODE_50
	result.ErrDesc = kons.ERR_CODE_50_MSG + " [Vendor]" + err.Error()
	fmt.Println("===== END-LOGIN UKIRAMA =========================================================")
	// result.Error = string(res)
	return result
}

// sendHTTP ...
func sendHTTP(url string, req interface{}) ([]byte, error) {

	return http.HttpPost(url, req, "60s", 1)
}

// sendGetHTTP ...
func sendGetHTTP(url string) ([]byte, error) {
	return http.HttpGetByParam(url, "60s")
}
