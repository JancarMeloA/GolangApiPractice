package DB

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DbConnection() {
	dsn := "host=localhost user=postgres password=1234 dbname=OkaneDev port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	if err != nil {

		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
}
