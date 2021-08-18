package config

import (
	"books-api/entity"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitializeDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err := gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	db.AutoMigrate(&entity.Book{})

	return db
}
