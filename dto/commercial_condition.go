package dto

type CommercialConditionsDTO struct {
	ID                         int64
	FixBonus                   int64
	FixPremium                 int64
	LogistBonus                int64
	PaymentDueDate             string
	AccountWork                string
	Inkoterms                  string
	MinimalOrderLot            string
	CommercialConditionsStatus string
	PaymentType                PaymentTypeDTO
	Organization               OrganizationDTO
}
