package models

type SubDistrict struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	DistrictID uint32 `gorm:"not null" json:"districtID"`
}