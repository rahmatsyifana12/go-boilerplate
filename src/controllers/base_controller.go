package controllers

import "github.com/sarulabs/di"

type Controller struct {	
}

func NewController(ioc di.Container) *Controller {
	return &Controller{}
}