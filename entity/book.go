package entity

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Book struct {
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Author      string `gorm:"type:varchar(255)" json:"author"`
	Description string `gorm:"type:text" json:"description"`

	gorm.Model
}
