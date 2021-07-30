package models

import "github.com/jinzhu/gorm"

type Record struct {
	gorm.Model
	MicroTime float64
	Open float64
	Highest float64
	Lowest float64
	Close float64
	Volume float64
}

type PartOfRecord struct {
	MicroTime float64
	Open float64
	Highest float64
	Lowest float64
	Close float64
	Volume float64
}
