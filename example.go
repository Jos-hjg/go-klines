package main

import (
	"gintest/config"
	"gintest/database"
	"gintest/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	//读取配置文件
	config.InitConf()
	//初始化 DB
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := router.InitRouter()
	r.Run(config.App.PORT)
}
