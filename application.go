package main

import "example.com/m/v2/controller"

type WebApplicationContext struct {
	AssortimentController             controller.AssortimentController
	CommercialConditionsController    controller.CommercialConditionsController
	CommercialProposalController      controller.CommercialProposalController
	EnterProductApplicationController controller.EnterProductApplicationController
	ProductTemplateController         controller.ProductTemplateController
	ProductController                 controller.ProductController
	RegistrationController            controller.RegistrationController
	LoginController                   controller.LoginController
	IndexController                   controller.IndexController
}

func NewWebApplicationContext() *WebApplicationContext {
	var appContext = WebApplicationContext{}
	return &appContext
}
