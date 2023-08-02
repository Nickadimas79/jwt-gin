package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	DBdriver := os.Getenv("DB_DRIVER")
	DBhost := os.Getenv("DB_HOST")
	DBuser := os.Getenv("DB_USER")
	DBpassword := os.Getenv("DB_PASSWORD")
	DBname := os.Getenv("DB_NAME")
	DBport := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBuser, DBpassword, DBhost, DBport, DBname)

	DB, err := gorm.Open(DBdriver, DBURL)
	if err != nil {
		fmt.Println("cannot connect to database:", DBdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("connected to database ", DBdriver)
	}

	DB.AutoMigrate(&User{})
}
