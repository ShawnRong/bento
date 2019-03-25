package controllers

import (
	"net/http"

	"github.com/ShawnRong/bento/db"

	"github.com/ShawnRong/bento/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var user models.User

func (u UserController) Retrieve(c *gin.Context) {
	db.GetDB().First(&user, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"users":  user,
	})
}

func (u UserController) Create(c *gin.Context) {
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (u UserController) Delete(c *gin.Context) {
	if err := db.GetDB().Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "delete success",
	})
}

func (u UserController) Update(c *gin.Context) {
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.GetDB().Model(&user).Updates(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
