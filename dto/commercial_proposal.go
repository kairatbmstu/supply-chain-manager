package dto

import "time"

type CommercialProposal struct {
	Id                       int64
	ProcessDate              time.Time
	CommercialProposalStatus string
	ProposalConditions       ProposalConditions
	Organization             OrganizationDTO
	Department               DepartmentDTO
	Currency                 string
	DeliveryConditions       DeliveryConditionsDTO
	CommercialProposalInfo   CommercialProposalInfo
	ProductItem              []ProductItemDTO
}

type ProposalConditions struct {
	ID              int64
	FixBonus        float64
	FixPremium      float64
	LogistBonus     float64
	PaymentDueDate  string
	AccountWork     string
	Inkoterms       string
	MinimalOrderLot string
	PaymentType     PaymentTypeDTO
	Regions         []RegionDTO
}

type CommercialProposalInfo struct {
}
