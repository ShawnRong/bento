package controllers

import (
	"net/http"

	"github.com/ShawnRong/bento/e"
	"github.com/ShawnRong/bento/models"
	"github.com/ShawnRong/bento/util"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid: "Required; MaxSize(50)"`
	Password string `valid: "Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//a := auth{Username: username, Password: password}

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	isExist := models.CheckAuth(username, password)
	if isExist {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = e.SUCCESS
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
