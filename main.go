package main

import (
	"encoding/gob"
	"log"

	"example.com/m/v2/database"
	"example.com/m/v2/dto"
	"example.com/m/v2/forms"
	"example.com/m/v2/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	database.DB, err = database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})

	gob.Register(dto.UserDTO{})
	gob.Register(forms.LoginForm{})
	gob.Register(forms.RegistrationForm{})

	router.Use(sessions.Sessions("mysession", store))
	router.Use(middleware.CheckAuthentification)
	router.Static("web/css", "web/css")
	router.Static("web/js", "web/js")
	router.LoadHTMLGlob("web/templates/**/*")

	var context = NewWebApplicationContext()

	router.GET("/", context.IndexController.GetIndexPage)
	router.GET("/login", context.LoginController.GetLoginPage)
	router.POST("/login", context.LoginController.PostLogin)
	router.GET("/logout", context.LoginController.Logout)
	router.GET("/register/resident_type", context.RegistrationController.GetResidentType)
	router.POST("/register/resident_type", context.RegistrationController.PostResidentType)
	router.GET("/register/org_main_info", context.RegistrationController.GetMainInfo)
	router.POST("/register/org_main_info", context.RegistrationController.PostOrgMainInfo)

	router.GET("/register/org_additional_info", context.RegistrationController.GetOrgAdditionalInfo)
	router.POST("/register/org_additional_info", context.RegistrationController.PostOrgAdditionalInfo)

	router.GET("/register/contact_person_info", context.RegistrationController.GetContantPersonInfo)
	router.POST("/register/contact_person_info", context.RegistrationController.PostContactPerson)

	router.GET("/register/enter_password", context.RegistrationController.GetEnterPassword)
	router.POST("/register/enter_password", context.RegistrationController.PostEnterPassword)

	router.GET("/register/complete_registration", context.RegistrationController.GetCompleteRegistration)
	router.POST("/register/complete_registration", context.RegistrationController.PostCompleteRegistration)

	router.GET("/register/waiting", context.RegistrationController.GetWaiting)
	router.POST("/register/waiting", context.RegistrationController.PostWaiting)

	router.GET("/product-templates/get-details", context.ProductTemplateController.GetDetails)
	router.GET("/product-templates/get-create-form", context.ProductTemplateController.GetCreateForm)
	router.POST("/product-templates/create", context.ProductTemplateController.Create)
	router.GET("/product-templates/get-update-form", context.ProductTemplateController.GetUpdateForm)
	router.POST("/product-templates/update", context.ProductTemplateController.Update)
	router.GET("/product-templates/get-product-templates-list", context.ProductTemplateController.GetProductTemplatesList)

	router.Run(":8080")
}
