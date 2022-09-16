package databases

import (
	"go-backend-crowdfunding/configs"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func MySQLConnect() {
	dsn := configs.DB_USER + ":" + configs.DB_PASSWORD + "@tcp(" + configs.DB_HOST + ":" + configs.DB_PORT + ")/" + configs.DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("DB Connection Error: ")
	} else {
		log.Println("DB Connection Success")
	}

	DBMigrate()
}
