package controllers

import (
	"github.com/maayarosama/Blogging_system/config"
	"github.com/maayarosama/Blogging_system/models"
)

type Controller struct {
	db         models.DB
	mailSender config.MailSender
	jwt        config.JwtToken
}

func NewController(db models.DB, mailSender config.MailSender, jwt config.JwtToken) *Controller {
	return &Controller{db, mailSender, jwt}
}
