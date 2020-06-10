package models
import (
	"github.com/jinzhu/gorm"
)
type SubDistrict struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	DistrictID uint32 `gorm:"not null" json:"districtID"`
}
func (u *SubDistrict) FindAllSubDistrict(db *gorm.DB,rid uint32,pid uint32,did uint32) (*[]SubDistrict, error) {
	var err error
	subDistricts := []SubDistrict{}
	err = db.Debug().Model(&SubDistrict{}).Where("district_id=?",did).Find(&subDistricts).Error
	if err != nil {
		return &[]SubDistrict{}, err
	}
	return &subDistricts, err
}