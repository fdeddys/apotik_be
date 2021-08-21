package reportservice

import (
	"distribution-system-be/database"
	"distribution-system-be/models/dto"
)

// StockOpnameService ...
type ReportPaymentCashService struct {
}

// Approve ...
func (o ReportPaymentCashService) GenerateReportPaymentCash(filterData dto.FilterReportDate) (filename string) {

	dateStart := filterData.StartDate + " 00:00:00"
	dateEnd := filterData.EndDate + " 23:59:59"
	datas := generateDataReport(dateStart, dateEnd)
	filename = ExportToCSV(datas, filterData.StartDate, filterData.EndDate, "report-payment")
	return
}

func generateDataReport(dateStart, dateEnd string) []dto.ReportPaymentCash {

	datas := database.ReportPaymentCashByDate(dateStart, dateEnd)
	return datas
}
