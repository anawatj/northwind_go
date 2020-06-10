package models
import (
	"errors"
	"github.com/jinzhu/gorm"
)
type Region struct {
	ID  uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}
func (region *Region) Validate() error {
	if region.Name == "" {
		return errors.New("Region name is required")
	}
	return nil
}
func (region *Region) SaveRegion(db *gorm.DB) (*Region, error){
	err := region.Validate()
	err = db.Debug().Create(&region).Error
	if err != nil {
		return &Region{}, err
	}
	return region, nil
}
func (u *Region) FindAllRegions(db *gorm.DB) (*[]Region, error) {
	var err error
	regions := []Region{}
	err = db.Debug().Model(&Region{}).Find(&regions).Error
	if err != nil {
		return &[]Region{}, err
	}
	return &regions, err
}
func (u *Region) FindRegionByID(db *gorm.DB, uid uint32) (*Region, error) {
	var err error
	err = db.Debug().Model(Region{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Region{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Region{}, errors.New("Region Not Found")
	}
	return u, err
}

func (region *Region) UpdateRegion(db *gorm.DB, uid uint32) (*Region, error) {

	err := region.Validate()
	
	db = db.Debug().Model(&Region{}).Where("id = ?", uid).Take(&Region{}).UpdateColumns(
		map[string]interface{}{
			"name":region.Name,
		},
	)
	if db.Error != nil {
		return &Region{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&Region{}).Where("id = ?", uid).Take(&region).Error
	if err != nil {
		return &Region{}, err
	}
	return region, nil
}

func (u *Region) DeleteRegion(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Region{}).Where("id = ?", uid).Take(&Region{}).Delete(&Region{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}