package app

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/maayarosama/Blogging_system/models"
	"github.com/stretchr/testify/assert"
)

func SetUp(t testing.TB) (a *App) {
	config := `{
		"server": {
			"host": "localhost",
			"port": ":3000"
		},
		"database": {
			"file": "./test.db"
		},
		"token": {
			"secret": "secret",
			"timeout": 200
		},
		"mailSender": {
		    "email" : "ossamam@incubaid.com",
        	"password": "Nirvana2021",
			"timeout": 10
		}

	}
	`
	dir := t.TempDir()
	configPath := dir + "/config.json"
	dbPath := dir + "testing.db"

	err := os.WriteFile(configPath, []byte(config), 0644)
	if err != nil {
		t.Errorf("Should've writteb to file, but got the following error:  %+v", err)
	}

	a, err = NewApp(configPath)
	db, _ := a.InitiateDB(dbPath)
	a.DB = db
	if err != nil {
		t.Errorf("Should've created a new app, but got the following error:  %+v", err)
	}

	return a
}

func CreateNewUser(a *App) error {
	body := []byte(`{
		"name":"mayarosama22",
		"email":"mayar@incubaid.com",
		"password":"12345678910",
		"quote":"sssssssss"
	}`)
	request := httptest.NewRequest("GET", "/user/signup", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	a.SignUp(response, request)

	got := response.Body.String()

	want := "Verification code has been sent to your email \n"

	if got != want {
		err := fmt.Sprintf(`
		error : got %q, want %q
	`, got, want)
		return errors.New(err)

	}

	return nil
}

func VerifyUser(a *App) error {
	code, err := a.DB.GetCodeByEmail("mayar@incubaid.com")
	println("err ", err)

	if err != nil {
		return err
	}

	verifyuser := fmt.Sprintf(`{
		"email":"mayar@incubaid.com",
		"code": %d
	}`, code)

	request := httptest.NewRequest("POST", "/user/verifyemail", bytes.NewBuffer([]byte(verifyuser)))
	response := httptest.NewRecorder()
	a.VerifyEmail(response, request)

	got := response.Body.String()
	want := "account verified successfully \n"
	if got != want {
		err := fmt.Sprintf(`
		error : got %q, want %q
	`, got, want)
		return errors.New(err)

	}

	return nil
}

func SignInTest(a *App) {

	body := []byte(`{
		"email":"mayar@incubaid.com",
		"password":"12345678910"
	}`)

	request := httptest.NewRequest("POST", "/user/signin", bytes.NewBuffer([]byte(body)))
	response := httptest.NewRecorder()
	a.SignIn(response, request)
}

func TestSignUp(t *testing.T) {
	a := SetUp(t)
	body := []byte(`{
		"name":"mayarosama22",
		"email":"mayar1701976@miuegypt.edu.eg",
		"password":"1234567891011",
		"quote":"sssssssss"
	}`)
	t.Run("sign up ", func(t *testing.T) {

		request := httptest.NewRequest("POST", "/user/signup", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		a.SignUp(response, request)

		got := response.Body.String()
		want := `{Verification code has been sent to your email \n}`

		if got == want {
			t.Errorf("error : got %q, want %q", got, want)
		}

	})

	t.Run("user already exists but not verified", func(t *testing.T) {

		a.DB.CreateUser(
			&models.User{Name: "mayarosama22", Email: "mayar1701976@miuegypt.edu.eg", Password: "$2a$04$3WF5ZN8c5OXmFwbj8oFsN.BdRvJRDUt8zfP0vS2A7Zzx6K2rMrmg.", Quote: "desc", VerificationCode: 999, Verified: false})

		request := httptest.NewRequest("POST", "/user/signup", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		a.SignUp(response, request)

		got := response.Body.String()
		want := `{user exists but not verified another verification code has been sent to your email \n}`
		if got == want {
			t.Errorf("error : got %q, want %q", got, want)
		}
	})

}

func TestGetUsers(t *testing.T) {
	a := SetUp(t)

	err := CreateNewUser(a)
	if err != nil {
		t.Errorf("%+v", err)
	}

	request := httptest.NewRequest("POST", "/user", nil)
	response := httptest.NewRecorder()
	a.GetUsers(response, request)
	assert.Equal(t, response.Code, http.StatusOK)

}

func TestVerifyEmail(t *testing.T) {
	a := SetUp(t)

	err := CreateNewUser(a)
	if err != nil {
		t.Errorf("%+v", err)
	}

	t.Run("verify incorrect signup code", func(t *testing.T) {

		verifyuser := fmt.Sprintf(`{
			"email":"mayar@incubaid.com",
			"code": %d
		}`, 5555)

		request := httptest.NewRequest("POST", "/user/verifyemail", bytes.NewBuffer([]byte(verifyuser)))
		response := httptest.NewRecorder()
		a.VerifyEmail(response, request)

		got := response.Body.String()
		want := "codes don't match \n"

		if got != want {
			t.Errorf("error : got %q, want %q", got, want)

		}

	})

	t.Run("verify correct signup code", func(t *testing.T) {

		code, err := a.DB.GetCodeByEmail("mayar@incubaid.com")

		if err != nil {
			t.Errorf("error : %q", err)

		}
		verifyuser := fmt.Sprintf(`{
			"email":"mayar@incubaid.com",
			"code": %d
		}`, code)

		request := httptest.NewRequest("POST", "/user/verifyemail", bytes.NewBuffer([]byte(verifyuser)))
		response := httptest.NewRecorder()
		a.VerifyEmail(response, request)

		got := response.Body.String()
		want := "account verified successfully \n"

		if got != want {
			t.Errorf("error : got %q, want %q", got, want)

		}

	})

	t.Run("verify already verified user", func(t *testing.T) {

		code, err := a.DB.GetCodeByEmail("mayar@incubaid.com")

		if err != nil {
			t.Errorf("error : %q", err)

		}
		verifyuser := fmt.Sprintf(`{
			"email":"mayar@incubaid.com",
			"code": %d
		}`, code)

		request := httptest.NewRequest("POST", "/user/verifyemail", bytes.NewBuffer([]byte(verifyuser)))
		response := httptest.NewRecorder()
		a.VerifyEmail(response, request)

		got := response.Body.String()
		want := "user Already exists \n"

		if got != want {
			t.Errorf("error : got %q, want %q", got, want)

		}

	})

}

func TestSignIn(t *testing.T) {
	a := SetUp(t)
	t.Run("signin with a user that doesn't exist", func(t *testing.T) {

		body := []byte(`{
			"email":"osm@incubaid.com",
			"password":"12345678910"
		}`)

		request := httptest.NewRequest("POST", "/user/signin", bytes.NewBuffer([]byte(body)))
		response := httptest.NewRecorder()
		a.VerifyEmail(response, request)

		got := response.Body.String()
		want := "user doesn't exists \n"

		if got != want {
			t.Errorf("error : got %q, want %q", got, want)

		}

	})

	err := CreateNewUser(a)
	if err != nil {
		t.Errorf("%+v", err)
	}

	t.Run("signin not verified user and correct password", func(t *testing.T) {

		body := []byte(`{
			"email":"mayar@incubaid.com",
			"password":"12345678910"
		}`)

		request := httptest.NewRequest("POST", "/user/signin", bytes.NewBuffer([]byte(body)))
		response := httptest.NewRecorder()
		a.SignIn(response, request)

		got := response.Body.String()
		want := "user isn't verified\n"

		if got != want {
			t.Errorf("error : got %q, want %q", got, want)

		}

	})

	err = VerifyUser(a)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	t.Run("signin verified user and correct password", func(t *testing.T) {

		body := []byte(`{
			"email":"mayar@incubaid.com",
			"password":"12345678910"
		}`)

		request := httptest.NewRequest("POST", "/user/signin", bytes.NewBuffer([]byte(body)))
		response := httptest.NewRecorder()
		a.SignIn(response, request)
		assert.Equal(t, response.Code, http.StatusOK)

	})

	t.Run("signin verified user and incorrect password", func(t *testing.T) {

		body := []byte(`{
			"email":"mayar@incubaid.com",
			"password":"12345910"
		}`)

		request := httptest.NewRequest("POST", "/user/signin", bytes.NewBuffer([]byte(body)))
		response := httptest.NewRecorder()
		a.SignIn(response, request)
		got := response.Body.String()
		want := "crypto/bcrypt: hashedPassword is not the hash of the given password"

		if got != want {
			t.Errorf("error : got %q, want %q", got, want)

		}

	})

}
