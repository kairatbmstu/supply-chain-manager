package forms

type RegistrationModel struct {
	ResidentType          string
	FormOfLawCode         string
	OrgName               string
	OrgPhone              string
	IIN_BIN               string
	BIC                   string
	AccountNumber         string
	KBE                   string
	GLN                   string
	OrgSite               string
	RegionCode            string
	OrgAddress            string
	STM                   string
	VAT                   string
	Email                 string
	ContactPersonFullname string
	ContactPersonPosition string
	Password              string
	ConfirmationPassword  string
}

type ResidentTypeForm struct {
	ResidentType string `form:"resident_type"`
}

type ResidentDetailsForm struct {
	OrgAddress    string `form:"org_address"`
	OrgName       string `form:"org_name"`
	FormOfLawCode string `form:"form_of_law_code"`
}
