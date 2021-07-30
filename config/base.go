package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var App *Conf

type Conf struct {
	PORT string `json:"APP_PORT"`
	DSN string `json:"APP_CONNECTSTRING"`
}

func InitConf(){
	config := &Conf{}
	ConfFile := "./settings.json"
	Content, err := ioutil.ReadFile(ConfFile)
	if err != nil{
		log.Println("Use default settings")
		//使用默认配置
		App = &Conf{PORT: ":8080", DSN: "root:Huangjg1001@tcp(localhost:3306)/tank?charset=utf8&parseTime=true&loc=Local"}
		return
	}

	if err := json.Unmarshal(Content,config); err != nil{
		log.Println("Use default settings")
		//使用默认配置
		App = &Conf{PORT: ":8080", DSN: "root:Huangjg1001@tcp(localhost:3306)/tank?charset=utf8&parseTime=true&loc=Local"}
		return
	}

	App = config
}