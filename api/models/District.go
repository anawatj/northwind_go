package models
import (
	"github.com/jinzhu/gorm"
)
type District struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	ProvinceID uint32 `gorm:"not null" json:"provinceId"`
}
func (u *District) FindAllDistricts(db *gorm.DB,rid uint32,pid uint32) (*[]District, error) {
	var err error
	districts := []District{}
	err = db.Debug().Model(&District{}).Where("province_id=?",pid).Find(&districts).Error
	if err != nil {
		return &[]District{}, err
	}
	return &districts, err
}