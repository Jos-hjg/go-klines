package models

import "github.com/go-playground/validator/v10"

type UserLogin struct {
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}

var LoginValidation = ErrorType{
	"Username":{
		"required":"用户名称必须",

	},
	"Password":{
		"required":"密码必须",
	},
}

func LoginError(errs validator.ValidationErrors) map[string]string{
	errMap := map[string]string{}
	for _, err := range errs {
		errMap[err.Field()] = LoginValidation[err.Field()][err.Tag()]
	}
	return errMap
}