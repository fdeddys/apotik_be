package dbmodels

import "time"

// Pelanggan ...
type Pelanggan struct {
	ID       int64      `json:"id" gorm:"column:id"`
	TglMasuk *time.Time `json:"tglMasuk" gorm:"column:tgl_masuk"`
	Nama     string     `json:"nama" gorm:"column:nama"`
	Instansi string     `json:"instansi" gorm:"column:instansi"`
	NoStr    string     `json:"noStr" gorm:"column:no_str"`
	Profesi  string     `json:"profesi" gorm:"column:profesi"`
	NoHp     string     `json:"noHp" gorm:"column:no_hp"`
}

// TableName ...
func (Pelanggan) TableName() string {
	return "public.pelanggan"
}
