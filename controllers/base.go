package controllers

import (
	"gintest/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB

func GetProduct(ctx *gin.Context) {
	var products []models.Product
	
	filter := &models.Product{}
	if err := ctx.ShouldBindQuery(&filter); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"data":"",
			"error":err,
		})
		log.Fatal(err)
	}
	DB.Find(&products, filter)
	ctx.JSON(http.StatusOK,gin.H{
		"data":products,
		"error":"",
	})
}

func AddProduct(ctx *gin.Context)  {
	
	var product models.Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"data":"",
			"error":err,
		})
		log.Fatal(err)
	}
	DB.Create(&product)
	ctx.JSON(http.StatusOK,gin.H{
		"data":product,
		"error":"",
	})
}

func DelProduct(ctx *gin.Context){
	
	var product models.Product
	var products []models.Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"data":"",
			"error":err,
		})
		log.Fatal(err)
	}
	DB.Delete(&products,product)
	ctx.JSON(http.StatusOK,gin.H{
		"data":products,
		"error":"",
	})
}

func GetRecord(ctx *gin.Context){
	var records []models.Record

	
	DB.Find(&records)

	ctx.JSON(http.StatusOK,gin.H{
		"data":records,
		"success":true,
	})

}

func GetNewlyRecord(ctx *gin.Context){
	var record models.Record
	//从数据库获取所有 record 数据
	DB.Order("id DESC").First(&record)
	ctx.JSON(http.StatusOK, gin.H{
		"data": record,
		"success":true,
	})
}

func UserLogin(ctx *gin.Context)  {
	token := "I am a token"
	data := &models.UserLogin{}
	if err := ctx.ShouldBind(data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok{
			ctx.JSON(http.StatusOK,gin.H{
				"error":models.LoginError(err.(validator.ValidationErrors)),
			})
		}

		return
	}
	log.Println(data)
	ctx.JSON(http.StatusOK,gin.H{
		"token": token,
		"error":"",
	})
}


