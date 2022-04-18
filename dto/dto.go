package dto

type RegistrationCommand struct {
	IsResident            bool
	FormOfLawID           int
	OrganizationPhone     string
	IINBIN                string
	BIC                   string
	AccountNumber         string
	KBE                   string
	GLN                   string
	OrganizationSite      string
	RegionID              int
	OrgAddress            string
	IsSTM                 bool
	IsVAT                 bool
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

type RoleDTO struct {
	Code   string
	NameKz string
	NameRu string
	NameEn string
}

type OrganizationDTO struct {
	ID         int64
	ExternalId int64
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
	IBAN       string
	Bic        string
	KbeCode    string
	Address    string
	FormOfLaw  FormOfLawDTO
	KBE        KBEDTO
}

type FormOfLawDTO struct {
	ID         int
	Code       string
	IsResident bool
	NameKz     string
	NameRu     string
	NameEn     string
}

type KBEDTO struct {
	ID     int
	Code   string
	NameKz string
	NameRu string
	NameEn string
}

type MeasurementDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type DepartmentDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type DeliveryConditionsDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type PackageTypeDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type BarcodeMaskDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type CountryDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type StorageAndTransportationConditionsDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type ColorDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type GroupDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type SubgroupDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type ChildWeightDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type TnvdDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type RegionDTO struct {
	ID         int
	ExternalId int
	NameKz     string
	NameRu     string
	NameEn     string
}

type CategoryDTO struct {
	ID            int
	NameKz        string
	NameRu        string
	NameEn        string
	SubCategories []SubCategoryDTO
}

type SubCategoryDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
	CategoryDTO
	ProductField []ProductFieldDTO
}

type PaymentTypeDTO struct {
	ID         int
	ExternalId int
	NameKz     string
	NameRu     string
	NameEn     string
}

type ProductFieldDTO struct {
	ID          int
	FieldCode   string
	FieldNameKz string
	FieldNameRu string
	FieldNameEn string
	SubCategory SubCategoryDTO
}
