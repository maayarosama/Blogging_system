package controllers

import (
	"github.com/maayarosama/Blogging_system/config"
	"github.com/maayarosama/Blogging_system/models"
)

type Controller struct {
	db         models.DB
	mailSender config.MailSender
}

func NewController(db models.DB, mailSender config.MailSender) *Controller {
	return &Controller{db, mailSender}
}
