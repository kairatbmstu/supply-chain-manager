package domain

type Currency string

const (
	KZT Currency = "KZT"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type CommercialProposalStatus string

const (
	CommercialProposalStatus_Draft         string = "draft"
	CommercialProposalStatus_ProcessingCM  string = "processing_cm"
	CommercialProposalStatus_ApprovedByCM  string = "approved_by_cm"
	CommercialProposalStatus_ProcessingRM  string = "processing_rm"
	CommercialProposalStatus_ApprovedByRM  string = "approved_by_rm"
	CommercialProposalStatus_FullyApproved string = "fully_approved"
)

const (
	EntryProposalStatus_Draft          string = "draft"
	EntryProposalStatus_ProcessingBUCP string = "processing_bucp"
	EntryProposalStatus_ApprovedByBUCP string = "approved_by_bucp"
	EntryProposalStatus_ProcessingOUKD string = "processing_oukd"
	EntryProposalStatus_ApprovedByOUKD string = "approved_by_oukd"
	EntryProposalStatus_FullyApproved  string = "fully_approved"
)

type User struct {
	Username string
	Email    string
	Password string
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

type FormOfLaw struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type KBE struct {
	ID     int
	Code   string
	NameKz string
	NameRu string
	NameEn string
}

type Measurement struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type Department struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type DeliveryConditions struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type PackageType struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type BarcodeMask struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type Country struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type StorageAndTransportationConditions struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type Color struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type Group struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type Subgroup struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type ChildWeight struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type TNVD struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
}

type Region struct {
	ID         int
	ExternalId int
	NameKz     string
	NameRu     string
	NameEn     string
}

type Category struct {
	ID            int
	NameKz        string
	NameRu        string
	NameEn        string
	SubCategories []SubCategory
}

type SubCategory struct {
	ID     int
	NameKz string
	NameRu string
	NameEn string
	Category
	ProductField []ProductField
}

type PaymentType struct {
	ID         int
	ExternalId int
	NameKz     string
	NameRu     string
	NameEn     string
}
