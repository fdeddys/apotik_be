package dbmodels

type TemplateProduct struct {
	ID        int64   `json:"id" gorm:"column:id"`
	Code      string  `json:"code" gorm:"column:code"`
	Nama      string  `json:"nama" gorm:"column:nama"`
	Plu       string  `json:"plu" gorm:"column:plu"`
	HargaJual float32 `json:"hargaJual" gorm:"column:harga_jual"`
	Status    int     `json:"status" gorm:"column:status"`
}

// TableName ...
func (t *TemplateProduct) TableName() string {
	return "public.template_product"
}
