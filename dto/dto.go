package dto

import "time"

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

type ProductItemDTO struct {
	Id                     int64
	StatusChangeDate       time.Time
	PriceExcludingVAT      float64
	VatRate                float64
	DiscountFromBasicPrice float64
	MarginalTradeMarkup    float64
	Currency               string
	PriceWithVAT           float64
	Product                ProductDTO
	ProductItemStatus      string
	Region                 RegionDTO
}

type RoleDTO struct {
	Code   string
	NameKz string
	NameRu string
	NameEn string
}

type OrganizationDTO struct {
	id         int64
	externalId int64
	bin        string
	nameRu     string
	nameKz     string
	nameEn     string
	isResident bool
	blocked    bool
	corpPhone  string
	contactFio string
	website    string
	gln        string
	isStm      bool
	isNds      bool
	iban       string
	bic        string
	kbeCode    string
	address    string
	FormOfLaw  FormOfLawDTO
	KBE        KBEDTO
}

type FormOfLawDTO struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
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
