package controller

import (
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/forms"
	"example.com/m/v2/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var RegistrationControllerInstance RegistrationController

type RegistrationController struct {
}

func (r RegistrationController) GetResidentType(c *gin.Context) {
	initRegistrationForm(c)
	c.HTML(http.StatusOK, "resident_type.html", gin.H{})
}

func (r RegistrationController) PostResidentType(c *gin.Context) {
	session := sessions.Default(c)
	var regForm forms.RegistrationForm
	if c.ShouldBind(&regForm) == nil {
		regForm.Println()

		if session.Get("registration") != nil {
			registrationForm := session.Get("registration").(forms.RegistrationForm)
			registrationForm.ResidentType = regForm.ResidentType
			err := session.Save()
			if err != nil {
				log.Println("error : ", err)
				c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
				return
			}
		}
	}

	c.Redirect(http.StatusFound, "/register/org_main_info")
}

func (r RegistrationController) GetMainInfo(c *gin.Context) {
	initRegistrationForm(c)
	session := sessions.Default(c)
	session.AddFlash("Пожалуйста заполните поле Наименование организации")
	session.Save()

	registrationForm := getRegistrationForm(c)
	c.HTML(http.StatusOK, "org_main_info.html", gin.H{
		"regForm": registrationForm,
	})

}

func (r RegistrationController) PostOrgMainInfo(c *gin.Context) {
	initRegistrationForm(c)
	session := sessions.Default(c)

	var regForm = getRegistrationForm(c)

	var err = c.ShouldBind(&regForm)

	if err != nil {
		log.Println("error : ", err)
		c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
		return
	}

	if regForm.OrgName == "" {
		session.AddFlash("Пожалуйста заполните поле Наименование организации")
		session.Save()
		c.Redirect(303, "/register/org_main_info")
		return
	}

	fmt.Println(" regForm : ")
	regForm.Println()
	session.Set("registration", regForm)
	err = session.Save()
	if err != nil {
		log.Println("error : ", err)
		c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
		return
	}

	c.Redirect(http.StatusFound, "/register/org_additional_info")
}

func (r RegistrationController) GetRegisterOrganizationAdditionalInfo(c *gin.Context) {
	initRegistrationForm(c)

	regForm := getRegistrationForm(c)
	c.HTML(http.StatusOK, "org_additional_info.html", gin.H{
		"regForm": regForm,
	})

}

func (r RegistrationController) PostRegisterOrganizationAdditionalInfo(c *gin.Context) {
	initRegistrationForm(c)

	var regForm = getRegistrationForm(c)
	if c.ShouldBind(&regForm) == nil {
		regForm.Println()
		err := saveSession(c, regForm)
		if err != nil {
			log.Println("error : ", err)
			c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
			return
		}

	}

	c.Redirect(http.StatusFound, "/register/contact_person_info")
}

func (r RegistrationController) GetContantPersonInfo(c *gin.Context) {
	initRegistrationForm(c)

	registrationForm := getRegistrationForm(c)
	c.HTML(http.StatusOK, "contact_person_info.html", gin.H{
		"regForm": registrationForm,
	})

}

func (r RegistrationController) PostContactPerson(c *gin.Context) {
	initRegistrationForm(c)

	var regForm = getRegistrationForm(c)
	if c.ShouldBind(&regForm) == nil {
		regForm.Println()
		session := sessions.Default(c)
		session.Set("registration", regForm)
		err := session.Save()
		if err != nil {
			log.Println("error : ", err)
			c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
			return
		}
	}

	c.Redirect(http.StatusFound, "/register/complete_registration")
}

func (r RegistrationController) GetWaiting(c *gin.Context) {
	c.HTML(http.StatusOK, "waiting.html", gin.H{})
}

func (r RegistrationController) PostWaiting(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}

func (r RegistrationController) GetCompleteRegistration(c *gin.Context) {
	initRegistrationForm(c)
	registrationForm := getRegistrationForm(c)
	c.HTML(http.StatusOK, "complete_registration.html", gin.H{
		"regForm": registrationForm,
	})
}

func (r RegistrationController) PostCompleteRegistration(c *gin.Context) {
	initRegistrationForm(c)
	var regForm = getRegistrationForm(c)
	if c.ShouldBind(&regForm) == nil {
		regForm.Println()
		err := saveSession(c, regForm)
		if err != nil {
			log.Println("error : ", err)
			c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
			return
		}

		registrationCommand := forms.RegistrationFormMapperInstance.Map(regForm)
		userDto, err := service.RegistrationServiceInstance.RegisterUser(registrationCommand)

		if err != nil {
			log.Println("error : ", err)
			c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
			return
		}

		if userDto != nil {
			c.Redirect(http.StatusFound, "/register/waiting")
			return
		}

	}
	c.HTML(http.StatusOK, "complete_registration.html", gin.H{})
}

func initRegistrationForm(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("registration") == nil {
		session.Set("registration", new(forms.RegistrationForm))
		err := session.Save()
		if err != nil {
			log.Println("error : ", err)
			c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
			return
		}
	}
}

func getRegistrationForm(c *gin.Context) *forms.RegistrationForm {
	session := sessions.Default(c)
	if session.Get("registration") != nil {
		regForm := session.Get("registration").(forms.RegistrationForm)
		return &regForm
	}
	return nil
}

func saveSession(c *gin.Context, registrationForm *forms.RegistrationForm) error {
	session := sessions.Default(c)
	session.Set("registration", registrationForm)
	err := session.Save()
	return err
}
