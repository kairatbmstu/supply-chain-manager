package main

import (
	"encoding/gob"

	"example.com/m/v2/controller"
	"example.com/m/v2/dto"
	"example.com/m/v2/forms"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	// connStr := "postgresql://supplier_portal:123456@localhost/todos?sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})

	gob.Register(dto.UserDTO{})
	gob.Register(forms.LoginForm{})
	gob.Register(forms.RegistrationForm{})

	router.Use(sessions.Sessions("mysession", store))
	router.Static("web/css", "web/css")
	router.Static("web/js", "web/js")
	router.LoadHTMLGlob("web/templates/**/*")

	router.GET("/", controller.IndexControllerInstance.GetIndexPage)
	router.GET("/login", controller.LoginControllerInstance.GetLoginPage)
	router.POST("/login", controller.LoginControllerInstance.PostLogin)
	router.GET("/register/resident_type", controller.RegistrationControllerInstance.GetResidentType)
	router.POST("/register/resident_type", controller.RegistrationControllerInstance.PostResidentType)
	router.GET("/register/org_main_info", controller.RegistrationControllerInstance.GetMainInfo)
	router.POST("/register/org_main_info", controller.RegistrationControllerInstance.PostOrgMainInfo)

	router.GET("/register/org_additional_info", controller.RegistrationControllerInstance.GetRegisterOrganizationAdditionalInfo)
	router.POST("/register/org_additional_info", controller.RegistrationControllerInstance.PostRegisterOrganizationAdditionalInfo)

	router.GET("/register/contact_person_info", controller.RegistrationControllerInstance.GetContantPersonInfo)
	router.POST("/register/contact_person_info", controller.RegistrationControllerInstance.PostContactPerson)

	router.GET("/register/complete_registration", controller.RegistrationControllerInstance.GetCompleteRegistration)
	router.POST("/register/complete_registration", controller.RegistrationControllerInstance.PostCompleteRegistration)

	router.Run(":8080")
}

type ApplicationContext struct {
}
