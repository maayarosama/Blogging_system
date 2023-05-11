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

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	// Validations for unique username and email need to be added

	createUser := &models.User{}
	utils.ParseBody(r, createUser)

	_, err := c.db.GetUserByEmail(createUser.Email)
	if err == nil {
		log.Error().Err(err).Send()

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User Already exists \n"))
		return
	}

	code := utils.GenerateRandomCode()
	message := utils.SignUpMailBody(code, c.mailSender.Timeout)

	err = utils.SendMail(c.mailSender.Email, c.mailSender.Password, createUser.Email, message)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	createUser.VerificationCode = code

	u := c.db.SignUp(createUser)
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	w.Write([]byte("Verification code has been sent to your email \n"))

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
