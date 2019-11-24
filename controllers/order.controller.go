package controllers

import (
	"distribution-system-be/database"
	"distribution-system-be/models"
	"distribution-system-be/models/dbModels"
	"distribution-system-be/models/dto"
	"distribution-system-be/services"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"distribution-system-be/constants"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/signintech/gopdf"
)

// OrderController ...
type OrderController struct {
	DB *gorm.DB
}

type DataDetail struct {
	Item     string
	Quantity int64
	Unit     string
	Price    int64
	Total    int64
}

type InvHdrInfo struct {
	CustCode  string
	CustName  string
	TransAt   string
	SourceDoc string
}

// OrderService ...
var OrderService = new(services.OrderService)

var (
	// length New Line
	spaceLen float64

	// page margin
	pageMargin float64

	// customer region
	spaceCustomerInfo float64
	spaceTitik        float64
	spaceValue        float64

	spaceSummaryInfo  float64
	spaceTitikSumamry float64
	spaceValueSummary float64

	// table
	tblCol1 float64
	tblCol2 float64
	tblCol3 float64
	tblCol4 float64
	tblCol5 float64
	tblCol6 float64

	curPage     int
	number      int
	dataDetails []DataDetail
	totalRec    int
	invoiceNumb string

	// count by system
	subTotal   int64
	tax        int64
	grandTotal int64

	invInfo InvHdrInfo
)

