package router

import (
	"gintest/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	r := gin.Default()
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:5000"}
	//r.Use(cors.New(config))
	r.Use(cors.Default())   //使用默认cors
	r.GET("/products",controllers.GetProduct)
	r.POST("/product",controllers.AddProduct)
	r.DELETE("/product",controllers.DelProduct)
	r.GET("/records",controllers.GetRecord)
	r.GET("/record",controllers.GetNewlyRecord)
	r.POST("/user-login",controllers.UserLogin)
	return r
}