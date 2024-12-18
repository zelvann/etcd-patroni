package product

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	ImagePath   string    `json:"image_path"`
	ExpiryDate  *time.Time `json:"expiry_date,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	ExpiryDate  *time.Time `json:"expiry_date,omitempty"`
}

type UpdateProductRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ExpiryDate  *time.Time `json:"expiry_date,omitempty"`
}
