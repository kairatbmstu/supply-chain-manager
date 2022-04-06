package main

import (
	"database/sql"
	"encoding/gob"
	"log"

	"example.com/m/v2/controller"
	"example.com/m/v2/dto"
	"example.com/m/v2/forms"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	connStr := "postgresql://supplier_portal:123456@localhost/todos?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})

	gob.Register(dto.UserDTO{})
	gob.Register(forms.LoginForm{})
	gob.Register(forms.RegistrationModel{})
	gob.Register(forms.ResidentDetailsForm{})
	gob.Register(forms.ResidentTypeForm{})

	router.Use(sessions.Sessions("mysession", store))
	router.Static("/css", "web/css")
	router.Static("/js", "web/js")
	router.LoadHTMLGlob("web/templates/**/*")

	router.GET("/", controller.IndexControllerInstance.GetIndexPage)
	router.GET("/login", controller.LoginControllerInstance.GetLoginPage)
	router.POST("/login", controller.LoginControllerInstance.PostLogin)
	router.GET("/register/resident_type", controller.RegistrationControllerInstance.GetRegistrationType)
	router.POST("/register/resident_type", controller.RegistrationControllerInstance.PostRegistrationType)
	router.GET("/register/resident_details", controller.RegistrationControllerInstance.GetResidentDetails)
	router.POST("/register/resident_details", controller.RegistrationControllerInstance.PostResidentDetails)

	router.GET("/register/organization_additional_info", controller.RegistrationControllerInstance.GetRegisterOrganizationAdditionalInfo)
	router.POST("/register/organization_additional_info", controller.RegistrationControllerInstance.PostRegisterOrganizationAdditionalInfo)

	router.GET("/register/contact_person_info", controller.RegistrationControllerInstance.GetContantPersonInfo)
	router.POST("/register/contact_person_info", controller.RegistrationControllerInstance.PostContactPerson)

	router.Run(":8080")
}

type ApplicationContext struct {
}
