package domain

import "time"

type CommercialProposal struct {
	Id                       int64
	ProcessDate              time.Time
	CommercialProposalStatus CommercialProposalStatus
	ProposalConditions       ProposalConditions
	Organization             Organization
	Department               Department
	Currency                 Currency
	DeliveryConditions       DeliveryConditions
	CommercialProposalInfo   CommercialProposalInfo
	ProductItem              []ProductItem
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
	PaymentType     PaymentType
	Regions         []Region
}

type CommercialProposalInfo struct {
}
