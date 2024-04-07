package module

import (
	"bybu/go-postgres/shared/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = ConnectToPostgres()

func ConnectToPostgres() (*gorm.DB) {
	dsn := config.Env.GetPostgressUrl();
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db;
}
