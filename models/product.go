package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code string `gorm:"unique_index" form:"code"`
	Title string `gorm:"index" form:"title"`
	Price int `form:"price"`
	Sales int `form:"sales"`
}

