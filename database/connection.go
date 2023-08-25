package database

import(
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)


func Connect() {
	_, err := gorm.Open(sqlite.Open(""), &gorm.Config{})
	
	if err != nil {
		panic("could not connect to the database")
	}
}