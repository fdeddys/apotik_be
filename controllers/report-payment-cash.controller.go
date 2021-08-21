package controllers

import (
	"distribution-system-be/models/dto"
	"distribution-system-be/services/reportservice"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// StockOpnameController ...
type ReportPaymentCashController struct {
	DB *gorm.DB
}

// StockOpnameService ...
var reportPaymentService = new(reportservice.ReportPaymentCashService)
var reportSalesService = new(reportservice.ReportSalesOrderService)

// DownloadTemplate ...
func (s *ReportPaymentCashController) DownloadReportPayment(c *gin.Context) {
	// fmt.Println("-------->", req)
	req := dto.FilterReportDate{}
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, unmarshal body Request to Receive stuct ", dataBodyReq)
		c.JSON(http.StatusBadRequest, "")
		c.Abort()
		return
	}

	filename := reportPaymentService.GenerateReportPaymentCash(req)

	header := c.Writer.Header()
	header["Content-type"] = []string{"text/csv"}
	header["Content-Disposition"] = []string{"attachment; filename=report.csv"}

	file, _ := os.Open(filename)

	io.Copy(c.Writer, file)

	os.Remove(filename)
	return
}

// DownloadTemplate ...
func (s *ReportPaymentCashController) DownloadReportSales(c *gin.Context) {
	// fmt.Println("-------->", req)
	req := dto.FilterReportDate{}
	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, unmarshal body Request to Receive stuct ", dataBodyReq)
		c.JSON(http.StatusBadRequest, "")
		c.Abort()
		return
	}

	filename := reportSalesService.GenerateReport(req)

	header := c.Writer.Header()
	header["Content-type"] = []string{"text/csv"}
	header["Content-Disposition"] = []string{"attachment; filename=report.csv"}

	file, _ := os.Open(filename)

	io.Copy(c.Writer, file)

	os.Remove(filename)
	return
}
