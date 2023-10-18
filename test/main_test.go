package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestSearchUserByIdHandler(t *testing.T) {
	database.ConnectDatabase()
	CreateUserMocked()
	defer DeleteUserMocked()
	r := SetupTestRoutes()
	r.GET("users/:id", controller.GetUserById)
	searchPath := "/users/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", searchPath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var mockUser user_model.UserModel
	json.Unmarshal(response.Body.Bytes(), &mockUser)

	log.Default().Println(mockUser)

	assert.Equal(t, "Raveline", mockUser.Name, "Should be equal")
	assert.Equal(t, "Seak for the truth", mockUser.Description)
	assert.Equal(t, "1212121212", mockUser.SocialNumber)
	assert.Equal(t, http.StatusOK, response.Code)

}

func TestDeleteUser(t *testing.T) {
	database.ConnectDatabase()
	CreateUserMocked()
	r := SetupTestRoutes()
	r.DELETE("users/:id", controller.DeleteUser)
	deletePath := "/users/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditUser(t *testing.T) {
	database.ConnectDatabase()
	CreateUserMocked()
	defer DeleteUserMocked()
	r := SetupTestRoutes()
	r.PATCH("users/:id", controller.EditUser)
	user := user_model.UserModel{
		Name:         "Raveline Junior",
		Description:  "Seak for the truth, even if it hurts",
		SocialNumber: "269498421.22",
	}

	jsonValue, _ := json.Marshal(user)
	editPath := "/users/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("PATCH", editPath, bytes.NewBuffer(jsonValue))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var mockedUser user_model.UserModel
	json.Unmarshal(response.Body.Bytes(), &mockedUser)

	log.Default().Println(mockedUser)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Raveline Junior", user.Name)
	assert.Equal(t, "Seak for the truth, even if it hurts", user.Description)
	assert.Equal(t, "269498421.22", user.SocialNumber)
}
