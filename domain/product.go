package domain

import "time"

type Product struct {
	ItemNameKz                     string
	ItemNameRu                     string
	ItemNameEn                     string
	DescriptionKz                  string
	DescriptionRu                  string
	DescriptionEn                  string
	ExternalId                     int64
	Barcode                        string
	SkuFeaturesForItemName         string
	ItemPriceTagName               string
	Supplier                       string
	Trademark                      string
	Brand                          string
	Manufacturer                   string
	CompositionKazakhLang          string
	CompositionRussianLang         string
	AssignBarcode                  bool
	VendorCode                     string
	Material                       string
	SalePeriodDays                 int
	ExpirationDateAfterOpeningDays int
	Density                        string
	Watt                           float64
	Ampere                         float64
	Plinth                         string
	Model                          string
	Completeness                   string
	HalalSign                      bool
	LactoseFreeSign                bool
	GlutenFreeSign                 bool
	SugarFreeSign                  bool
	KosherSign                     bool
	OrganicSign                    bool
	VeganSign                      bool
	DiabeticSign                   bool
	NaturalLossRate                float64
	Proteins                       float64

	/**
	 * Жиры
	 */
	Fats float64
	/**
	 * Углеводы
	 */
	Carbohydrates float64
	/**
	 * Энергетическая ценность
	 */
	EnergyValueKcal                    float64
	TouchSensitivitySign               bool
	EthyleneSensitivity                bool
	AlcoholStrength                    float64
	TareWeight                         float64
	DryWeight                          float64
	MinimumStorageTemperature          string
	MaximumStorageTemperature          string
	BlockAttribute                     bool
	NumberOfPieceskgInBlock            float64
	LengthDepthCm                      float64
	WidthCm                            float64
	HeightCm                           float64
	WeightCm                           float64
	UnitGrossWeightKg                  float64
	ConformityCertificateEndDate       time.Time
	PositionNameInStateLanguage        string
	PriceSegment                       string
	SalePriceWithVATtgCC               float64
	SalePriceWithVATtgATAK             float64
	PinCode                            string
	Packaging                          string
	Load                               float64
	Foliage                            string
	FatContent                         float64
	LogistictsOptions                  string
	Size                               string
	AdditionalOptions                  string
	AgeCategory                        string
	BlockAttributes                    BlockAttributes
	BarcodeMask                        BarcodeMask
	Country                            Country
	TNVD                               TNVD
	Color                              Color
	Taste                              string
	PackageType                        PackageType
	ChildWeight                        ChildWeight
	StorageAndTransportationConditions StorageAndTransportationConditions
	DeliveryConditions                 DeliveryConditions
	Measurement                        Measurement
	Organization                       Organization
	SubCategory                        SubCategory
	Subgroup                           Subgroup
	ItemPhotoWithLine                  int64
	ItemPhotoFrontSide                 int64
	CompositionPhoto                   int64
	ItemPhotoForScales                 int64
}

type BlockAttributes struct {
	/**
	 * Block - length deep
	 */
	LengthDeep float64

	/**
	 * Block - width
	 */
	Width float64

	/**
	 * Block - height
	 */
	Height float64

	/**
	 * Block - netWeight
	 */
	NetWeight float64

	/**
	 * Block - grossWeight
	 */
	GrossWeight float64

	/**
	 * Block - код блока
	 */
	WblockCode string
}

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

type ProductField struct {
	ID          int
	FieldCode   string
	FieldNameKz string
	FieldNameRu string
	FieldNameEn string
	SubCategory SubCategory
}
