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

// ReceiveService ...
type ReceiveService struct {
}

// GetDataPage ...
func (r ReceiveService) GetDataPage(param dto.FilterReceive, page, limit, status int) models.ResponsePagination {
	var res models.ResponsePagination

	offset := (page - 1) * limit
	data, totalData, err := database.GetReceivePage(param, offset, limit, status)

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

// GetDataReceiveByID ...
func (r ReceiveService) GetDataReceiveByID(reveiveID int64) dbmodels.Receive {

	var res dbmodels.Receive
	// var err error
	res, _ = database.GetReceiveByReceiveID(reveiveID)

	return res
}

// Save ...
func (r ReceiveService) Save(receive *dbmodels.Receive) (errCode, errDesc, receiveNo string, receiveID int64, status int8) {

	if receive.ID == 0 {
		newNumber, errCode, errMsg := generateNewReceiveNo()
		if errCode != constants.ERR_CODE_00 {
			return errCode, errMsg, "", 0, 0
		}
		receive.ReceiveNo = newNumber
		receive.Status = 10
	}
	receive.LastUpdateBy = dto.CurrUser
	receive.LastUpdate = time.Now()

	// fmt.Println("isi order ", order)
	err, errDesc, _, status := database.SaveReceive(receive)
	if err != constants.ERR_CODE_00 {
		return err, errDesc, "", 0, 0
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG, receive.ReceiveNo, receive.ID, status
}

// ApproveReceive ...
func (r ReceiveService) ApproveReceive(order *dbmodels.Receive) (errCode, errDesc string) {

	// fmt.Println("isi order ", order)
	err, errDesc := database.SaveReceiveApprove(order)
	if err != constants.ERR_CODE_00 {
		return err, errDesc
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

// RejectReceive ...
func (o OrderService) RejectReceive(receive *dbmodels.Receive) (errCode, errDesc string) {

	// cek qty
	// validateQty()
	// fmt.Println("isi order ", order)
	err, errDesc := database.RejectReceive(receive)
	if err != constants.ERR_CODE_00 {
		return err, errDesc
	}
	return constants.ERR_CODE_00, constants.ERR_CODE_00_MSG
}

func generateNewReceiveNo() (newNumber string, errCode string, errMsg string) {

	t := time.Now()
	bln := t.Format("01")
	thn := t.Format("06")
	header := "RV"

	err, number, errdesc := database.AddSequence(bln, thn, header)
	if err != constants.ERR_CODE_00 {
		return "", err, errdesc
	}
	newNumb := fmt.Sprintf("00000%v", number)
	newNumb = newNumb[len(newNumb)-5 : len(newNumb)]
	newNumber = fmt.Sprintf("%v%v%v%v", header, thn, bln, newNumb)

	return newNumber, constants.ERR_CODE_00, constants.ERR_CODE_00_MSG

}
