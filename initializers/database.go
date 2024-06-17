package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

  var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=ep-plain-wood-a22ufdb7.eu-central-1.aws.neon.tech user='Sunset Gocrud_owner' password='iG7ZYAKNw1CI' dbname='Sunset Gocrud' port=5432 sslmode=require"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal("Failed to connect to Database")
	}
}
