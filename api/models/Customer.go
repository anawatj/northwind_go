package models

type Customer struct{
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CompanyName  string `gorm:"size:255;null;unique" json:"companyName"`
	ContactName  string `gorm:"size:255;null" json:"contactName"`
	ContactTtle  string `gorm:"size:255;null" json:"contactTitle"`
	Address      string `gorm:"size:1000;null" json:"address"`
	RegionID     uint32 `gorm:"null" json:"regionId"`
	ProvinceID   uint32 `gorm:"null" json:"provinceId"`
	DistrictID   uint32 `gorm:"null" json:"districtId"`
	SubDistrictID uint32 `gorm:"null" json:"subDistrictId"`
	PostCode  string `gorm:"size:100;null" json:"postCode"`
	Phone string `gorm:"size:100;null" json:"phone"`
	Mobile string `gorm:"size:100;null" json:"mobile"`
	Fax string `gorm:"size:100;null" json:"fax"`
	
}