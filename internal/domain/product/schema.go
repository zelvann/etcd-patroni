package product

import (
	"time"

	"gorm.io/gorm"
)

type Timestamp struct {
	CreatedAt time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

type ProductSchema struct {
	ID          uint      `json:"id" gorm:"null;primaryKey;autoIncrement" `
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ExpireDate  time.Time `json:"expire_date"`
	ImageName   string    `json:"image_name"`

	Timestamp
}
