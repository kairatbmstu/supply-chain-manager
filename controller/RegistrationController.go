package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/forms"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var RegistrationControllerInstance RegistrationController

type RegistrationController struct {
}

func (r RegistrationController) GetRegistrationType(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("registration") == nil {
		session.Set("registration", new(forms.RegistrationForm))
		session.Save()
	}

	c.HTML(http.StatusOK, "register_resident_type.html", gin.H{})
}

func (r RegistrationController) PostRegistrationType(c *gin.Context) {
	session := sessions.Default(c)
	var residentTypeForm forms.ResidentTypeForm
	if c.ShouldBind(&residentTypeForm) == nil {
		log.Println(residentTypeForm.ResidentType)

		if session.Get("registration") != nil {
			registrationForm := session.Get("registration").(forms.RegistrationForm)
			registrationForm.ResidentType = residentTypeForm.ResidentType
			session.Save()
		}
	}

	c.Redirect(http.StatusFound, "/register/resident_details")
}

func (r RegistrationController) GetResidentDetails(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("registration") == nil {
		session.Set("registration", new(forms.RegistrationForm))
		session.Save()
	}

	if session.Get("registration") != nil {
		registrationForm := session.Get("registration").(*forms.RegistrationForm)
		log.Println("registrationForm : ", registrationForm)
		log.Println("orgName : ", session.Get("orgName"))
		log.Println("orgAddress : ", session.Get("orgAddress"))
		c.HTML(http.StatusOK, "register_resident_details.html", gin.H{
			"orgName":    registrationForm.OrgName,
			"orgAddress": registrationForm.OrgAddress,
		})
	} else {
		c.HTML(http.StatusOK, "register_resident_details.html", gin.H{})
	}

}

func (r RegistrationController) PostResidentDetails(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("registration") == nil {
		session.Set("registration", new(forms.RegistrationForm))
		session.Save()
	}

	var residentDetailsForm forms.ResidentDetailsForm
	if c.ShouldBind(&residentDetailsForm) == nil {
		log.Println("residentDetailsForm 1 : ", residentDetailsForm)

		if session.Get("registration") != nil {
			registrationForm := session.Get("registration").(*forms.RegistrationForm)
			registrationForm.OrgName = residentDetailsForm.OrgName
			registrationForm.OrgAddress = residentDetailsForm.OrgAddress
			registrationForm.FormOfLawCode = residentDetailsForm.FormOfLawCode
			session.Set("registration", registrationForm)
			session.Set("orgName", registrationForm.OrgName)
			session.Set("orgAddress", registrationForm.OrgAddress)
			log.Println("orgName :: ", session.Get("orgName"))
			log.Println("orgAddress :: ", session.Get("orgAddress"))
			session.Save()
			log.Println("registrationForm 2 : ", registrationForm)
		}
	}

	c.Redirect(http.StatusFound, "/register/resident_details")
}