// GetByOrderId ...
func (s *OrderController) GetByOrderId(c *gin.Context) {

	res := dbmodels.SalesOrder{}

	orderID, errPage := strconv.Atoi(c.Param("id"))
	if errPage != nil {
		logs.Info("error", errPage)
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = OrderService.GetDataOrderById(orderID)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return

}

// FilterData ...
func (s *OrderController) FilterData(c *gin.Context) {
	req := dto.FilterOrder{}
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
	log.Println("is release ", req.InternalStatus)

	intStatus := -1
	if intVal, errconv := strconv.Atoi(req.InternalStatus); errconv == nil {
		intStatus = intVal
	}
	res = OrderService.GetDataPage(req, page, count, intStatus)

	c.JSON(http.StatusOK, res)

	return
}

// SaveSO ...
func (s *OrderController) SaveSO(c *gin.Context) {

	req := dbmodels.SalesOrder{}
	body := c.Request.Body
	res := dto.OrderSaveResult{}
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, unmarshal body Request to Sales Order stuct ", dataBodyReq)
		res.ErrDesc = constants.ERR_CODE_03_MSG
		res.ErrCode = constants.ERR_CODE_03
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	errCode, errMsg := OrderService.Save(&req)
	res.ErrDesc = errMsg
	res.ErrCode = errCode
	// res.OrderNo = newNumb
	c.JSON(http.StatusOK, res)

	return
}

// PrintInvoice ...
func (s *OrderController) PrintInvoice(c *gin.Context) {

	req := dbmodels.SalesOrder{}
	body := c.Request.Body
	res := dto.OrderSaveResult{}
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, unmarshal body Request to Sales Order stuct ", dataBodyReq)
		res.ErrDesc = constants.ERR_CODE_03_MSG
		res.ErrCode = constants.ERR_CODE_03
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	// fmt.Println("-------->", req)

	generateRep(req)

	header := c.Writer.Header()
	// header["Content-type"] = []string{"application/octet-stream"}
	header["Content-type"] = []string{"application/x-pdf"}
	header["Content-Disposition"] = []string{"attachment; filename= tes.pdf"}

	// file, _ := os.Open("/Users/deddysyuhendra/go/src/tes-print/invoice.pdf")
	file, _ := os.Open("invoice.pdf")

	io.Copy(c.Writer, file)
	return
}

func generateRep(order dbmodels.SalesOrder) {

	spaceLen = beego.AppConfig.DefaultFloat("report.space-len", 15)
	pageMargin = beego.AppConfig.DefaultFloat("report.page-margin", 12)

	curPage = 1

	spaceCustomerInfo = 300
	spaceTitik = spaceCustomerInfo + 150
	spaceValue = spaceCustomerInfo + 160

	spaceSummaryInfo = spaceCustomerInfo
	spaceTitikSumamry = spaceTitik
	spaceValueSummary = spaceValue

	tblCol1 = 25
	tblCol2 = 80
	tblCol3 = 300
	tblCol4 = 370
	tblCol5 = 430
	tblCol6 = 500

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.SetMargins(pageMargin, pageMargin, pageMargin, pageMargin)
	pdf.AddPage()

	if err := pdf.AddTTFFont("open-sans", "font/OpenSans-Regular.ttf"); err != nil {
		log.Print(err.Error())
		return
	}

	if err := pdf.AddTTFFont("open-sans-bold", "font/OpenSans-Bold.ttf"); err != nil {
		log.Print(err.Error())
		return
	}

	// untuk nomor urut barang
	number = 1

	// get Data mockup utk display ke grid
	fmt.Println("data order send to fillData Details : ", order)
	dataDetails := fillDataDetail(order.SalesOrderNo)

	fmt.Println("hasil fill")
	for i, ordDetail := range dataDetails {
		fmt.Println(i, "====", ordDetail)
	}
	fmt.Println("=============")
	// setFont(&pdf, 12)
	setHeader(&pdf)
	pdf.Br(20)

	setDetail(&pdf, dataDetails)
	setSummary(&pdf)
	setSign(&pdf, "Admin", "Salesman", "Customer")
	// 595, H: 842
	// pdf.SetFont("open-sans", "", 14)

	// pdf.SetFont("open-sans", "", 10)
	// for i := 2; i <= 83; i++ {
	// 	pdf.SetX(1)
	// 	pdf.SetY(10 * float64(i))
	// 	pdf.Text(fmt.Sprintf("%v", i))
	// }
	pdf.WritePdf("invoice.pdf")

}

func fillDataDetail(orderNo string) []DataDetail {

	order, err := database.GetOrderByOrderNo(orderNo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(order)

	// invoiceNumb = "IVyymm999999"
	invoiceNumb = order.SalesOrderNo

	orderDetails := database.GetAllDataDetail(order.ID)

	fmt.Println("orderDetails : ", orderDetails)

	go fillDataCustomer(order)
	// tdk blh kosong
	// per halaman max 25 item detail
	totalRec = len(orderDetails)
	res := make([]DataDetail, totalRec+1)
	var data DataDetail

	for i, ordDetail := range orderDetails {
		data.Item = ordDetail.Product.Name
		data.Quantity = int64(ordDetail.Qty)
		data.Unit = ordDetail.UOM.Name
		data.Price = int64(ordDetail.Price)
		total := data.Price * data.Quantity
		data.Total = int64(ordDetail.Price)
		subTotal += total
		res[i+1] = data
		fmt.Println("total sub total", subTotal)
	}
	totalRec = len(res)
	fmt.Println("Jumlah record [fill] =>", totalRec)

	tax = subTotal / 10
	grandTotal = subTotal + tax

	return res
}

func fillDataCustomer(order dbmodels.SalesOrder) {
	invInfo.CustCode = order.Customer.Code
	invInfo.CustName = order.Customer.Name
	invInfo.TransAt = order.OrderDate.Format("02-01-2006")
	invInfo.SourceDoc = order.SalesOrderNo
}

func setHeader(pdf *gopdf.GoPdf) {

	showLogo(pdf)
	showCompany(pdf)
	space(pdf)
	showLine(pdf)
	showInvNo(pdf)

}

func showInvNo(pdf *gopdf.GoPdf) {

	pdf.SetY(30)
	pdf.SetX(450)
	setFontBold(pdf, 10)
	pdf.Text("SALES INVOICE")

	space(pdf)
	setFont(pdf, 12)
	pdf.SetX(450)
	pdf.Text(invoiceNumb)
}

func showCompany(pdf *gopdf.GoPdf) {

	line1 := beego.AppConfig.DefaultString("report.line1", "PT. Reksa Transaksi Sukses Makmur")
	line2 := beego.AppConfig.DefaultString("report.line2", "Plaza Mutiara Lt 21 Suite 2105")
	line3 := beego.AppConfig.DefaultString("report.line3", "Jl. DR. Ide Anak Agung Gde Agung")
	line4 := beego.AppConfig.DefaultString("report.line4", "Kav")
	line5 := beego.AppConfig.DefaultString("report.line5", "Setiabudi")
	line6 := beego.AppConfig.DefaultString("report.line6", "Postal code")

	pdf.Br(15)

	setFontBold(pdf, 10)
	pdf.SetX(200)
	pdf.Text(line1)

	space(pdf)
	setFont(pdf, 10)
	pdf.SetX(200)
	pdf.Text(line2)

	space(pdf)
	pdf.SetX(200)
	pdf.Text(line3)

	space(pdf)
	pdf.SetX(200)
	pdf.Text(line4)

	space(pdf)
	pdf.SetX(200)
	pdf.Text(line5)

	space(pdf)
	pdf.SetX(200)
	pdf.Text(line6)
}

func showLogo(pdf *gopdf.GoPdf) {

	imgSize := spaceLen * 5
	posX := 20.0
	posY := spaceLen

	pdf.Image("imgs/logo.png", posX, posY, &gopdf.Rect{W: imgSize + 68, H: imgSize})
}

func setDetail(pdf *gopdf.GoPdf, data []DataDetail) {

	setPageNumb(pdf)
	pdf.SetX(20)
	pdf.SetY(spaceLen * 8)

	showCustomer(pdf)

	space(pdf)
	showHeaderTable(pdf)

	fmt.Println("Panjang array ", len(data), "] ")
	fmt.Println("Total rec => set detail => ", totalRec, "] ")
	fmt.Println("start iterate")
	// var dataDetail DataDetail
	if totalRec > 1 {
		for i := 1; i <= 25; i++ {
			fmt.Println("idx ke [", i, "]", data[number])
			space(pdf)
			showData(pdf, fmt.Sprintf("%v", number), data[number].Item, data[number].Unit, data[number].Quantity, data[number].Price, data[number].Total)
			number++
			if number >= totalRec {
				break
			}
		}
	}
	// }

	space(pdf)
	showLine(pdf)

	// jika data masih ada utk next page
	// 1. add page
	// 2. set header
	// 3. rekursif
	if totalRec > number {
		fmt.Println("NEW page")
		curPage++
		pdf.AddPage()
		setHeader(pdf)
		setDetail(pdf, data)
	}
}

func setSummary(pdf *gopdf.GoPdf) {

	setFont(pdf, 10)

	space(pdf)
	// pdf.SetY(spaceLen * 42)
	pdf.SetX(spaceSummaryInfo)
	pdf.Text("Subtotal")
	pdf.SetX(spaceTitikSumamry)
	pdf.Text(":")
	pdf.SetX(spaceValueSummary)
	// fmt.Println("isi subtotal utk summart ", subTotal)
	pdf.Text(fmt.Sprintf("%v", subTotal))

	space(pdf)
	pdf.SetX(spaceSummaryInfo)
	pdf.Text("Tax ")
	pdf.SetX(spaceTitikSumamry)
	pdf.Text(":")
	pdf.SetX(spaceValueSummary)
	pdf.Text(fmt.Sprintf("%v", tax))

	space(pdf)
	pdf.SetX(spaceSummaryInfo)
	pdf.Text("GrandTotal ")
	pdf.SetX(spaceTitikSumamry)
	pdf.Text(":")
	pdf.SetX(spaceValueSummary)
	pdf.Text(fmt.Sprintf("%v", grandTotal))

}

func showHeaderTable(pdf *gopdf.GoPdf) {

	showLine(pdf)
	space(pdf)
	setFontBold(pdf, 10)
	pdf.SetX(tblCol1)
	pdf.Text("#")

	pdf.SetX(tblCol2)
	pdf.Text("Item")

	pdf.SetX(tblCol3)
	pdf.Text("Quantity")

	pdf.SetX(tblCol4)
	pdf.Text("Unit")

	pdf.SetX(tblCol5)
	pdf.Text("Price")

	pdf.SetX(tblCol6)
	pdf.Text("Total")

	space(pdf)
	showLine(pdf)
}

func showData(pdf *gopdf.GoPdf, no, item, unit string, qty, price, total int64) {

	setFont(pdf, 10)
	pdf.SetX(tblCol1)
	pdf.Text(no)

	pdf.SetX(tblCol2)
	pdf.Text(item)

	pdf.SetX(tblCol3)
	pdf.Text(fmt.Sprintf("%v", qty))

	pdf.SetX(tblCol4)
	pdf.Text(unit)

	pdf.SetX(tblCol5)
	pdf.Text(fmt.Sprintf("%v", price))

	pdf.SetX(tblCol6)
	pdf.Text(fmt.Sprintf("%v", total))
}

func showCustomer(pdf *gopdf.GoPdf) {
	// , code, name, transDate, ssNo string
	space(pdf)
	setFont(pdf, 10)

	pdf.SetX(spaceCustomerInfo)
	pdf.Text("Customer Code")
	pdf.SetX(spaceTitik)
	pdf.Text(":")
	pdf.SetX(spaceValue)
	pdf.Text(invInfo.CustCode)

	space(pdf)
	pdf.SetX(spaceCustomerInfo)
	pdf.Text("Customer ")
	pdf.SetX(spaceTitik)
	pdf.Text(":")
	pdf.SetX(spaceValue)
	pdf.Text(invInfo.CustName)

	space(pdf)
	pdf.SetX(spaceCustomerInfo)
	pdf.Text("Transaction at ")
	pdf.SetX(spaceTitik)
	pdf.Text(":")
	pdf.SetX(spaceValue)
	pdf.Text(invInfo.TransAt)

	space(pdf)
	pdf.SetX(spaceCustomerInfo)
	pdf.Text("Source Document ")
	pdf.SetX(spaceTitik)
	pdf.Text(":")
	pdf.SetX(spaceValue)
	pdf.Text(invInfo.SourceDoc)

}

func setSign(pdf *gopdf.GoPdf, sign1, sign2, sign3 string) {

	// pdf.SetY(spaceLen * 48)

	xSign1 := tblCol1
	xSign2 := tblCol1 + 200
	xSign3 := tblCol1 + 400
	maxLengLine := 100

	xLengSign1 := xSign1 + float64(maxLengLine)
	xLengSign2 := xSign2 + float64(maxLengLine)
	xLengSign3 := xSign3 + float64(maxLengLine)

	space(pdf)
	space(pdf)
	space(pdf)
	space(pdf)

	if sign1 != "" {
		pdf.SetX(xSign1)
		pdf.Text(sign1)
	}

	if sign2 != "" {
		pdf.SetX(xSign2)
		pdf.Text(sign2)
	}

	if sign3 != "" {
		pdf.SetX(xSign3)
		pdf.Text(sign3)
	}

	space(pdf)
	space(pdf)
	space(pdf)
	space(pdf)

	if sign1 != "" {
		pdf.SetX(xSign1)
		pdf.Line(xSign1, pdf.GetY(), xLengSign1, pdf.GetY())
	}

	if sign2 != "" {
		pdf.SetX(xSign2)
		pdf.Line(xSign2, pdf.GetY(), xLengSign2, pdf.GetY())
	}

	if sign3 != "" {
		pdf.SetX(xSign3)
		pdf.Line(xSign3, pdf.GetY(), xLengSign3, pdf.GetY())
	}

}

func setPageNumb(pdf *gopdf.GoPdf) {

	setFont(pdf, 10)
	pdf.SetX(595 - pageMargin - 40)
	pdf.SetY(842 - (pageMargin * 2))
	pdf.Text(fmt.Sprintf("Page %v", curPage))

}

func showLine(pdf *gopdf.GoPdf) {
	pdf.SetX(200)
	pdf.Line(pdf.MarginLeft(), pdf.GetY(), 575.0, pdf.GetY())
}

func setFont(pdf *gopdf.GoPdf, size int) {
	if err := pdf.SetFont("open-sans", "", size); err != nil {
		log.Print(err.Error())
		return
	}
	// pdf.SetFont("open-sans", "", size)
}

func setFontBold(pdf *gopdf.GoPdf, size int) {
	if err := pdf.SetFont("open-sans-bold", "", size); err != nil {
		log.Print(err.Error())
		return
	}
	// pdf.SetFont("open-sans", "", size)
}

func space(pdf *gopdf.GoPdf) {
	pdf.Br(spaceLen)
}
