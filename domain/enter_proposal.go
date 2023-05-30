package domain

type EnterProductApplication struct {
	ProductItem []ProductItem
}

type EnterProductApplicationStatus string

const (
	EnterProductApplicationStatus_Draft          string = "draft"
	EnterProductApplicationStatus_ProcessingBUCP string = "processing_bucp"
	EnterProductApplicationStatus_ApprovedByBUCP string = "approved_by_bucp"
	EnterProductApplicationStatus_ProcessingOUKD string = "processing_oukd"
	EnterProductApplicationStatus_ApprovedByOUKD string = "approved_by_oukd"
	EnterProductApplicationStatus_FullyApproved  string = "fully_approved"
)
