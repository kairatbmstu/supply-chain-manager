package forms

import "fmt"

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

func (r RegistrationForm) Println() {
	fmt.Println("ResidentType : 		 ", r.ResidentType)
	fmt.Println("FormOfLawID : 			 ", r.FormOfLawID)
	fmt.Println("OrgName : 				 ", r.OrgName)
	fmt.Println("OrgPhone : 			 ", r.OrgPhone)
	fmt.Println("IINBIN : 				 ", r.IINBIN)
	fmt.Println("BIC : 					 ", r.BIC)
	fmt.Println("AccountNumber : 		 ", r.AccountNumber)
	fmt.Println("KBE : 					 ", r.KBE)
	fmt.Println("GLN : 					 ", r.GLN)
	fmt.Println("OrgSite : 				 ", r.OrgSite)
	fmt.Println("RegionID : 			 ", r.RegionID)
	fmt.Println("OrgAddress : 			 ", r.OrgAddress)
	fmt.Println("STM : 					 ", r.STM)
	fmt.Println("VAT : 					 ", r.VAT)
	fmt.Println("Email : 				 ", r.ResidentType)
	fmt.Println("ContactPersonFullname : ", r.ContactPersonFullname)
	fmt.Println("ContactPersonPosition : ", r.ContactPersonPosition)
	fmt.Println("Password : 			 ", r.Password)
	fmt.Println("ConfirmationPassword :  ", r.ConfirmationPassword)
}
