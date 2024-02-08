package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Roles struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(30);not null" json:"name"`
}

func SeedRoles(db *gorm.DB) {
	var count int64
	db.Model(&Roles{}).Count(&count)

	if count == 0 {
		roles := []Roles{
			{
				Name: "Super Admin",
			},
			{
				Name: "Common User",
			},
		}

		for _, role := range roles {
			if err := db.Create(&role).Error; err != nil {
				fmt.Println("Error seeding role:", err)
			}
		}
	}
}
