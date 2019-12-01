package services

import (
	repository "distribution-system-be/database"
	"distribution-system-be/models"
	dbmodels "distribution-system-be/models/dbModels"
	"distribution-system-be/models/dto"
	"time"
)

// HistoryStockService ...
type HistoryStockService struct {
}

// SaveHistoryStock ...
func (h HistoryStockService) SaveHistoryStock(history *dbmodels.HistoryStock) models.NoContentResponse {
	history.LastUpdate = time.Now()
	history.LastUpdateBy = dto.CurrUser
	// var res models.ResponseSave
	res := repository.SaveHistory(*history)

	return res
}
