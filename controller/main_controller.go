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

	if err := user_model.ValidateUserData(&user); err != nil {
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

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Fail. The user with id " + id + " doesnt exist",
		})
		return
	}

	err := database.DB.Where("id = ?", id).Delete(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func EditUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user user_model.UserModel
	database.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Fail. The user with id " + id + " doesnt exist",
		})

		return
	}

	if err := user_model.ValidateUserData(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// if record is empty return error message to the client
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Model(&user).UpdateColumns(user)
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
			"message": "Fail. The user with id " + id + " doesnt exist",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func FindUserBySocialNumber(c *gin.Context) {
	var user user_model.UserModel
	socialNumber := c.Param("social_number")

	database.DB.Where(&user_model.UserModel{SocialNumber: socialNumber}).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Fail. The user with social number " + socialNumber + " doesnt exist",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

func Greetings(c *gin.Context) {

	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"Greetings": "Hello pretty soul of " + name + ", how are you?",
	})
}
