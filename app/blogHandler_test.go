package app

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maayarosama/Blogging_system/internal"
	"github.com/maayarosama/Blogging_system/models"
	"github.com/stretchr/testify/assert"
)

func TestGetBlogs(t *testing.T) {
	a := SetUp(t)

	err := CreateNewBlog(a)
	if err != nil {
		t.Errorf("%+v", err)
	}

	request := httptest.NewRequest("GET", "/blogs", nil)
	response := httptest.NewRecorder()
	a.GetBlogs(response, request)
	assert.Equal(t, response.Code, http.StatusOK)

}

func TestCreateBlog(t *testing.T) {

	body := []byte(`{
		"title":"ossamam@incubaid.com",
		"content":"123456789"
	}`)
	a := SetUp(t)
	err := CreateNewUser(a)
	if err != nil {
		t.Errorf("%+v", err)
	}

	user, err := a.DB.GetUserByEmail("mayar@incubaid.com")
	assert.NoError(t, err)
	token, _ := models.GenerateToken(user.ID, user.Email, a.Config.Token.Secret, a.Config.Token.Timeout)

	request := httptest.NewRequest("POST", "/blog", bytes.NewBuffer(body))

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	response := httptest.NewRecorder()

	ctx := context.WithValue(request.Context(), internal.UserIDKey("UserID"), user.ID)
	newRequest := request.WithContext(ctx)
	a.CreateBlog(response, newRequest)
	assert.Equal(t, response.Code, http.StatusOK)

}

func CreateNewBlog(a *App) error {

	body := []byte(`{
		"title":"ossamam@incubaid.com",
		"content":"123456789"
	}`)
	err := CreateNewUser(a)
	if err != nil {
		return err
	}

	user, err := a.DB.GetUserByEmail("mayar@incubaid.com")
	if err != nil {
		return err
	}
	token, _ := models.GenerateToken(user.ID, user.Email, a.Config.Token.Secret, a.Config.Token.Timeout)

	request := httptest.NewRequest("POST", "/blog", bytes.NewBuffer(body))

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	response := httptest.NewRecorder()

	ctx := context.WithValue(request.Context(), internal.UserIDKey("UserID"), user.ID)
	newRequest := request.WithContext(ctx)
	a.CreateBlog(response, newRequest)
	return nil
}
