package utils

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func SendMail(sender string, password string, receiver string, message string) error {
	auth := smtp.PlainAuth("", sender, password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, sender, []string{receiver}, []byte(message))
	// fmt.Printf("%+v\n", err)

	return err
}

func SignUpMailBody(code int, timeout int) string {
	subject := "Welcome to Blogging System \n\n"
	body := fmt.Sprintf("We are so glad to have you here.\n\nYour code is %s\nThe code will expire in %d seconds.\nPlease don't share it with anyone.", strconv.Itoa(code), timeout)
	message := subject + body

	return message
}

func GenerateRandomCode() int {
	min := 1000
	max := 9999
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
