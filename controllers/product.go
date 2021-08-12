package controllers

import (
	"gintest/models"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"strconv"
)

func GetProduct(ctx *gin.Context) {
	var products []models.Product
	limit := 3
	total := 0
	offset := 0
	count := 0
	prev := 0
	next := 0
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	offset = (page - 1) * limit
	DB.Model(&models.Product{}).Count(&count)
	total = int(math.Ceil(float64(count) / float64(limit)))
	if err != nil || page <= 0 || page > total {
		page = 1
	}
	DB.
		Offset(offset).Limit(limit).
		Find(&products)
	if page == 1 {
		prev = total
	} else {
		prev = page - 1
	}
	if page == total {
		next = 1
	} else {
		next = page + 1
	}
	pager := map[string]interface{}{
		"limit": limit,
		"total": total,
		"page":  page,
		"prev":  prev,
		"next":  next,
		"rows":  len(products),
	}
	ctx.JSON(http.StatusOK, gin.H{
		"pager": pager,
		"data":  products,
		"error": "",
	})
}

func AddProduct(ctx *gin.Context) {

	var product models.Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  "",
			"error": err,
		})
		log.Fatal(err)
	}
	DB.Create(&product)
	ctx.JSON(http.StatusOK, gin.H{
		"data":  product,
		"error": "",
	})
}

func DelProduct(ctx *gin.Context) {
	var product models.Product
	var products []models.Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  "",
			"error": err,
		})
		log.Fatal(err)
	}
	DB.Delete(&products, product)
	ctx.JSON(http.StatusOK, gin.H{
		"data":  products,
		"error": "",
	})
}
