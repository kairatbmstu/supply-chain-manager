package dto

type RegistrationCommand struct {
	ResidentType          string
	FormOfLawCode         string
	OrganizationPhone     string
	IIN_BIN               string
	BIC                   string
	AccountNumber         string
	KBE                   string
	GLN                   string
	OrganizationSite      string
	RegionCode            string
	Address               string
	STM                   string
	VAT                   string
	Email                 string
	ContactPersonFullname string
	ContactPersonPosition string
	Password              string
	ConfirmationPassword  string
}

type UserDTO struct {
	Username string
	Email    string
	Password string
}
