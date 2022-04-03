package domain

import "time"

type Currency string

const (
	KZT Currency = "KZT"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type ProductItem struct {
	Id                     int64
	StatusChangeDate       time.Time
	PriceExcludingVAT      float64
	VatRate                float64
	DiscountFromBasicPrice float64
	MarginalTradeMarkup    float64
	Currency               Currency
	PriceWithVAT           float64
	Product                Product
	ProductItemStatus      ProductItemStatus
	Region                 Region
}

type ProductItemStatus string

const (
	ProductItemStatus_Draft             string = "draft"
	ProductItemStatus_ProcessingCM      string = "processing_cm"
	ProductItemStatus_ApprovedByCM      string = "approved_by_cm"
	ProductItemStatus_ProcessingRM      string = "processing_rm"
	ProductItemStatus_ApprovedByRM      string = "approved_by_rm"
	ProductItemStatus_FullyApprovedInCP string = "fully_approved_in_cp"

	ProductItemStatus_ReadyForEnterAssortiment string = "ready_for_enter_in_assortiment"
	ProductItemStatus_ProcessingBUCP           string = "processing_bucp"
	ProductItemStatus_ApprovedByBUCP           string = "approved_by_bucp"
	ProductItemStatus_ProcessingOUKD           string = "processing_oukd"
	ProductItemStatus_ApprovedByOUKD           string = "approved_by_oukd"
	ProductItemStatus_FullyApprovedInEP        string = "fully_approved_in_cp"

	ProductItemStatus_ProcessingEnter         string = "approved_by_oukd" // "Ожидает ввода"
	ProductItemStatus_EnteredInMagnum         string = "approved_by_oukd" // "Введен в ассортимент"
	ProductItemStatus_AssortimentEntryDelayed string = "approved_by_oukd" // "Ввод в АМ отложен"
	ProductItemStatus_Closed                  string = "approved_by_oukd" // Выведен
	ProductItemStatus_Blocked                 string = "approved_by_oukd" // Заблокирован
	ProductItemStatus_AwaitingWithdraw                                    // Ожидает вывода
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

type EntryProposal struct {
	ProductItem []ProductItem
}

type EntryProposalStatus string

const (
	EntryProposalStatus_Draft          string = "draft"
	EntryProposalStatus_ProcessingBUCP string = "processing_bucp"
	EntryProposalStatus_ApprovedByBUCP string = "approved_by_bucp"
	EntryProposalStatus_ProcessingOUKD string = "processing_oukd"
	EntryProposalStatus_ApprovedByOUKD string = "approved_by_oukd"
	EntryProposalStatus_FullyApproved  string = "fully_approved"
)

type ProductField struct {
	ID          int
	FieldCode   string
	FieldNameKz string
	FieldNameRu string
	FieldNameEn string
	SubCategory SubCategory
}

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

//DeliveryConditions   PackageType  BarcodeMask  Country  Color  TNVD
//PackageType   ChildWeight  StorageAndTransportationConditions  DeliveryConditions

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
