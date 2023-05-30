package domain

type FormOfLaw struct {
	ID         int
	Code       string
	IsResident bool
	NameKz     string
	NameRu     string
	NameEn     string
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
