package repositories

import (
	"fmt"
	"log"

	"github.com/ttnsp/go-boilerplate/auth"
	"github.com/ttnsp/go-boilerplate/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(host string, port int, dbname string, user string, password string) {

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=Europe/Paris", host, port, dbname, user, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Foo{}, &models.Users{})

	// create user admin with default password if not exists yet
	var exists bool
	db.Model(models.Users{}).Select("count(*) > 0").Where("name = ?", "admin").Find(&exists)
	if !exists {
		password, err := auth.GeneratehashPassword("admin")
		if err != nil {
			log.Fatal(err)
		}
		var admin models.Users = models.Users{Name: "admin", Password: password}
		db.Model(models.Users{}).Create(&admin)
		log.Printf("%v", admin)
	}

	DB = db
}
