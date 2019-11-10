package dbmodels

import (
	"time"
)

type SupplierNooChecklist struct {
	ID   				int64		`json:"id" gorm:"column:id"`
	SupplierID			int64 		`json:"supplierID" gor:"column:supplier_id"`
	LookupCode			string		`json:"lookupCode" gorm:"column:lookup_code"`
	Lookup 				Lookup	`json:"Lookup"`
	// Lookup 				[]Lookup 	`gorm:"many2many:supplier_noo_check_list;foreignkey:lookup_code;association_foreignkey:code;association_jointable_foreignkey:lookup_code;jointable_foreignkey:lookup_code;"`
	LastUpdateBy		string 		`json:"last_update_by"`
	LastUpdate			time.Time	`json:"last_update"`
	IsMandatory			int 		`json:"isMandatory" gorm:"column:is_mandatory"`
	ChecklistType 		string 		`json:"type" gorm:"column:checklist_type"`
}

type ResponseSupplierNooChecklist struct {
	Code			string		`json:"code" gorm:"column:code"`
	Name			string		`json:"name" gorm:"column:name"`
	SupplierID		int64		`json:"supplierID" gorm:"column:supplier_id"`
	LookupCode		string		`json:"lookupCode" gorm:"column:lookup_code"`
	IsMandatory		int			`json:"isMandatory" gor:"column:is_mandatory"`
	ChecklistType	string		`json:"checklistType" gorm:"column:checklist_type"`
	ID				int64 		`json:"id" gorm:"column:id"`
	Status			int64		`json:"status" gorm:"column:status"`
	LookupGroup		string		`json:"lookupGroup" gorm:"column:lookup_group"`
	LastUpdateBy	string 		`json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate		time.Time	`json:"last_update" gorm:"last_update"`
	LookupID		int64		`json:"lookupID" gorm:"column:lookup_id"`	
}


// TableName ...
func (s *SupplierNooChecklist) TableName() string {
	return "public.supplier_noo_check_list"
}