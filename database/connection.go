package database

import(
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/Sohbetbackend/golang-react/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=backend2003 dbname=golang-react-postgres port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}