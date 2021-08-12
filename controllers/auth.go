package controllers

import (
	"gintest/config"
	"gintest/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strings"
	"time"
)

func UserLogin(ctx *gin.Context) {
	data := &models.UserLogin{}
	if err := ctx.ShouldBind(data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusOK, gin.H{
				"error": models.LoginError(err.(validator.ValidationErrors)),
			})
		}

		return
	}
	//log.Println(data)
	claims := &jwt.StandardClaims{
		Audience:  data.Username,
		ExpiresAt: time.Now().Add(10 * 60 * time.Second).Unix(),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	sign, err := token.SignedString([]byte(config.App.SigninKey))
	if err != nil {
		ctx.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"error": "token 签名失败",
		})
	}
	//log.Println(sign)
	ctx.JSON(http.StatusOK, gin.H{
		"token": sign,
		"error": "",
	})

}

func CheckAuth(ctx *gin.Context) {
	headers := ctx.GetHeader("Authorization")
	log.Println(headers)
	if haspre := strings.HasPrefix(headers, "Bearer "); !haspre {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "error",
		})
		return
	}
	if len(headers) <= 7 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "have no authorization",
		})
		return
	}
	tokenStr := headers[7:]

	//token := sha256.Sum256([]byte("I am a token"))
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.App.SigninKey), nil
	})
	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "token is not available",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
	})
}
