package models


import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
)


type User struct{
	gorm.Models
	Name string
	Email string `gorm:"not null;unique_index"`
}