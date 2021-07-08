package report

import (
	"distribution-system-be/database"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/signintech/gopdf"
)

var (
	purchaseOrderNumber string
)

func GeneratePurchaseOrderReport(purchaseOrderID int64) {

	title = "Purchase Order"

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
	fmt.Println("data  send to fillData Details : ", purchaseOrderID)
	dataDetails := fillDataDetailPurchaseOrder(purchaseOrderID)

	fmt.Println("hasil fill")
	for i, ordDetail := range dataDetails {
		fmt.Println(i, "====", ordDetail)
	}
	fmt.Println("=============")
	// setFont(&pdf, 12)
	setHeader(&pdf)
	pdf.Br(20)

	setDetail(&pdf, dataDetails, "")
	setSummary(&pdf)
	setSign(&pdf, "Admin", "Salesman", "Customer")

	pdf.WritePdf("purchase-order.pdf")

}

func fillDataDetailPurchaseOrder(purchaseOrderID int64) []DataDetail {

	purchaseOrder, err := database.GetPurchaseOrderByPurchaseOrderID(purchaseOrderID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(purchaseOrder)

	purchaseOrderNumber = purchaseOrder.PurchaserNo
	// purchaseOrderNo = purchaseOrder.PurchaserNo

	purchaseOrderDetails := database.GetAllDataDetailPurchaseOrder(purchaseOrderID)

	fmt.Println("Details : ", purchaseOrderDetails)

	go fillDataCustomer(
		purchaseOrder.Supplier.Code,
		purchaseOrder.Supplier.Name,
		purchaseOrder.PurchaserDate.Format("02-01-2006"),
		purchaseOrder.PurchaserNo,
	)
	// tdk blh kosong
	// per halaman max 25 item detail
	totalRec = len(purchaseOrderDetails)
	res := make([]DataDetail, totalRec+1)
	var data DataDetail

	subTotal = 0
	tax = 0
	grandTotal = 0
	for i, detail := range purchaseOrderDetails {
		data.Item = detail.Product.Name
		data.Quantity = int64(detail.Qty)
		data.Unit = detail.UOM.Name
		data.Price = int64(detail.Price)
		total := data.Price * data.Quantity
		data.Total = int64(detail.Price) * int64(detail.Qty)
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
