package dto

import (
	"time"

	"example.com/m/v2/domain"
)

type ProductDTO struct {
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
	BarcodeMask                        domain.BarcodeMask
	Country                            CountryDTO
	TNVD                               TnvdDTO
	Color                              ColorDTO
	Taste                              string
	PackageType                        PackageTypeDTO
	ChildWeight                        ChildWeightDTO
	StorageAndTransportationConditions StorageAndTransportationConditionsDTO
	DeliveryConditions                 DeliveryConditionsDTO
	Measurement                        MeasurementDTO
	Organization                       OrganizationDTO
	SubCategory                        SubCategoryDTO
	Subgroup                           SubgroupDTO
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
