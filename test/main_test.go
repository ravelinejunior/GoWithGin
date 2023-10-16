package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ravelinejunior/go_api_gin/controller"
	"github.com/ravelinejunior/go_api_gin/database"
	user_model "github.com/ravelinejunior/go_api_gin/model"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

var ID int

func CreateUserMocked() {
	user := user_model.UserModel{
		Name:         "Raveline",
		Description:  "Seak for the truth",
		SocialNumber: "1212121212",
	}

	database.DB.Create(&user)
	ID = int(user.ID)
}

func DeleteUserMocked() {
	var user user_model.UserModel
	database.DB.Delete(&user, ID)
}

func TestVerifyStatusCodeFromGreetingWithParam(t *testing.T) {

	r := SetupTestRoutes()
	r.GET("/:name", controller.Greetings)
	req, _ := http.NewRequest("GET", "/raveline", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Must be equals")
	responseMock := `{"Greetings":"Hello pretty soul of raveline, how are you?"}`
	responseBody, _ := io.ReadAll(res.Body)
	assert.Equal(t, responseMock, string(responseBody))

}

func TestListingAllUsersWithHandlers(t *testing.T) {
	database.ConnectDatabase()
	CreateUserMocked()
	defer DeleteUserMocked()
	r := SetupTestRoutes()
	r.GET("/users", controller.ShowAllUsers)
	req, err := http.NewRequest("GET", "/users", nil)
	if assert.NoError(t, err) {
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	}
}

func TestSearchUserBySocialNumberHandler(t *testing.T) {
	database.ConnectDatabase()
	CreateUserMocked()
	defer DeleteUserMocked()
	r := SetupTestRoutes()
	r.GET("/users/social/:social_number", controller.FindUserBySocialNumber)
	req, _ := http.NewRequest("GET", "/users/social/1212121212", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}
