package models
import (
	"errors"
	"github.com/jinzhu/gorm"

)
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

func (c *Customer) Validate(action string) error {
	if c.CompanyName==""{
		return errors.New("Please input company name")
	}
	if c.ContactName==""{
		return errors.New("Please input contact name")
	}
	if c.ContactTtle==""{
		return errors.New("Please input contact title")
	}
	if c.Address==""{
		return errors.New("Please input address")
	}
	if c.RegionID==0{
		return errors.New("Please input region")
	}
	if c.ProvinceID == 0{
		return errors.New("Please input province")
	}
	if c.DistrictID==0{
		return errors.New("Please input district")
	}
	if c.SubDistrictID==0{
		return errors.New("please input sub district")
	}
	if c.PostCode==""{
		return errors.New("Please input postal code")
	}
	if c.Phone==""{
		return errors.New("Please input phone")
	}
	if c.Mobile==""{
		return errors.New("Please input mobile")
	}
	if c.Fax==""{
		return errors.New("Please input fax")
	}

	return nil
}
func (c *Customer) SaveCustomer(db *gorm.DB) (*Customer, error) {

	var err error
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}
