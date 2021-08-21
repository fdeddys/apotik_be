package reportservice

import (
	"distribution-system-be/database"
	"distribution-system-be/models/dto"
)

// StockOpnameService ...
type ReportSalesOrderService struct {
}

// Approve ...
func (o ReportSalesOrderService) GenerateReport(filterData dto.FilterReportDate) (filename string) {

	dateStart := filterData.StartDate + " 00:00:00"
	dateEnd := filterData.EndDate + " 23:59:59"
	datas := generateDataReportSales(dateStart, dateEnd)
	filename = ExportToCSV(datas, filterData.StartDate, filterData.EndDate, "report-sales-order")
	return
}

func generateDataReportSales(dateStart, dateEnd string) []dto.ReportSales {

	datas := database.ReportSalesByDate(dateStart, dateEnd)
	return datas
}
