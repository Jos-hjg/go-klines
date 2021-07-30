package database

import (
	"gintest/config"
	"gintest/controllers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func InitDB() (*gorm.DB, error){

	db, err := gorm.Open("mysql", config.App.DSN)
	if err != nil {
		return nil, err
	}
	controllers.DB = db
	return db, nil
}

