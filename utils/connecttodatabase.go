package utils

import (
	"graphQLGEN/graph/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDB() {
	var err error
	db, err = gorm.Open(postgres.Open("host=localhost user=postgres dbname=GraphQL-Jokes sslmode=disable password=baki@Hanma"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&model.Joke{})
}
func GetDB() *gorm.DB {
	return db
}
