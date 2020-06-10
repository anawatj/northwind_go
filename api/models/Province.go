package models


type Province struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	RegionID uint32 `gorm:"not null" json:"regionId"`
}

