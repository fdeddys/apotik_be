package dto

// RoleMenuDto ...
type RoleMenuDto struct {
	MenuID          int    `json:"menuId" gorm:"column:menuid"`
	MenuDescription string `json:"menuDescription" gorm:"column:menudescription"`
	Status          string `json:"status" gorm:"column:status"`
}
