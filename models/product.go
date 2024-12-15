package models

type Product struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	Name             string
	Description      string
	Price            float64
	ProductImages    []string `gorm:"type:text[]"`
	CompressedImages []string `gorm:"type:text[]"`
}
