package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ravelinejunior/go_api_gin/database"
	user_model "github.com/ravelinejunior/go_api_gin/model"
)

func CreateNewUser(c *gin.Context) {
	var user user_model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user user_model.UserModel
	err := database.DB.Where("id = ?", id).Delete(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Record not found",
		})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Result": "Fail. The user with id " + id + " doesnt exist",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func ShowAllUsers(c *gin.Context) {
	var users []user_model.UserModel
	database.DB.Find(&users)
	c.JSON(200, users)
}

func GetUserById(c *gin.Context) {
	var user user_model.UserModel
	id := c.Params.ByName("id")
	database.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Result": "Fail. The user with id" + id + "doesnt exist",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func Greetings(c *gin.Context) {

	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"Greetings ": "How are you " + name,
	})
}
