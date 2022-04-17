package controller

import (
	"strings"

	"example.com/m/v2/flash"
	"example.com/m/v2/forms"
	"github.com/gin-gonic/gin"
)

type MainInfoValidator struct {
	Errors []string
}

func (a *MainInfoValidator) validate(regForm *forms.OrgMainInfoForm) {
	if regForm.OrgName == "" {
		a.Errors = append(a.Errors, "OrgName is required")
	}

	if regForm.OrgAddress == "" {
		a.Errors = append(a.Errors, "OrgAddress is required")
	}

	if regForm.FormOfLawID == 0 {
		a.Errors = append(a.Errors, "FormOfLawID is required")
	}
}

func (a *MainInfoValidator) hasErrors() bool {
	return len(a.Errors) > 0
}

type AdditionalInfoValidator struct {
	Errors []string
}

func (a *AdditionalInfoValidator) validate(regForm *forms.OrgAdditionalInfo) {
	if regForm.OrgPhone == "" {
		a.Errors = append(a.Errors, "OrgPhone is required")
	}

	if regForm.OrgSite == "" {
		a.Errors = append(a.Errors, "OrgSite is required")
	}

	if regForm.IINBIN == "" {
		a.Errors = append(a.Errors, "IINBIN is required")
	}

	if regForm.AccountNumber == "" {
		a.Errors = append(a.Errors, "AccountNumber is required")
	}

	if regForm.BIC == "" {
		a.Errors = append(a.Errors, "BIC is required")
	}

	if regForm.KBE == "" {
		a.Errors = append(a.Errors, "KBE is required")
	}
}

func (a *AdditionalInfoValidator) hasErrors() bool {
	return len(a.Errors) > 0
}

type ContactPersonValidator struct {
	Errors []string
}

func (a *ContactPersonValidator) validate(regForm *forms.ContactPersonForm) {
	if regForm.Email == "" {
		a.Errors = append(a.Errors, "Email is required")
	}

	if regForm.ContactPersonFullname == "" {
		a.Errors = append(a.Errors, "ContactPersonFullname is required")
	}

	if regForm.ContactPersonPosition == "" {
		a.Errors = append(a.Errors, "ContactPersonPosition is required")
	}

}

func (a *ContactPersonValidator) hasErrors() bool {
	return len(a.Errors) > 0
}

type RegistrationFormValidator struct {
	Errors []string
}

func (a *RegistrationFormValidator) validate(c *gin.Context, regForm *forms.RegistrationForm) {

	if regForm.ResidentType == "" {
		flash.SetFlash(c, "ErrorResidentType", "ResidentType is required")
		a.Errors = append(a.Errors, "ResidentType is required")
	}

	if regForm.OrgName == "" {
		flash.SetFlash(c, "ErrorOrgName", "OrgName is required")
		a.Errors = append(a.Errors, "OrgName is required")
	}

	if regForm.OrgAddress == "" {
		flash.SetFlash(c, "ErrorOrgAddress", "OrgAddress is required")
		a.Errors = append(a.Errors, "OrgAddress is required")
	}

	if regForm.FormOfLawID == 0 {
		flash.SetFlash(c, "ErrorFormOfLawID", "FormOfLawID is required")
		a.Errors = append(a.Errors, "FormOfLawID is required")
	}
	if regForm.OrgPhone == "" {
		flash.SetFlash(c, "ErrorOrgPhone", "OrgPhone is required")
		a.Errors = append(a.Errors, "OrgPhone is required")
	}

	if regForm.OrgSite == "" {
		flash.SetFlash(c, "ErrorOrgSite", "OrgSite is required")
		a.Errors = append(a.Errors, "OrgSite is required")
	}

	if regForm.IINBIN == "" {
		flash.SetFlash(c, "ErrorIINBIN", "IINBIN is required")
		a.Errors = append(a.Errors, "IINBIN is required")
	}

	if regForm.AccountNumber == "" {
		flash.SetFlash(c, "ErrorAccountNumber", "AccountNumber is required")
		a.Errors = append(a.Errors, "AccountNumber is required")
	}

	if regForm.BIC == "" {
		flash.SetFlash(c, "ErrorBIC", "BIC is required")
		a.Errors = append(a.Errors, "BIC is required")
	}

	if regForm.KBE == "" {
		flash.SetFlash(c, "ErrorKBE", "KBE is required")
		a.Errors = append(a.Errors, "KBE is required")
	}

	if regForm.Email == "" {
		flash.SetFlash(c, "ErrorEmail", "Email is required")
		a.Errors = append(a.Errors, "Email is required")
	}

	if regForm.ContactPersonFullname == "" {
		flash.SetFlash(c, "ErrorContactPersonFullname", "ContactPersonFullname is required")
		a.Errors = append(a.Errors, "ContactPersonFullname is required")
	}

	if regForm.ContactPersonPosition == "" {
		flash.SetFlash(c, "ErrorContactPersonPosition", "ContactPersonPosition is required")
		a.Errors = append(a.Errors, "ContactPersonPosition is required")
	}

	if regForm.Password == "" {
		flash.SetFlash(c, "ErrorPassword", "Password is required")
		a.Errors = append(a.Errors, "Password is required")
	}

	if regForm.ConfirmationPassword == "" {
		flash.SetFlash(c, "ErrorConfirmationPassword", "ConfirmationPassword is required")
		a.Errors = append(a.Errors, "ConfirmationPassword is required")
	}

	if !strings.EqualFold(regForm.Password, regForm.ConfirmationPassword) {
		flash.SetFlash(c, "ErrorPasswordConfirmationPasswordShoudBeEqual", "Password and ConfirmationPassword should be equal")
		a.Errors = append(a.Errors, "Password and ConfirmationPassword should be equal")
	}

	if len(regForm.Password) >= 8 {
		flash.SetFlash(c, "ErrorPasswordFormat", "Password must contain more than 8 symbols")
		a.Errors = append(a.Errors, "Password must contain more than 8 symbols")
	}

}

func (a *RegistrationFormValidator) hasErrors() bool {
	return len(a.Errors) > 0
}

type EnterPasswordValidator struct {
	Errors []string
}

func (a *EnterPasswordValidator) validate(form *forms.EnterPasswordForm) {
	if form.Password == "" {
		a.Errors = append(a.Errors, "Password is required")
	}

	if form.ConfirmationPassword == "" {
		a.Errors = append(a.Errors, "ConfirmationPassword is required")
	}

	if !strings.EqualFold(form.Password, form.ConfirmationPassword) {
		a.Errors = append(a.Errors, "Password and ConfirmationPassword should be equal")
	}

	if len(form.Password) < 8 {
		a.Errors = append(a.Errors, "Password must contain more than 8 symbols")
	}

}

func (a *EnterPasswordValidator) hasErrors() bool {
	return len(a.Errors) > 0
}
