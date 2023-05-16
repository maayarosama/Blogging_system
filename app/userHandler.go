package app

import (
	"encoding/json"
	"net/http"

	"github.com/maayarosama/Blogging_system/internal"
	"github.com/maayarosama/Blogging_system/models"

	"github.com/rs/zerolog/log"
)

// Retrieve all users from db
func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	newUsers := a.db.GetUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// verify user
func (a *App) VerifyEmail(w http.ResponseWriter, r *http.Request) {

	type VerifyUserInput struct {
		Email string `json:"email"  binding:"required"`
		Code  int    `json:"code"  binding:"required"`
	}

	verifyData := &VerifyUserInput{}
	ParseBody(r, verifyData)

	user, EmailErr := a.db.GetUserByEmail(verifyData.Email)

	if EmailErr != nil {
		log.Error().Err(EmailErr).Send()
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user doesn't exists \n"))
		return
	}

	if EmailErr == nil && user.Verified {
		log.Error().Err(EmailErr).Send()
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user Already exists \n"))
		return
	}

	if verifyData.Code != user.VerificationCode {
		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte("codes don't match \n"))
		return

	}
	user.Verified = true
	err := a.db.UpdateUserVerfied(user)
	if err != nil {
		log.Error().Err(err).Send()
		w.Write([]byte("error while verifying\n"))
		return
	}

	w.Write([]byte("account verified successfully \n"))

}

// sign up and sending verfication code
func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {

	createUser := &models.User{}
	ParseBody(r, createUser)

	// Check if user exists and verified
	user, EmailErr := a.db.GetUserByEmail(createUser.Email)
	if EmailErr == nil && user.Verified {
		log.Error().Err(EmailErr).Send()

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user Already exists \n"))
		return
	}

	// Generate verfication code and set it to the created user
	code := internal.GenerateRandomCode()
	message := internal.SignUpMailBody(code, a.config.MailSender.Timeout)

	createUser.Password, _ = internal.HashPassword(user.Password)

	err := internal.SendMail(a.config.MailSender.Email, a.config.MailSender.Password, createUser.Email, message)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	createUser.VerificationCode = code

	// Check if user exists but not verified if so update the user's info
	if EmailErr == nil && !user.Verified {
		createUser.ID = user.ID
		err = a.db.UpdateUser(createUser)
		if err != nil {
			log.Error().Err(err).Send()
			return
		}
		w.Write([]byte("user exists but not verified another verification code has been sent to your email \n"))
		return
	}

	// user doesn't exist, so create new user
	u := a.db.CreateUser(createUser)
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	w.Write([]byte("Verification code has been sent to your email \n"))
	return

}

// user sign in
func (a *App) SignIn(w http.ResponseWriter, r *http.Request) {
	type SignInInput struct {
		Email    string `json:"email"  binding:"required"`
		Password string `json:"password"  binding:"required"`
	}

	input := &SignInInput{}
	ParseBody(r, input)
	userDetails, err := a.db.GetUserByEmail(input.Email)

	if err != nil {
		log.Error().Err(err).Send()

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no user found\n"))
		return
	}

	if err == nil && !userDetails.Verified {
		log.Error().Err(err).Send()

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user isn't verified\n"))

		return
	}

	// compare := internal.VerifyPassword(userDetails.Password, input.Password)
	// fmt.Printf("%+v\n", []byte(userDetails.Password))
	// fmt.Printf("%+v\n", []byte(input.Password))
	// compare := bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(input.Password))
	// fmt.Printf("userDetails.Password %+v\n", userDetails.Password)
	// fmt.Printf("input.Password %+v\n", input.Password)
	// p, _ := internal.HashPassword(input.Password)

	// fmt.Printf("Hashed input.Password %+v\n", p)

	// if compare != nil {
	// 	log.Error().Err(err).Send()
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	log.Printf("failed signin: %v", compare)
	// 	w.Write([]byte(compare.Error()))

	// 	return
	// }

	// res, _ := json.Marshal(userDetails)
	token, err := models.GenerateToken(userDetails.ID, userDetails.Email, a.config.Token.Secret, a.config.Token.Timeout)

	data, err := json.Marshal(map[string]string{"access_token": token})
	// r.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	// reqToken := r.Header.Get("Authorization")
	// splitToken := strings.Split(reqToken, "Bearer ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.Write([]byte("Signed in successfully \n"))

}
