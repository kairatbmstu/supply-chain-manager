package forms

type RegistrationModel struct {
	IsResident            bool
	FormOfLawID           int
	OrgName               string
	OrgPhone              string
	IINBIN                string
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
	OrgAddress  string `form:"org_address"`
	OrgName     string `form:"org_name"`
	FormOfLawID int    `form:"form_of_law_code"`
}

type ResidentOrgAddInfoForm struct {
	OrgPhone      string `form:"org_phone"`
	OrgSite       string `form:"org_site"`
	IINBIN        string `form:"iin_bin"`
	AccountNumber string `form:"account_number"`
	BIC           string `form:"bic"`
	KBE           string `form:"kbe"`
	RegionID      int    `form:"region_id"`
	GLN           string `form:"gln"`
}
