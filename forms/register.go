package forms

type RegistrationForm struct {
	ResidentType          string `form:"resident_type"`
	FormOfLawID           int    `form:"form_of_law_code"`
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
