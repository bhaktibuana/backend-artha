package models

import (
	"api-artha/src/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", configs.DbConfig().DbUser, configs.DbConfig().DbPassword, configs.DbConfig().DbHost, configs.DbConfig().DbPort, configs.DbConfig().DbName)
	database, err := gorm.Open(mysql.Open(connection))

	if err != nil {
		panic(err)
	}

	MigrateDatabase(database)
	SeedDatabase(database)

	DB = database
}

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(
		&Genders{},
		&Roles{},
		&Users{},
	)
}

func SeedDatabase(db *gorm.DB) {
	SeedGenders(db)
	SeedRoles(db)
	SeedUsers(db)
}
