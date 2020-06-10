package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/northwind_go/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname:"Tao",
		Email:"ajarusiripot@gmail.com",
		Password:"P@ssw0rd",
	},
}
var regions =[]models.Region{
	models.Region{
		Name:"North",
	},
	models.Region{
		Name:"Middle",
	},
	models.Region{
		Name:"NorthEast",
	},
	models.Region{
		Name:"South",
	},

}



func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists( &models.User{},&models.Region{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{},&models.Region{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}



	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	

	
	}
	for i, _ := range regions{
		err = db.Debug().Model(&models.Region{}).Create(&regions[i]).Error
		if err != nil {
			log.Fatalf("cannot seed regions table: %v",err)
		}
	}
}