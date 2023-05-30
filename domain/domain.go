package domain

import "time"

type User struct {
	ID                    int
	IIN                   string
	FirstName             string
	LastName              string
	MiddleName            string
	Phone                 string
	IsActive              bool
	ActiveDirectoryLink   string
	RegDatetime           time.Time
	OtdelID               int
	OrganizationID        int
	Username              string
	Email                 string
	Password              string
	AllowSelfRegistration bool
}

type Role struct {
	Code   string
	NameKz string
	NameRu string
	NameEn string
}

type Organization struct {
	ID         int
	ExternalId int
	Bin        string
	NameRu     string
	NameKz     string
	NameEn     string
	IsResident bool
	Blocked    bool
	CorpPhone  string
	ContactFio string
	Website    string
	Gln        string
	IsStm      bool
	IsNds      bool
	Iban       string
	Bic        string
	Address    string
	FormOfLaw  FormOfLaw
	KBE        KBE
}
