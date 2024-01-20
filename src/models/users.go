package models

import (
	"api-artha/src/helpers"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int64          `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Username  string         `gorm:"type:varchar(16);not null" json:"username"`
	TagLine   string         `gorm:"type:varchar(5);not null" json:"tag_line"`
	Email     string         `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password"`
	BirthDate sql.NullTime   `gorm:"type:date" json:"birth_date"`
	GenderId  sql.NullInt64  `gorm:"type:integer" json:"gender_id"`
	Gender    Genders        `gorm:"foreignKey:GenderId"`
	RoleId    int64          `gorm:"type:integer;not null" json:"role_id"`
	Role      Roles          `gorm:"foreignKey:RoleId"`
	Status    string         `gorm:"type:varchar(20);check:status IN ('verified', 'unverified', 'deleted');not null" json:"status"`
	ImageUrl  sql.NullString `gorm:"type:varchar(255)" json:"image_url"`
	CreatedAt time.Time      `gorm:"type:datetime;not null" json:"created_at"`
	UpdatedAt sql.NullTime   `gorm:"type:datetime" json:"updated_at"`
	DeletedAt sql.NullTime   `gorm:"type:datetime" json:"deleted_at"`
}

func SeedUsers(db *gorm.DB) {
	var count int64
	db.Model(&Users{}).Count(&count)

	if count == 0 {
		var role Roles

		if roleError := db.First(&role, "name = ?", "Super Admin").Error; roleError != nil {
			switch roleError {
			case gorm.ErrRecordNotFound:
				fmt.Println("No role found with the specified name")
				return
			default:
				fmt.Println("Error:", roleError)
				return
			}
		}

		users := []Users{
			{
				Name:      "Artha Admin",
				Username:  "Artha",
				TagLine:   "00000",
				Email:     "admin@artha.bhaktibuana.com",
				Password:  helpers.HashPassword("@rthA1234567890"),
				BirthDate: sql.NullTime{time.Time{}, false},
				GenderId:  sql.NullInt64{Int64: 0, Valid: false},
				RoleId:    role.Id,
				Status:    "verified",
				CreatedAt: time.Now(),
				UpdatedAt: sql.NullTime{time.Time{}, false},
				DeletedAt: sql.NullTime{time.Time{}, false},
			},
		}

		for _, user := range users {
			if err := db.Create(&user).Error; err != nil {
				fmt.Println("Error seeding user:", err)
			}
		}
	}
}
