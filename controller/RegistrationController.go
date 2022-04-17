package controller

import (
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

func (r RegistrationController) GetMainInfo(c *gin.Context) {
	initRegistrationForm(c)

	registrationForm := getRegistrationForm(c)
	c.HTML(http.StatusOK, "org_main_info.html", gin.H{
		"regForm": registrationForm,
	})

}

func (r RegistrationController) GetContantPersonInfo(c *gin.Context) {
	initRegistrationForm(c)

	registrationForm := getRegistrationForm(c)
	c.HTML(http.StatusOK, "contact_person_info.html", gin.H{
		"regForm": registrationForm,
	})

}

func (r RegistrationController) GetOrgAdditionalInfo(c *gin.Context) {
	initRegistrationForm(c)

	regForm := getRegistrationForm(c)
	c.HTML(http.StatusOK, "org_additional_info.html", gin.H{
		"regForm": regForm,
	})

}

func (r RegistrationController) GetEnterPassword(c *gin.Context) {
	initRegistrationForm(c)

	regForm := getRegistrationForm(c)
	c.HTML(http.StatusOK, "enter_password.html", gin.H{
		"regForm": regForm,
	})

}

func (r RegistrationController) GetCompleteRegistration(c *gin.Context) {
	initRegistrationForm(c)
	registrationForm := getRegistrationForm(c)
	c.HTML(http.StatusOK, "complete_registration.html", gin.H{
		"regForm": registrationForm,
	})
}

func (r RegistrationController) GetWaiting(c *gin.Context) {
	c.HTML(http.StatusOK, "waiting.html", gin.H{})
}

func (r RegistrationController) PostResidentType(c *gin.Context) {
	session := sessions.Default(c)
	var residentTypeForm forms.ResidentTypeForm
	if c.ShouldBind(&residentTypeForm) == nil {

		if session.Get("registration") != nil {
			registrationForm := session.Get("registration").(forms.RegistrationForm)
			registrationForm.ResidentType = residentTypeForm.ResidentType
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

func (r RegistrationController) PostOrgMainInfo(c *gin.Context) {
	initRegistrationForm(c)
	session := sessions.Default(c)

	regForm := getRegistrationForm(c)

	var mainInfo forms.OrgMainInfoForm
	var err = c.ShouldBind(&mainInfo)
	if err != nil {
		log.Println("error : ", err)
		c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
		return
	}

	validator := MainInfoValidator{}
	validator.validate(&mainInfo)

	if validator.hasErrors() {
		c.HTML(200, "error400.html", gin.H{
			"errors": validator.Errors,
		})
		return
	}

	regForm.FormOfLawID = mainInfo.FormOfLawID
	regForm.OrgAddress = mainInfo.OrgAddress
	regForm.OrgName = mainInfo.OrgName

	session.Set("registration", regForm)
	err = session.Save()
	if err != nil {
		log.Println("error : ", err)
		c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
		return
	}

	c.Redirect(http.StatusFound, "/register/org_additional_info")
}

func (r RegistrationController) PostOrgAdditionalInfo(c *gin.Context) {
	initRegistrationForm(c)

	var regForm = getRegistrationForm(c)
	var addinfo forms.OrgAdditionalInfo
	err := c.ShouldBind(&addinfo)
	if err == nil {
		if err != nil {
			log.Println("error : ", err)
			c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
			return
		}

		validator := AdditionalInfoValidator{}
		validator.validate(&addinfo)
		if validator.hasErrors() {
			c.HTML(http.StatusInternalServerError, "error400.html", gin.H{
				"errors": validator.Errors,
			})
		}

		regForm.AccountNumber = addinfo.AccountNumber
		regForm.BIC = addinfo.BIC
		regForm.IINBIN = addinfo.IINBIN
		regForm.KBE = addinfo.KBE
		regForm.OrgPhone = addinfo.OrgPhone
		regForm.OrgSite = addinfo.OrgSite
		regForm.STM = addinfo.STM
		regForm.VAT = addinfo.VAT
		err := saveSession(c, regForm)
		if err != nil {
			log.Println("error : ", err)
			c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
			return
		}
	}
	if err != nil {
		log.Println("error : ", err)
		c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
		return
	}
	c.Redirect(http.StatusFound, "/register/contact_person_info")
}

func (r RegistrationController) PostContactPerson(c *gin.Context) {
	initRegistrationForm(c)

	var regForm = getRegistrationForm(c)
	var contactPerson forms.ContactPersonForm
	err := c.ShouldBind(&contactPerson)
	if err != nil {
		log.Println("error : ", err)
		c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
		return
	}

	validator := ContactPersonValidator{}
	validator.validate(&contactPerson)
	if validator.hasErrors() {
		c.HTML(http.StatusBadRequest, "error400.html", gin.H{
			"errors": validator.Errors,
		})
		return
	}

	regForm.ContactPersonFullname = contactPerson.ContactPersonFullname
	regForm.ContactPersonPosition = contactPerson.ContactPersonPosition
	regForm.Email = contactPerson.Email

	session := sessions.Default(c)
	session.Set("registration", regForm)
	err = session.Save()
	if err != nil {
		log.Println("error : ", err)
		c.HTML(http.StatusInternalServerError, "error500.html", gin.H{})
		return
	}

	c.Redirect(http.StatusFound, "/register/enter_password")
}

func (r RegistrationController) PostCompleteRegistration(c *gin.Context) {
	initRegistrationForm(c)
	var regForm = getRegistrationForm(c)
	if regForm != nil {
		var regFormValidator RegistrationFormValidator
		regFormValidator.validate(c, regForm)

		if regFormValidator.hasErrors() {
			c.Redirect(http.StatusFound, "/register/complete_registration")
			return
		}

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

func (r RegistrationController) PostEnterPassword(c *gin.Context) {
	initRegistrationForm(c)
	var regForm = getRegistrationForm(c)
	var enterPasswordForm forms.EnterPasswordForm
	if c.ShouldBind(&enterPasswordForm) == nil {

		validator := EnterPasswordValidator{}
		validator.validate(&enterPasswordForm)

		if validator.hasErrors() {
			c.HTML(http.StatusBadRequest, "error400.html", gin.H{
				"errors": validator.Errors,
			})
			return
		}

		regForm.Password = enterPasswordForm.Password
		regForm.ConfirmationPassword = enterPasswordForm.ConfirmationPassword

		var regFormValidator RegistrationFormValidator
		regFormValidator.validate(c, regForm)

		if regFormValidator.hasErrors() {
			c.Redirect(http.StatusFound, "/register/complete_registration")
			return
		}

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

func (r RegistrationController) PostWaiting(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
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
