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
	initRegistrationForm(c)
	c.HTML(http.StatusOK, "register_resident_type.html", gin.H{})
}

func (r RegistrationController) PostRegistrationType(c *gin.Context) {
	session := sessions.Default(c)
	var residentTypeForm forms.ResidentTypeForm
	if c.ShouldBind(&residentTypeForm) == nil {
		log.Println(residentTypeForm.ResidentType)

		if session.Get("registration") != nil {
			registrationForm := session.Get("registration").(forms.RegistrationModel)
			registrationForm.ResidentType = residentTypeForm.ResidentType
			err := session.Save()
			if err != nil {
				log.Println("error : ", err)
				c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
				return
			}
		}
	}

	c.Redirect(http.StatusFound, "/register/resident_details")
}

func (r RegistrationController) GetResidentDetails(c *gin.Context) {
	initRegistrationForm(c)

	if getRegistrationForm(c) != nil {
		registrationForm := getRegistrationForm(c)
		c.HTML(http.StatusOK, "register_resident_details.html", gin.H{
			"registrationForm": registrationForm,
		})
	} else {
		c.HTML(http.StatusOK, "register_resident_details.html", gin.H{})
	}

}

func (r RegistrationController) PostResidentDetails(c *gin.Context) {
	initRegistrationForm(c)

	var residentDetailsForm forms.ResidentDetailsForm
	if c.ShouldBind(&residentDetailsForm) == nil {

		if getRegistrationForm(c) != nil {
			registrationForm := getRegistrationForm(c)
			registrationForm.OrgName = residentDetailsForm.OrgName
			registrationForm.OrgAddress = residentDetailsForm.OrgAddress
			registrationForm.FormOfLawID = residentDetailsForm.FormOfLawID
			saveRegistrationForm(c, registrationForm)
		}
	}

	c.Redirect(http.StatusFound, "/register/organization_additional_info")
}

func (r RegistrationController) GetRegisterOrganizationAdditionalInfo(c *gin.Context) {
	initRegistrationForm(c)
	c.HTML(http.StatusOK, "register_organization_additional_info.html", gin.H{})
}

func (r RegistrationController) PostRegisterOrganizationAdditionalInfo(c *gin.Context) {
	initRegistrationForm(c)
	c.HTML(http.StatusOK, "register_organization_additional_info.html", gin.H{})
}

func (r RegistrationController) GetContantPersonInfo(c *gin.Context) {
	initRegistrationForm(c)
	c.HTML(http.StatusOK, "register_organization_additional_info.html", gin.H{})
}

func (r RegistrationController) PostContactPerson(c *gin.Context) {
	initRegistrationForm(c)
	c.HTML(http.StatusOK, "register_organization_additional_info.html", gin.H{})
}

func initRegistrationForm(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("registration") == nil {
		session.Set("registration", new(forms.RegistrationModel))
		err := session.Save()
		if err != nil {
			log.Println("error : ", err)
			c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
			return
		}
	}
}

func getRegistrationForm(c *gin.Context) *forms.RegistrationModel {
	session := sessions.Default(c)
	if session.Get("registration") != nil {
		regForm := session.Get("registration").(forms.RegistrationModel)
		return &regForm
	}
	return nil
}

func saveRegistrationForm(c *gin.Context, registrationForm *forms.RegistrationModel) {
	session := sessions.Default(c)
	session.Set("registration", registrationForm)
	err := session.Save()
	if err != nil {
		log.Println("error : ", err)
		c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
		return
	}
}
