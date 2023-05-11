package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/maayarosama/Blogging_system/models"
	"github.com/maayarosama/Blogging_system/utils"
	"github.com/rs/zerolog/log"
)

func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	newUsers := c.db.GetUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *Controller) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	verifyData := &models.VerifyUserInput{}

	utils.ParseBody(r, verifyData)
	user, EmailErr := c.db.GetUserByEmail(verifyData.Email)

	if EmailErr != nil {
		log.Error().Err(EmailErr).Send()
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User doeasn't exists \n"))
		return
	}

	if EmailErr == nil && user.Verified {
		log.Error().Err(EmailErr).Send()
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User Already exists \n"))
		return
	}

	if verifyData.Code != user.VerificationCode {
		w.Write([]byte("Codes don't match \n"))
		return

	}
	user.Verified = true
	err := c.db.UpdateUserVerfied(user)
	if err != nil {
		log.Error().Err(err).Send()
		w.Write([]byte("Error while verifying\n"))

		return
	}

	w.Write([]byte("Account verified successfully \n"))

}

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {

	createUser := &models.User{}
	utils.ParseBody(r, createUser)

	// Check if user exists and verified
	user, EmailErr := c.db.GetUserByEmail(createUser.Email)
	if EmailErr == nil && user.Verified {
		log.Error().Err(EmailErr).Send()

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User Already exists \n"))
		return
	}

	// Generate verfication code and set it to the created user
	code := utils.GenerateRandomCode()
	message := utils.SignUpMailBody(code, c.mailSender.Timeout)

	err := utils.SendMail(c.mailSender.Email, c.mailSender.Password, createUser.Email, message)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	createUser.VerificationCode = code

	// Check if user exists but not verified if so update the user's info
	if EmailErr == nil && !user.Verified {
		createUser.ID = user.ID
		err = c.db.UpdateUser(createUser)
		if err != nil {
			log.Error().Err(err).Send()
			return
		}
		w.Write([]byte("User exists but not verified another verification code has been sent to your email \n"))

	}

	// check if user doesn't exist if so create new user
	if EmailErr != nil {
		u := c.db.SignUp(createUser)
		res, _ := json.Marshal(u)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		w.Write([]byte("Verification code has been sent to your email \n"))
	}

}

func (c *Controller) SignIn(w http.ResponseWriter, r *http.Request) {
	//Ignore this needs a lot of modification

	var input models.SignInInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Error().Err(err).Send()

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userDetails, err := c.db.GetUserByEmail(input.Email)

	if err != nil {
		log.Error().Err(err).Send()

		w.WriteHeader(http.StatusBadRequest)
		println("No user found")
		return
	}

	res, _ := json.Marshal(userDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
