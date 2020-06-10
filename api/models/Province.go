package models
import (
	"github.com/jinzhu/gorm"
)

type Province struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	RegionID uint32 `gorm:"not null" json:"regionId"`
}

func (u *Province) FindAllProvince(db *gorm.DB,rid uint32) (*[]Province, error) {
	var err error
	provinces := []Province{}
	err = db.Debug().Model(&Province{}).Where("region_id = ?", rid).Find(&provinces).Error
	if err != nil {
		return &[]Province{}, err
	}
	return &provinces, err
}