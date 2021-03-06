package services

import (
	"distribution-system-be/constants"
	"distribution-system-be/database"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"fmt"
	"time"
)

// AdjustmentService ...
type AdjustmentService struct {
}

// GetDataPage ...
func (a AdjustmentService) GetDataPage(param dto.FilterAdjustment, page, limit, status int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetAdjustmentPage(param, offset, limit, status)

	if err != nil {
		res.Error = err.Error()
		return res
	}

	res.Contents = data
	res.TotalRow = totalData
	res.Page = page
	res.Count = limit

	return res
}

// GetDataAdjustmentByID ...
func (a AdjustmentService) GetDataAdjustmentByID(reveiveID int64) dbmodels.Adjustment {

	var res dbmodels.Adjustment
	// var err error
	res, _ = database.GetAdjustmentByAdjustmentID(reveiveID)

	return res
}

// Save ...
func (a AdjustmentService) Save(adjustment *dbmodels.Adjustment) (errCode, errDesc, adjustmentNo string, adjustmentID int64, status int8) {

	if adjustment.ID == 0 {
		newNumber, errCode, errMsg := generateNewAdjustmentNo()
		if errCode != constants.ERR_CODE_00 {
			return errCode, errMsg, "", 0, 0
		}
		adjustment.AdjustmentNo = newNumber
	}
	adjustment.Status = 10
	adjustment.LastUpdateBy = dto.CurrUser
	adjustment.LastUpdate = time.Now()

	// fmt.Println("isi order ", order)
	err, errDesc, _, status := database.SaveAdjustment(adjustment)
	if err != constants.ERR_CODE_00 {
		return err, errDesc, "", 0, 0
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, adjustment.AdjustmentNo, adjustment.ID, adjustment.Status
}

// ApproveAdjustment ...
func (a AdjustmentService) ApproveAdjustment(order *dbmodels.Adjustment) (errCode, errDesc string) {

	// fmt.Println("isi order ", order)
	err, errDesc := database.SaveAdjustmentApprove(order)
	if err != constants.ERR_CODE_00 {
		return err, errDesc
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// RejectAdjustment ...
func (a AdjustmentService) RejectAdjustment(adjustment *dbmodels.Adjustment) (errCode, errDesc string) {

	// cek qty
	// validateQty()
	// fmt.Println("isi order ", order)
	err, errDesc := database.RejectAdjustment(adjustment)
	if err != constants.ERR_CODE_00 {
		return err, errDesc
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

func generateNewAdjustmentNo() (newNumber string, errCode string, errMsg string) {

	t := time.Now()
	bln := t.Format("01")
	thn := t.Format("06")
	header := "AJ"

	err, number, errdesc := database.AddSequence(bln, thn, header)
	if err != constants.ERR_CODE_00 {
		return "", err, errdesc
	}
	newNumb := fmt.Sprintf("00000%v", number)
	newNumb = newNumb[len(newNumb)-5 : len(newNumb)]
	newNumber = fmt.Sprintf("%v%v%v%v", header, thn, bln, newNumb)

	return newNumber, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG

}
