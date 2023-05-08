package controllers

import "github.com/maayarosama/Blogging_system/models"

type Controller struct {
	db models.DB
}

func NewController(db models.DB) *Controller {
	return &Controller{db}
}
