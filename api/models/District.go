package models

type District struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	ProvinceID uint32 `gorm:"not null" json:"provinceId"`
}