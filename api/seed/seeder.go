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
var provinces = []models.Province{
	models.Province{
		Name:"Bangkok",
		RegionID:2,
	},
	models.Province{
		Name:"Nontaburi",
		RegionID:2,
	},
	models.Province{
		Name:"Patumtani",
		RegionID:2,
	},
	models.Province{
		Name:"Samutprakarn",
		RegionID:2,
	},
	models.Province{
		Name:"Chachengsao",
		RegionID:2,
	},
	models.Province{
		Name:"Nakornpathom",
		RegionID:2,
	},
	models.Province{
		Name:"Samutsakorn",
		RegionID:2,
	},
	models.Province{
		Name:"Samutsongkarm",
		RegionID:2,
	},
	models.Province{
		Name:"Singburi",
		RegionID:2,
	},
	models.Province{
		Name:"Saraburi",
		RegionID:2,
	},
	models.Province{
		Name:"Ayuttaya",
		RegionID:2,
	},
	models.Province{
		Name:"Lopburi",
		RegionID:2,
	},
	models.Province{
		Name:"Angthong",
		RegionID:2,
	},
	models.Province{
		Name:"Supanburi",
		RegionID:2,
	},
	models.Province{
		Name:"Chainat",
		RegionID:2,
	},
	models.Province{
		Name:"Ratchaburi",
		RegionID:2,
	},
	models.Province{
		Name:"Petchaburi",
		RegionID:2,
	},
	models.Province{
		Name:"Prajubkirikhan",
		RegionID:2,
	},
	models.Province{
		Name:"Kanjanaburi",
		RegionID:2,
	},
	models.Province{
		Name:"Chonburi",
		RegionID:2,

	},
	models.Province{
		Name:"Rayong",
		RegionID:2,
	},
	models.Province{
		Name:"Chantaburi",
		RegionID:2,
	},
	models.Province{
		Name:"Trad",
	},
	models.Province{
		Name:"Nakornnayok",
		RegionID:2,
	},
	models.Province{
		Name:"Prajinburi",
		RegionID:2,
	},
	models.Province{
		Name:"Srakaew",
		RegionID:2,
	},
	models.Province{
		Name:"Chiangmai",
		RegionID:1,
	},
	models.Province{
		Name:"Chiangrai",
		RegionID:1,
	},
	models.Province{
		Name:"Payao",
		RegionID:1,
	},
	models.Province{
		Name:"Lampoon",
		RegionID:1,
	},
	models.Province{
		Name:"Maehongson",
		RegionID:1,
	},
	models.Province{
		Name:"Lampang",
		RegionID:1,
	},
	models.Province{
		Name:"Phare",
		RegionID:1,
	},
	models.Province{
		Name:"Nan",
		RegionID:1,
	},
	models.Province{
		Name:"Uttradit",
		RegionID:1,
	},
	models.Province{
		Name:"Sukhothai",
		RegionID:1,
	},
	models.Province{
		Name:"Tak",
		RegionID:1,

	},
	models.Province{
		Name:"Kamphangpetch",
		RegionID:1,
	},
	models.Province{
		Name:"Petchaboon",
		RegionID:1,
	},
	models.Province{
		Name:"Pitsanulok",
		RegionID:1,
	},
	models.Province{
		Name:"Pichit",
		RegionID:1,
	},
	models.Province{
		Name:"Nakornsawan",
		RegionID:1,
	},
	models.Province{
		Name:"Uthaitani",
		RegionID:1,
	},
	models.Province{
		Name:"Nakornratchasima",
		RegionID:3,
	},
	models.Province{
		Name:"Chaiyapoom",
		RegionID:3,
	},
	models.Province{
		Name:"Loei",
		RegionID:3,
	},
	models.Province{
		Name:"Khonken",
		RegionID:3,
	},
	models.Province{
		Name:"Nongbualumpoo",
		RegionID:3,
	},
	models.Province{
		Name:"Udonthani",
		RegionID:3,
	},
	models.Province{
		Name:"Nongkhai",
		RegionID:3,
	},
	models.Province{
		Name:"Mahasarakham",
		RegionID:3,
	},
	models.Province{
		Name:"Kalasin",
		RegionID:3,
	},
	models.Province{
		Name:"Roiet",
		RegionID:3,
	},
	models.Province{
		Name:"Yasothorn",
		RegionID:3,
	},
	models.Province{
		Name:"Sakonnakorn",
		RegionID:3,
	},
	models.Province{
		Name:"Nakornpanom",
		RegionID:3,
	},
	models.Province{
		Name:"Mukdaharn",
		RegionID:3,
	},
	models.Province{
		Name:"Bungkarn",
		RegionID:3,
	},
	models.Province{
		Name:"Amnatjaroen",
		RegionID:3,
	},
	models.Province{
		Name:"Ubonratchatani",
		RegionID:3,
	},
	models.Province{
		Name:"Srisaket",
		RegionID:3,
	},
	models.Province{
		Name:"Surin",
		RegionID:3,
	},
	models.Province{
		Name:"Burirum",
		RegionID:3,
	},
	models.Province{
		Name:"Chumporn",
		RegionID:4,
	},
	models.Province{
		Name:"Suratthani",
		RegionID:4,
	},
	models.Province{
		Name:"Nakornsrithamarat",
		RegionID:4,
	},
	models.Province{
		Name:"Pattalung",
		RegionID:4,
	},
	models.Province{
		Name:"Songkha",
		RegionID:4,
	},
	models.Province{
		Name:"Pattani",
		RegionID:4,
	},
	models.Province{
		Name:"Yala",
		RegionID:4,
	},
	models.Province{
		Name:"Naratiwat",
		RegionID:4,
	},
	models.Province{
		Name:"Ranong",
		RegionID:4,
	},
	models.Province{
		Name:"Pangnga",
		RegionID:4,
	},
	models.Province{
		Name:"Phuket",
		RegionID:4,
	},
	models.Province{
		Name:"Krabi",
		RegionID:4,
	},
	models.Province{
		Name:"Trang",
		RegionID:4,
	},
	models.Province{
		Name:"Satoon",
	},
}



func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists( &models.User{},&models.Region{},&models.Province{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{},&models.Region{},&models.Province{}).Error
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
	for i, _ := range provinces{
		err = db.Debug().Model(&models.Province{}).Create(&provinces[i]).Error
		if err != nil {
			log.Fatalf("cannot seed province table: %v",err)
		}
	}
}