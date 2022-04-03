package domain

type CommercialConditions struct {
	ID                         int64
	FixBonus                   int64
	FixPremium                 int64
	LogistBonus                int64
	PaymentDueDate             string
	AccountWork                string
	Inkoterms                  string
	MinimalOrderLot            string
	CommercialConditionsStatus CommercialConditionsStatus
	PaymentType                PaymentType
	Organization               Organization
}

type CommercialConditionsStatus string

const (
	CommercialConditionsStatus_Draft      = "draft"
	CommercialConditionsStatus_ReadyForKP = "ready_for_kp"
)
