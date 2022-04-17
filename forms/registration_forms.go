package forms

type RegistrationForm struct {
	ResidentType          string `form:"resident_type"`
	FormOfLawID           int    `form:"form_of_law"`
	OrgName               string `form:"org_name"`
	OrgPhone              string `form:"org_phone"`
	IINBIN                string `form:"iin_bin"`
	BIC                   string `form:"bic"`
	AccountNumber         string `form:"account_number"`
	KBE                   string `form:"kbe"`
	GLN                   string `form:"gln"`
	OrgSite               string `form:"org_site"`
	RegionID              int    `form:"region_id"`
	OrgAddress            string `form:"org_address"`
	STM                   string `form:"stm"`
	VAT                   string `form:"vat"`
	Email                 string `form:"email"`
	ContactPersonFullname string `form:"contact_fullname"`
	ContactPersonPosition string `form:"contact_position"`
	Password              string `form:"password"`
	ConfirmationPassword  string `form:"confirmation_password"`
}

type ResidentTypeForm struct {
	ResidentType string `form:"resident_type"`
}

type OrgMainInfoForm struct {
	OrgName     string `form:"org_name"`
	OrgAddress  string `form:"org_address"`
	FormOfLawID int    `form:"form_of_law"`
}

type OrgAdditionalInfo struct {
	IINBIN        string `form:"iin_bin"`
	OrgSite       string `form:"org_site"`
	OrgPhone      string `form:"org_phone"`
	AccountNumber string `form:"account_number"`
	BIC           string `form:"bic"`
	KBE           string `form:"kbe"`
	STM           string `form:"stm"`
	VAT           string `form:"vat"`
}

type ContactPersonForm struct {
	Email                 string `form:"email"`
	ContactPersonFullname string `form:"contact_fullname"`
	ContactPersonPosition string `form:"contact_position"`
}

type EnterPasswordForm struct {
	Password             string `form:"password"`
	ConfirmationPassword string `form:"confirmation_password"`
}
