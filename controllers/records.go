package controllers

import (
	"gintest/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRecord(ctx *gin.Context) {
	var records []models.Record
	DB.Find(&records)

	ctx.JSON(http.StatusOK, gin.H{
		"data":    records,
		"success": true,
	})
}

func GetNewlyRecord(ctx *gin.Context) {
	var record models.Record
	//从数据库获取所有 record 数据
	DB.Order("id DESC").First(&record)
	ctx.JSON(http.StatusOK, gin.H{
		"data":    record,
		"success": true,
	})
}
