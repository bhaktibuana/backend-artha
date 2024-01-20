package models

import "time"

type ModelTemplate struct {
	CreatedAt     time.Time `gorm:"type:datetime;not null" json:"created_at"`
	CreatedBy     uint64    `gorm:"type:integer;not null" json:"created_by"`
	CreatedByUser Users     `gorm:"foreignKey:CreatedBy"`
	UpdatedAt     time.Time `gorm:"type:datetime" json:"updated_at"`
	UpdatedBy     uint64    `gorm:"type:integer" json:"updated_by"`
	UpdatedByUser Users     `gorm:"foreignKey:UpdatedBy"`
	DeletedAt     time.Time `gorm:"type:datetime" json:"deleted_at"`
	DeletedBy     uint64    `gorm:"type:integer" json:"deleted_by"`
	DeletedByUser Users     `gorm:"foreignKey:DeletedBy"`
}
