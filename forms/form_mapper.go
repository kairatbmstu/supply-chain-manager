package forms

import (
	"strings"

	"example.com/m/v2/dto"
)

var RegistrationFormMapperInstance RegistrationFormMapper

type RegistrationFormMapper struct {
}

func (r RegistrationFormMapper) Map(registrationForm *RegistrationForm) dto.RegistrationDTO {
	var result = dto.RegistrationDTO{
		FormOfLawID:           registrationForm.FormOfLawID,
		OrganizationPhone:     registrationForm.OrgPhone,
		IINBIN:                registrationForm.IINBIN,
		BIC:                   registrationForm.BIC,
		GLN:                   registrationForm.GLN,
		Email:                 registrationForm.Email,
		ContactPersonFullname: registrationForm.ContactPersonFullname,
		ContactPersonPosition: registrationForm.ContactPersonPosition,
		Password:              registrationForm.Password,
		ConfirmationPassword:  registrationForm.ConfirmationPassword,
		OrgAddress:            registrationForm.OrgAddress,
		RegionID:              registrationForm.RegionID,
		OrganizationSite:      registrationForm.OrgSite,
		AccountNumber:         registrationForm.AccountNumber,
		KBE:                   registrationForm.KBE,
	}

	switch registrationForm.ResidentType {
	case "resident":
		result.IsResident = true
	case "non_resident":
		result.IsResident = false
	}

	if strings.EqualFold(registrationForm.STM, "checked") {
		result.IsSTM = true
	}

	if strings.EqualFold(registrationForm.VAT, "checked") {
		result.IsVAT = true
	}

	return result
}
