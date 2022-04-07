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

func (r RegistrationController) GetResidentType(c *gin.Context) {
	initRegistrationForm(c)
	c.HTML(http.StatusOK, "resident_type.html", gin.H{})
}

func (r RegistrationController) PostResidentType(c *gin.Context) {
	session := sessions.Default(c)
	var regForm forms.RegistrationForm
	if c.ShouldBind(&regForm) == nil {
		log.Println("regForm : ", regForm)

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

	if getRegistrationForm(c) != nil {
		registrationForm := getRegistrationForm(c)
		c.HTML(http.StatusOK, "org_main_info.html", gin.H{
			"regForm": registrationForm,
		})
	} else {
		c.HTML(http.StatusOK, "org_main_info.html", gin.H{})
	}

}

func (r RegistrationController) PostOrgMainInfo(c *gin.Context) {
	initRegistrationForm(c)

	var regForm = getRegistrationForm(c)
	if c.ShouldBind(&regForm) == nil {
		log.Println("regForm : ", regForm)
		saveRegistrationForm(c, regForm)
	}

	c.Redirect(http.StatusFound, "/register/org_additional_info")
}

func (r RegistrationController) GetRegisterOrganizationAdditionalInfo(c *gin.Context) {
	initRegistrationForm(c)

	if getRegistrationForm(c) != nil {
		regForm := getRegistrationForm(c)
		c.HTML(http.StatusOK, "org_additional_info.html", gin.H{
			"regForm": regForm,
		})
	} else {
		c.HTML(http.StatusOK, "org_additional_info.html", gin.H{})
	}

}

func (r RegistrationController) PostRegisterOrganizationAdditionalInfo(c *gin.Context) {
	initRegistrationForm(c)

	var regForm = getRegistrationForm(c)
	if c.ShouldBind(&regForm) == nil {
		log.Println("regForm : ", regForm)
		saveRegistrationForm(c, regForm)
	}

	c.Redirect(http.StatusFound, "/register/contact_person_info")
}

func (r RegistrationController) GetContantPersonInfo(c *gin.Context) {
	initRegistrationForm(c)

	if getRegistrationForm(c) != nil {
		registrationForm := getRegistrationForm(c)
		c.HTML(http.StatusOK, "contact_person_info.html", gin.H{
			"regForm": registrationForm,
		})
	} else {
		c.HTML(http.StatusOK, "contact_person_info.html", gin.H{})
	}

}

func (r RegistrationController) GetCompleteRegistration(c *gin.Context) {
	initRegistrationForm(c)
	if getRegistrationForm(c) != nil {
		registrationForm := getRegistrationForm(c)
		c.HTML(http.StatusOK, "complete_registration.html", gin.H{
			"regForm": registrationForm,
		})
	} else {
		c.HTML(http.StatusOK, "complete_registration.html", gin.H{})
	}
}

func (r RegistrationController) PostCompleteRegistration(c *gin.Context) {
	initRegistrationForm(c)
	var regForm = getRegistrationForm(c)
	if c.ShouldBind(&regForm) == nil {
		log.Println("regForm : ", regForm)
		saveRegistrationForm(c, regForm)
	}
	c.HTML(http.StatusOK, "complete_registration.html", gin.H{})
}

func (r RegistrationController) PostContactPerson(c *gin.Context) {
	initRegistrationForm(c)

	var regForm = getRegistrationForm(c)
	if c.ShouldBind(&regForm) == nil {
		log.Println("regForm : ", regForm)
		saveRegistrationForm(c, regForm)
	}

	c.Redirect(http.StatusFound, "/register/complete_registration")
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

func saveRegistrationForm(c *gin.Context, registrationForm *forms.RegistrationForm) {
	session := sessions.Default(c)
	session.Set("registration", registrationForm)
	err := session.Save()
	if err != nil {
		log.Println("error : ", err)
		c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
		return
	}
}
