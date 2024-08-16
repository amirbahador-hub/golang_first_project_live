package shop 

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id int `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Price int `json:"price"`
	BrandID int `json:"brand_id"`
	Brand Brand `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:BrandID"`
	CategoryID int `json:"category_id"`
	Category Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CategoryID"`
	DiscountPercent float32 `json:"discount_percent"`
	DiscountedPrice int `json:"discounted_price"`
}

type Brand struct {
	gorm.Model
	Id int `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
}

type Category struct {
	gorm.Model
	Id int `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Parent int `json:"parent"`
}