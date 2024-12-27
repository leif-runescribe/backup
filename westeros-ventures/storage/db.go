// database/db.go
package database

import (
	"gorm.io/driver/postgres"

	"github.com/leif-runescribe/westeros-roster/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "user=postgres password=123 dbname=Houses host=localhost port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.User{})

	return DB
}
