package dbmodels

// Lookup model ...
type Lookup struct {
	ID          int64  `json:"id" gorm:"column:id"`
	Status      int64  `json:"status" gorm:"column:status"`
	Code        string `json:"code" gorm:"column:code"`
	LookupGroup string `json:"lookup_group" gorm:"column:lookup_group"`
	Name        string `json:"name" gorm:"column:name"`
	IsViewable  string `json:"isViewable" gorm:"column:is_viewable"`
}

// TableName ...
func (t *Lookup) TableName() string {
	return "public.lookup"
}
