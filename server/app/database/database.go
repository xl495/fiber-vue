package database

import (
	"fiber-vue/app/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		username, password, host, db,
	)

	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(&model.User{}, &model.Menu{})
	if err != nil {
		log.Fatal(err)
	}

}
