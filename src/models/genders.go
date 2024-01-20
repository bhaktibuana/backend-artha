package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Genders struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(6);not null" json:"name"`
}

func SeedGenders(db *gorm.DB) {
	var count int64
	db.Model(&Genders{}).Count(&count)

	if count == 0 {
		genders := []Genders{
			{
				Name: "Male",
			},
			{
				Name: "Female",
			},
		}

		for _, gender := range genders {
			if err := db.Create(&gender).Error; err != nil {
				fmt.Println("Error seeding gender:", err)
			}
		}
	}
}
